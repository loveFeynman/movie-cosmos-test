package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "movie/testutil/keeper"
	"movie/testutil/nullify"
	"movie/x/movie/keeper"
	"movie/x/movie/types"
)

func createNMovie(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Movie {
	items := make([]types.Movie, n)
	for i := range items {
		items[i].Id = keeper.AppendMovie(ctx, items[i])
	}
	return items
}

func TestMovieGet(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNMovie(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMovie(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMovieRemove(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNMovie(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMovie(ctx, item.Id)
		_, found := keeper.GetMovie(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMovieGetAll(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNMovie(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMovie(ctx)),
	)
}

func TestMovieCount(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNMovie(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMovieCount(ctx))
}
