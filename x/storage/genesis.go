package storage

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v3/x/storage/keeper"
	"github.com/jackalLabs/canine-chain/v3/x/storage/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the contracts
	for _, elem := range genState.ContractsList {
		k.SetContracts(ctx, elem)
	}
	// Set all the activeDeals
	for _, elem := range genState.ActiveDealsList {
		k.SetActiveDeals(ctx, elem)
	}
	// Set all the Providers
	for _, elem := range genState.ProvidersList {
		k.SetProviders(ctx, elem)
	}

	// Set all the strays
	for _, elem := range genState.StraysList {
		k.SetStrays(ctx, elem)
	}
	// Set all the fidCid
	for _, elem := range genState.FidCidList {
		k.SetFidCid(ctx, elem)
	}

	// Set all the paymentinfo
	for _, elem := range genState.PaymentInfoList {
		k.SetStoragePaymentInfo(ctx, elem)
	}

	// Set all the collateral
	for _, elem := range genState.CollateralList {
		k.SetCollateral(ctx, elem)
	}

	// Set all the attestations
	for _, elem := range genState.AttestForms {
		k.SetAttestationForm(ctx, elem)
	}

	// Set all the reports
	for _, elem := range genState.ReportForms {
		k.SetReportForm(ctx, elem)
	}

	// Set all the active providers
	for _, elem := range genState.ActiveProvidersList {
		k.SetActiveProviders(ctx, elem)
	}

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Copy the Strays objects into a new slice
	straysList := make([]types.Strays, len(k.GetAllStrays(ctx)))
	for i, stray := range k.GetAllStrays(ctx) {
		straysList[i] = *stray
	}

	genesis.ContractsList = k.GetAllContracts(ctx)
	genesis.ActiveDealsList = k.GetAllActiveDeals(ctx)
	genesis.ProvidersList = k.GetAllProviders(ctx)
	genesis.StraysList = straysList
	genesis.FidCidList = k.GetAllFidCid(ctx)
	genesis.PaymentInfoList = k.GetAllStoragePaymentInfo(ctx)
	genesis.CollateralList = k.GetAllCollateral(ctx)
	genesis.ActiveProvidersList = k.GetAllActiveProviders(ctx)
	genesis.ReportForms = k.GetAllReport(ctx)
	genesis.AttestForms = k.GetAllAttestation(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
