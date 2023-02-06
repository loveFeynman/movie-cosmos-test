package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"movie/x/movie/types"
)

func (k Keeper) ReviewAll(c context.Context, req *types.QueryAllReviewRequest) (*types.QueryAllReviewResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var reviews []types.Review
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	reviewStore := prefix.NewStore(store, types.KeyPrefix(types.ReviewKey))

	pageRes, err := query.Paginate(reviewStore, req.Pagination, func(key []byte, value []byte) error {
		var review types.Review
		if err := k.cdc.Unmarshal(value, &review); err != nil {
			return err
		}

		reviews = append(reviews, review)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllReviewResponse{Review: reviews, Pagination: pageRes}, nil
}

func (k Keeper) Review(c context.Context, req *types.QueryGetReviewRequest) (*types.QueryGetReviewResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	review, found := k.GetReview(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetReviewResponse{Review: review}, nil
}
