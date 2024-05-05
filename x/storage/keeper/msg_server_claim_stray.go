package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k msgServer) ClaimStray(goCtx context.Context, msg *types.MsgClaimStray) (*types.MsgClaimStrayResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	stray, ok := k.GetStrays(ctx, msg.Cid)
	if !ok {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "stray contract either no longer is stray, or has been removed by the user")
	}

	ls := k.ListFiles(ctx, stray.Fid)

	provider, found := k.GetProviders(ctx, msg.ForAddress)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not a provider")
	}

	for _, l := range ls {
		if l == provider.Ip {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot claim a stray you own.")
		}
	}

	deal := activeDealsPool.Get().(*types.ActiveDeals)
	deal.Cid = stray.Cid
	deal.Signee = stray.Signee
	deal.Provider = msg.ForAddress
	deal.Creator = msg.Creator
	deal.Fid = stray.Fid
	deal.Merkle = stray.Merkle

	size, ok := sdk.NewIntFromString(stray.Filesize)
	if !ok {
		activeDealsPool.Put(deal)
		return nil, fmt.Errorf("failed to parse filesize: %s %v", "%t", ok)
	}

	if size.IsZero() {
		activeDealsPool.Put(deal)
		return nil, fmt.Errorf("filesize is nil")
	}
	pieces := size.Quo(sdk.NewInt(k.GetParams(ctx).ChunkSize))
	deal.Startblock = fmt.Sprintf("%d", ctx.BlockHeight())
	deal.Endblock = fmt.Sprintf("%d", stray.End)
	deal.Filesize = stray.Filesize
	deal.LastProof = ctx.BlockHeight()
	deal.Blocktoprove = fmt.Sprintf("%d", ctx.BlockHeight()%pieces.Int64())
	deal.Proofsmissed = "0"

	if msg.ForAddress != msg.Creator {
		for _, claimer := range provider.AuthClaimers {
			if claimer == msg.Creator {
				k.SetActiveDeals(ctx, *deal)
				activeDealsPool.Put(deal)
				return &types.MsgClaimStrayResponse{}, nil
			}
		}
		activeDealsPool.Put(deal)
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "cannot claim a stray for someone else")
	}

	k.SetActiveDeals(ctx, *deal)
	activeDealsPool.Put(deal)

	return &types.MsgClaimStrayResponse{}, nil
}
