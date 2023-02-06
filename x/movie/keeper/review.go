package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"movie/x/movie/types"
)

// GetReviewCount get the total number of review
func (k Keeper) GetReviewCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ReviewCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetReviewCount set the total number of review
func (k Keeper) SetReviewCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ReviewCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendReview appends a review in the store with a new id and update the count
func (k Keeper) AppendReview(
	ctx sdk.Context,
	review types.Review,
) uint64 {
	// Create the review
	count := k.GetReviewCount(ctx)

	// Set the ID of the appended value
	review.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewKey))
	appendedValue := k.cdc.MustMarshal(&review)
	store.Set(GetReviewIDBytes(review.Id), appendedValue)

	// Update review count
	k.SetReviewCount(ctx, count+1)

	return count
}

// SetReview set a specific review in the store
func (k Keeper) SetReview(ctx sdk.Context, review types.Review) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewKey))
	b := k.cdc.MustMarshal(&review)
	store.Set(GetReviewIDBytes(review.Id), b)
}

// GetReview returns a review from its id
func (k Keeper) GetReview(ctx sdk.Context, id uint64) (val types.Review, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewKey))
	b := store.Get(GetReviewIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveReview removes a review from the store
func (k Keeper) RemoveReview(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewKey))
	store.Delete(GetReviewIDBytes(id))
}

// GetAllReview returns all review
func (k Keeper) GetAllReview(ctx sdk.Context) (list []types.Review) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Review
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetReviewIDBytes returns the byte representation of the ID
func GetReviewIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetReviewIDFromBytes returns ID in uint64 format from a byte array
func GetReviewIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
