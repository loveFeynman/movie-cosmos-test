package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MovieList: []Movie{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in movie
	movieIdMap := make(map[uint64]bool)
	movieCount := gs.GetMovieCount()
	for _, elem := range gs.MovieList {
		if _, ok := movieIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for movie")
		}
		if elem.Id >= movieCount {
			return fmt.Errorf("movie id should be lower or equal than the last id")
		}
		movieIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
