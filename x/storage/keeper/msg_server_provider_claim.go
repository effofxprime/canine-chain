package keeper

import (
	"context"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/jackalLabs/canine-chain/x/storage/types"
)

func (k msgServer) AddProviderClaimer(goCtx context.Context, msg *types.MsgAddClaimer) (*types.MsgAddClaimerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Creator)

	if !found {
		return nil, types.ErrProviderNotFound
	}

	if provider.AuthClaimers == nil {
		provider.AuthClaimers = []string{}
	}

	for _, claimer := range provider.AuthClaimers {
		if claimer == msg.ClaimAddress {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrConflict, "cannot add the same claimer twice")
		}
	}

	provider.AuthClaimers = append(provider.AuthClaimers, msg.ClaimAddress)

	k.SetProviders(ctx, provider)

	return &types.MsgAddClaimerResponse{}, nil
}

func (k msgServer) RemoveProviderClaimer(goCtx context.Context, msg *types.MsgRemoveClaimer) (*types.MsgRemoveClaimerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	provider, found := k.GetProviders(ctx, msg.Creator)
	authClaimers := provider.AuthClaimers

	if !found {
		return nil, types.ErrProviderNotFound
	}

	if len(authClaimers) == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrConflict, "Provider has no claimer addresses")
	}

	newClaim := []string{}

	index := sort.SearchStrings(authClaimers, msg.ClaimAddress)

	if index == len(authClaimers) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrConflict, "this address is not a claimer")
	}

	for i := 0; i < len(authClaimers); i++ {
		if i != index {
			newClaim = append(newClaim, authClaimers[i])
		}
	}
	provider.AuthClaimers = newClaim

	k.SetProviders(ctx, provider)

	return &types.MsgRemoveClaimerResponse{}, nil
}