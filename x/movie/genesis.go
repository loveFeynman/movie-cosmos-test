package movie

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"movie/x/movie/keeper"
	"movie/x/movie/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the movie
	for _, elem := range genState.MovieList {
		k.SetMovie(ctx, elem)
	}

	// Set movie count
	k.SetMovieCount(ctx, genState.MovieCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.MovieList = k.GetAllMovie(ctx)
	genesis.MovieCount = k.GetMovieCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
