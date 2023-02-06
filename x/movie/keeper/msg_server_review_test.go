package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"movie/x/movie/types"
)

func TestReviewMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)

	// create a list of movies
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: "B", Title: "Movie Title " + strconv.Itoa(i)})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}

	creator := "A"
	// Check if we can create a review on existing movie
	resp, err := srv.CreateReview(ctx, &types.MsgCreateReview{Creator: creator, MovieId: 3})
	require.NoError(t, err)
	require.Equal(t, 0, int(resp.Id))

	// throw an error when creating a review while the movie doesn’t exist.
	_, err = srv.CreateReview(ctx, &types.MsgCreateReview{Creator: creator, MovieId: 8})
	require.EqualError(t, err, "Cannot perform this tx for reviews: no such a movie")

}

func TestReviewMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateReview
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateReview{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateReview{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateReview{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateReview(ctx, &types.MsgCreateReview{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateReview(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestReviewMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteReview
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteReview{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteReview{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteReview{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateReview(ctx, &types.MsgCreateReview{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteReview(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
