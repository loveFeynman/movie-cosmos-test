package movie

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"movie/testutil/sample"
	moviesimulation "movie/x/movie/simulation"
	"movie/x/movie/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = moviesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateMovie = "op_weight_msg_movie"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateMovie int = 100

	opWeightMsgUpdateMovie = "op_weight_msg_movie"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMovie int = 100

	opWeightMsgDeleteMovie = "op_weight_msg_movie"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteMovie int = 100

	opWeightMsgCreateReview = "op_weight_msg_review"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateReview int = 100

	opWeightMsgUpdateReview = "op_weight_msg_review"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateReview int = 100

	opWeightMsgDeleteReview = "op_weight_msg_review"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteReview int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	movieGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		MovieList: []types.Movie{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		MovieCount: 2,
		ReviewList: []types.Review{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		ReviewCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&movieGenesis)
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

	var weightMsgCreateMovie int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateMovie, &weightMsgCreateMovie, nil,
		func(_ *rand.Rand) {
			weightMsgCreateMovie = defaultWeightMsgCreateMovie
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateMovie,
		moviesimulation.SimulateMsgCreateMovie(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateMovie int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMovie, &weightMsgUpdateMovie, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMovie = defaultWeightMsgUpdateMovie
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMovie,
		moviesimulation.SimulateMsgUpdateMovie(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteMovie int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteMovie, &weightMsgDeleteMovie, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteMovie = defaultWeightMsgDeleteMovie
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteMovie,
		moviesimulation.SimulateMsgDeleteMovie(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateReview int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateReview, &weightMsgCreateReview, nil,
		func(_ *rand.Rand) {
			weightMsgCreateReview = defaultWeightMsgCreateReview
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateReview,
		moviesimulation.SimulateMsgCreateReview(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateReview int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateReview, &weightMsgUpdateReview, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateReview = defaultWeightMsgUpdateReview
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateReview,
		moviesimulation.SimulateMsgUpdateReview(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteReview int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteReview, &weightMsgDeleteReview, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteReview = defaultWeightMsgDeleteReview
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteReview,
		moviesimulation.SimulateMsgDeleteReview(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
