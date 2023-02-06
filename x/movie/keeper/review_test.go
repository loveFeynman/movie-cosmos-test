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

func createNReview(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Review {
	items := make([]types.Review, n)
	for i := range items {
		items[i].Id = keeper.AppendReview(ctx, items[i])
	}
	return items
}

func TestReviewGet(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNReview(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetReview(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestReviewRemove(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNReview(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveReview(ctx, item.Id)
		_, found := keeper.GetReview(ctx, item.Id)
		require.False(t, found)
	}
}

func TestReviewGetAll(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNReview(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllReview(ctx)),
	)
}

func TestReviewCount(t *testing.T) {
	keeper, ctx := keepertest.MovieKeeper(t)
	items := createNReview(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetReviewCount(ctx))
}
