package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const LastBlockKey = "LastBlockMinted"

func (k Keeper) SetLastBlockMinted(ctx sdk.Context, coin sdk.DecCoin) error {

	coinData, err := coin.Marshal()
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(LastBlockKey), coinData)

	return nil
}

func (k Keeper) defaultCoin(ctx sdk.Context) sdk.DecCoin {
	params := k.GetParams(ctx)
	return sdk.NewDecCoin(params.MintDenom, sdk.NewInt(6_000_000)) // some default amount
}

func (k Keeper) GetLastBlockMinted(ctx sdk.Context) sdk.DecCoin {

	store := ctx.KVStore(k.storeKey)
	coinData := store.Get([]byte(LastBlockKey))
	if coinData == nil {
		return k.defaultCoin(ctx) // return default amount
	}

	var oldCoin sdk.DecCoin
	err := oldCoin.Unmarshal(coinData)
	if err != nil {
		return k.defaultCoin(ctx) // return default amount
	}

	return oldCoin

}
