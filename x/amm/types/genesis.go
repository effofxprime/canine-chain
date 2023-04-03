package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PoolList:           []Pool{},
		ProviderRecordList: []ProviderRecord{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in Pool
	lPoolIndexMap := make(map[string]struct{})

	for _, elem := range gs.PoolList {
		index := string(PoolKey(elem.Index))
		if _, ok := lPoolIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for lPool")
		}
		lPoolIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in ProviderRecord
	lProviderRecordIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProviderRecordList {
		index := string(ProviderRecordKey(elem.PoolName, elem.Provider))
		if _, ok := lProviderRecordIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for lProviderRecord")
		}
		lProviderRecordIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}