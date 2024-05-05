package keeper

import (
	"strconv"
	"sync"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

var straysPool = sync.Pool{
	New: func() interface{} {
		return new(types.Strays)
	},
}

// SetStrays set a specific strays in the store from its index
func (k Keeper) SetStrays(ctx sdk.Context, strays types.Strays) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StraysKeyPrefix))
	b := k.cdc.MustMarshal(&strays)
	store.Set(types.StraysKey(strays.Cid), b)
}

// GetStrays returns a strays from its index
func (k Keeper) GetStrays(ctx sdk.Context, cid string) (val types.Strays, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StraysKeyPrefix))
	b := store.Get(types.StraysKey(cid))
	if b == nil {
		return val, false
	}
	val, found = k.decodeStrays(b)
	return
}

func (k Keeper) decodeStrays(bz []byte) (val types.Strays, found bool) {
	valPtr := straysPool.Get().(*types.Strays)
	if valPtr == nil {
		return types.Strays{}, false
	}
	val = *valPtr
	found = k.cdc.Unmarshal(bz, &val) == nil
	straysPool.Put(valPtr) // Reset the pool
	return
}

// RemoveStrays removes a strays from the store
func (k Keeper) RemoveStrays(ctx sdk.Context, cid string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StraysKeyPrefix))
	store.Delete(types.StraysKey(cid))
}

// GetAllStrays returns all strays
func (k Keeper) GetAllStrays(ctx sdk.Context) []*types.Strays {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StraysKeyPrefix))
	iter := store.Iterator(nil, nil)
	defer iter.Close()

	var allStrays []*types.Strays
	for ; iter.Valid(); iter.Next() {
		val, _ := k.decodeStrays(iter.Value())
		allStrays = append(allStrays, &val)
	}
	return allStrays
}

func (k Keeper) ClearDeadFiles(ctx sdk.Context) {
	strays := k.GetAllStrays(ctx)
	deals := k.GetAllActiveDeals(ctx)

	var dealFids map[string]struct{}

	for _, stray := range strays {
		found := false
		if dealFids == nil {
			dealFids = make(map[string]struct{}, len(deals))
			for _, deal := range deals {
				dealFids[deal.Fid] = struct{}{}
			}
		}
		_, found = dealFids[stray.Fid]
		if found {
			continue
		}

		paymentInfo, found := k.GetStoragePaymentInfo(ctx, stray.Signee)
		if found {
			fSize, err := strconv.ParseInt(stray.Filesize, 10, 64)
			if err == nil {
				paymentInfo.SpaceUsed -= fSize // remove the deal from the users paid amount.
			}
		}
		k.RemoveStrays(ctx, stray.Cid)
	}
}
