package jklmining

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/jackal-dao/canine/testutil/sample"
	jklminingsimulation "github.com/jackal-dao/canine/x/jklmining/simulation"
	"github.com/jackal-dao/canine/x/jklmining/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = jklminingsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgAllowSave = "op_weight_msg_allow_save"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAllowSave int = 100

	opWeightMsgCreateSaveRequests = "op_weight_msg_save_requests"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateSaveRequests int = 100

	opWeightMsgUpdateSaveRequests = "op_weight_msg_save_requests"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateSaveRequests int = 100

	opWeightMsgDeleteSaveRequests = "op_weight_msg_save_requests"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteSaveRequests int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	jklminingGenesis := types.GenesisState{
		SaveRequestsList: []types.SaveRequests{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&jklminingGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAllowSave int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAllowSave, &weightMsgAllowSave, nil,
		func(_ *rand.Rand) {
			weightMsgAllowSave = defaultWeightMsgAllowSave
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAllowSave,
		jklminingsimulation.SimulateMsgAllowSave(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}