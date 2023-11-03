package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/jklmint/types"
	storeTypes "github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// Mint creates a new token to mint from the previous blocks information, it returns a decimal free coin to mint
func (k Keeper) Mint(ctx sdk.Context) (sdk.Coin, error) {

	mintParams := k.GetParams(ctx)

	mintDecrease, err := sdk.NewDecFromStr(mintParams.MintDecrease)
	if err != nil {
		return sdk.Coin{}, err
	}

	lastMinted := k.GetLastBlockMinted(ctx) // if this is the first time this is run, it will return the default minting amount
	lastMintedAmt := lastMinted.Amount

	blocksPerYear := sdk.NewDec(5_256_000)

	newMinted := lastMintedAmt.Sub(mintDecrease.Quo(blocksPerYear))

	newMintedDecCoin := sdk.NewDecCoin(mintParams.MintDenom, newMinted.TruncateInt())

	err = k.SetLastBlockMinted(ctx, newMintedDecCoin)
	if err != nil {
		return sdk.Coin{}, err
	}

	toMint, _ := newMintedDecCoin.TruncateDecimal()

	return toMint, nil
}

// BlockMint handles every blocks mint functionality
//
// TODO: Completely revamp function, currently it handles all minting & distribution logic which is bad
func (k Keeper) BlockMint(ctx sdk.Context) {
	tokensPerBlock := k.GetParams(ctx).TokensPerBlock

	mintTokens := sdk.NewDec(tokensPerBlock * 1_000_000)
	denom := k.GetParams(ctx).MintDenom

	pRatio := k.GetParams(ctx).ProviderRatio
	valRatio := 10 - pRatio

	providerRatio := sdk.NewDec(pRatio)
	providerRatio = providerRatio.QuoInt64(10)
	validatorRatio := sdk.NewDec(valRatio)
	validatorRatio = validatorRatio.QuoInt64(10)

	// get correct ratio
	providerTokens := mintTokens.Mul(providerRatio)
	validatorTokens := mintTokens.Mul(validatorRatio)

	// mint provider coins, update supply
	provCount := providerTokens.TruncateInt()
	providerCoin := sdk.NewCoin(denom, provCount)
	providerCoins := sdk.NewCoins(providerCoin)

	// mint validator coins, update supply
	valCount := validatorTokens.TruncateInt()
	validatorCoin := sdk.NewCoin(denom, valCount)
	validatorCoins := sdk.NewCoins(validatorCoin)

	totalCount := provCount.Add(valCount)
	// mint coins, update supply
	totalIntCoin := sdk.NewCoin(denom, totalCount)
	mintedCoin := totalIntCoin
	mintedCoins := sdk.NewCoins(mintedCoin)
	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, storeTypes.ModuleName, providerCoins)
	if err != nil {
		panic(err)
	}

	// send the minted validator coins to the fee collector account
	err = k.AddCollectedFees(ctx, validatorCoins)
	if err != nil {
		panic(err)
	}

	// alerting network on mint amount
	defer telemetry.ModuleSetGauge(types.ModuleName, float32(totalCount.Int64()), "minted_tokens")

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			sdk.NewAttribute(sdk.AttributeKeyAmount, fmt.Sprintf("%d", totalCount.Int64())),
		),
	)
}
