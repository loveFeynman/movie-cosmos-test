package keeper

import (
	"context"
	"fmt"

	"movie/x/movie/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateReview(goCtx context.Context, msg *types.MsgCreateReview) (*types.MsgCreateReviewResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var review = types.Review{
		Creator:     msg.Creator,
		MovieId:     msg.MovieId,
		RatingUint:  msg.RatingUint,
		Description: msg.Description,
	}

	// Only allow creating reviews once per account for the same movie
	reviewList := k.GetAllReview(ctx)
	for _, individualReview := range reviewList {
		if individualReview.Creator == review.Creator && individualReview.MovieId == review.MovieId {
			return nil, sdkerrors.Wrapf(types.ErrReviewDuplicationForOneAccount, "Cannot perform this tx for reviews")
		}
	}

	// throw an error when creating a review while the movie doesn’t exist.
	movieList := k.GetAllMovie(ctx)
	for i, individualMovie := range movieList {
		if individualMovie.Id == review.MovieId {
			break
		}
		if i == len(movieList)-1 {
			return nil, sdkerrors.Wrapf(types.ErrNoSuchMoive, "Cannot perform this tx for reviews")
		}
	}

	id := k.AppendReview(
		ctx,
		review,
	)

	return &types.MsgCreateReviewResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateReview(goCtx context.Context, msg *types.MsgUpdateReview) (*types.MsgUpdateReviewResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var review = types.Review{
		Creator:     msg.Creator,
		Id:          msg.Id,
		MovieId:     msg.MovieId,
		RatingUint:  msg.RatingUint,
		Description: msg.Description,
	}

	// Checks that the element exists
	val, found := k.GetReview(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetReview(ctx, review)

	return &types.MsgUpdateReviewResponse{}, nil
}

func (k msgServer) DeleteReview(goCtx context.Context, msg *types.MsgDeleteReview) (*types.MsgDeleteReviewResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetReview(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveReview(ctx, msg.Id)

	return &types.MsgDeleteReviewResponse{}, nil
}
