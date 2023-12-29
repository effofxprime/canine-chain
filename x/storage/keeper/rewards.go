package keeper

import (
	"fmt"
	"strconv"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k Keeper) burnContract(ctx sdk.Context, providerAddress string) {
	prov, found := k.GetProviders(ctx, providerAddress)
	if !found {
		return
	}

	burned, err := strconv.ParseInt(prov.BurnedContracts, 10, 64)
	if err != nil {
		ctx.Logger().Error("cannot parse providers burn count")
		return
	}

	prov.BurnedContracts = fmt.Sprintf("%d", burned+1)
	k.SetProviders(ctx, prov)
}

// manageProof checks the status of a given proof, if the file is too young, we skip it. If it's old enough and the
// prover has either failed to prove it or the proof simply never existed we remove it.
func (k Keeper) manageProof(ctx sdk.Context, sizeTracker *map[string]int64, file *types.ActiveDeals, proofWindow int64) {
	st := *sizeTracker

	// If we check the file and there is a proof delegated but the provider hasn't proven it yet we remove it.
	// However, we need to check if the file is new and is being caught by accident
	if !file.IsYoung(ctx.BlockHeight(), proofWindow) { // give first window grace before removal
		currentHeight := ctx.BlockHeight()

		proven := file.Proofverified == "true" // parsing bool

		if !proven { // if file has not been proven yet
			ctx.Logger().Info(fmt.Sprintf("proof has not been proven within the last window at %d", currentHeight))
			k.RemoveActiveDeals(ctx, file.Cid)
			k.burnContract(ctx, file.Provider)
			return
		}

		size, err := strconv.ParseInt(file.Filesize, 10, 64)
		if err == nil {
			st[file.Provider] += size // only give rewards to providers who have held onto the file for a full window
		}
	}
}

// TODO: Completely change the way this is done in Econ v2
func (k Keeper) rewardProviders(ctx sdk.Context, totalSize int64, sizeTracker *map[string]int64) {
	networkValue := sdk.NewDec(totalSize)

	storageWallet := k.accountkeeper.GetModuleAddress(types.ModuleName)

	tokens := k.bankkeeper.GetBalance(ctx, storageWallet, "ujkl")
	tokenAmountDec := tokens.Amount.ToDec()

	for prover, worth := range *sizeTracker {

		providerValue := sdk.NewDec(worth)

		networkPercentage := providerValue.Quo(networkValue)

		tokensValueOwed := networkPercentage.Mul(tokenAmountDec).TruncateInt64()

		coin := sdk.NewInt64Coin("ujkl", tokensValueOwed)
		coins := sdk.NewCoins(coin)

		pAddress, err := sdk.AccAddressFromBech32(prover)
		if err != nil {
			ctx.Logger().Error(sdkerrors.Wrapf(err, "failed to convert prover address %s to bech32", prover).Error())
			continue
		}
		err = k.bankkeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, pAddress, coins)
		if err != nil {
			ctx.Logger().Error(sdkerrors.Wrapf(err, "failed to send %d tokens to %s", tokensValueOwed, prover).Error())
			continue
		}
	}
}

func (k Keeper) removeFileIfDeserved(ctx sdk.Context, file *types.ActiveDeals, proofWindow int64) bool {
	if !file.IsYoung(ctx.BlockHeight(), proofWindow) { // give first window grace
		if file.Proofverified == "false" {
			k.RemoveActiveDeals(ctx, file.Cid)
			return true
		}
	}

	return false
}

// ManageRewards loops through every file on the network and manages it in some way.
func (k Keeper) ManageRewards(ctx sdk.Context, proofWindow int64) {
	var totalSize int64
	s := make(map[string]int64)
	sizeTracker := &s

	k.IterateAllActiveDeals(ctx, func(deal *types.ActiveDeals) {
		s, err := strconv.ParseInt(deal.Filesize, 10, 64)
		if err != nil { // skip errors
			return
		}

		totalSize += s

		if k.removeFileIfDeserved(ctx, deal, proofWindow) { // delete file if it meets the conditions to be deleted
			return
		}

		k.manageProof(ctx, sizeTracker, deal, proofWindow)
	})

	k.rewardProviders(ctx, totalSize, sizeTracker)
}

func (k Keeper) RunRewardBlock(ctx sdk.Context) {
	DayBlocks := k.GetParams(ctx).ProofWindow // checks more often than proofs take to catch them more frequently

	if ctx.BlockHeight()%DayBlocks > 0 { // runs once each window (usually a full days worth of blocks)
		ctx.Logger().Debug("skipping reward handling for this block")
		return
	}

	k.ManageRewards(ctx, DayBlocks)
}
