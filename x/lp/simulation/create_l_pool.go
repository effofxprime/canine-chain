package simulation

import (
	"math/rand"

	"github.com/jackal-dao/canine/x/lp/keeper"
	"github.com/jackal-dao/canine/x/lp/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateLPool(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateLPool{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateLPool simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateLPool simulation not implemented"), nil, nil
	}
}