package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"movie/x/movie/types"
)

func (k msgServer) CreateMovie(goCtx context.Context, msg *types.MsgCreateMovie) (*types.MsgCreateMovieResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var movie = types.Movie{
		Creator:     msg.Creator,
		Title:       msg.Title,
		Description: msg.Description,
		Year:        msg.Year,
	}

	id := k.AppendMovie(
		ctx,
		movie,
	)

	return &types.MsgCreateMovieResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateMovie(goCtx context.Context, msg *types.MsgUpdateMovie) (*types.MsgUpdateMovieResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var movie = types.Movie{
		Creator:     msg.Creator,
		Id:          msg.Id,
		Title:       msg.Title,
		Description: msg.Description,
		Year:        msg.Year,
	}

	// Checks that the element exists
	val, found := k.GetMovie(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetMovie(ctx, movie)

	return &types.MsgUpdateMovieResponse{}, nil
}

func (k msgServer) DeleteMovie(goCtx context.Context, msg *types.MsgDeleteMovie) (*types.MsgDeleteMovieResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetMovie(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMovie(ctx, msg.Id)

	return &types.MsgDeleteMovieResponse{}, nil
}
