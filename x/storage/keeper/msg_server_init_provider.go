package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

func (k msgServer) InitProvider(goCtx context.Context, msg *types.MsgInitProvider) (*types.MsgInitProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetProviders(ctx, msg.Creator)
	if found {
		return nil, types.ErrProviderExists
	}

	provider := types.Providers{
		Address:         msg.Creator,
		Ip:              msg.Ip,
		Totalspace:      fmt.Sprintf("%d", msg.TotalSpace),
		Creator:         msg.Creator,
		BurnedContracts: "0",
		KeybaseIdentity: msg.Keybase,
		AuthClaimers:    []string{},
	}

	k.SetProviders(ctx, provider)

	return &types.MsgInitProviderResponse{}, nil
}

func (k msgServer) ShutdownProvider(goCtx context.Context, msg *types.MsgShutdownProvider) (*types.MsgShutdownProviderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetProviders(ctx, msg.Creator)
	if !found {
		return nil, types.ErrProviderNotFound
	}

	k.RemoveProviders(ctx, msg.Creator)

	return &types.MsgShutdownProviderResponse{}, nil
}
