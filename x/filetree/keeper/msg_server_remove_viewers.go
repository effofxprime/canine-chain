package keeper

import (
	"context"
	"encoding/json"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackal-dao/canine/x/filetree/types"
)

func (k msgServer) RemoveViewers(goCtx context.Context, msg *types.MsgRemoveViewers) (*types.MsgRemoveViewersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	file, found := k.GetFiles(ctx, msg.Address, msg.Fileowner)
	if !found {
		return nil, types.ErrFileNotFound
	}

	hasEdit := HasEditAccess(file, msg.Creator)
	if !hasEdit {
		return nil, types.ErrNoAccess
	}

	pvacc := file.ViewingAccess

	jvacc := make(map[string]string)
	json.Unmarshal([]byte(pvacc), &jvacc)

	ids := strings.Split(msg.ViewerIds, ",")
	for _, v := range ids {
		delete(jvacc, v)
	}

	vaccbytes, err := json.Marshal(jvacc)
	if err != nil {
		return nil, types.ErrCantMarshall
	}
	newviewers := string(vaccbytes)

	file.ViewingAccess = newviewers

	k.SetFiles(ctx, file)

	//notify viewers
	bool, error := notify(k, ctx, msg.Notifyviewers, string("You can no longer view this file"), msg.Creator, file.Address, file.Owner)
	if !bool {
		return nil, error
	}
	return &types.MsgRemoveViewersResponse{}, nil
}