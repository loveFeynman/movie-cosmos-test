package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"movie/x/movie/types"
)

func TestMovieMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator, Title: "Movie Title " + strconv.Itoa(i)})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}

	//  throw an error when creating a duplicate movie with the same title.
	_, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator, Title: "Movie Title 3"})
	require.EqualError(t, err, "Cannot perform this tx: title duplication error")
}

func TestMovieMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateMovie
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateMovie{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateMovie{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateMovie{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateMovie(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMovieMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteMovie
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteMovie{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteMovie{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteMovie{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateMovie(ctx, &types.MsgCreateMovie{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteMovie(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
