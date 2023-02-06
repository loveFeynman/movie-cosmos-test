package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "movie/testutil/keeper"
	"movie/x/movie/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MovieKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
