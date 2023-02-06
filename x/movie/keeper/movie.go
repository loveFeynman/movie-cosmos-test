package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"movie/x/movie/types"
)

// GetMovieCount get the total number of movie
func (k Keeper) GetMovieCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MovieCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMovieCount set the total number of movie
func (k Keeper) SetMovieCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.MovieCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMovie appends a movie in the store with a new id and update the count
func (k Keeper) AppendMovie(
	ctx sdk.Context,
	movie types.Movie,
) uint64 {
	// Create the movie
	count := k.GetMovieCount(ctx)

	// Set the ID of the appended value
	movie.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MovieKey))
	appendedValue := k.cdc.MustMarshal(&movie)
	store.Set(GetMovieIDBytes(movie.Id), appendedValue)

	// Update movie count
	k.SetMovieCount(ctx, count+1)

	return count
}

// SetMovie set a specific movie in the store
func (k Keeper) SetMovie(ctx sdk.Context, movie types.Movie) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MovieKey))
	b := k.cdc.MustMarshal(&movie)
	store.Set(GetMovieIDBytes(movie.Id), b)
}

// GetMovie returns a movie from its id
func (k Keeper) GetMovie(ctx sdk.Context, id uint64) (val types.Movie, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MovieKey))
	b := store.Get(GetMovieIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMovie removes a movie from the store
func (k Keeper) RemoveMovie(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MovieKey))
	store.Delete(GetMovieIDBytes(id))
}

// GetAllMovie returns all movie
func (k Keeper) GetAllMovie(ctx sdk.Context) (list []types.Movie) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MovieKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Movie
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMovieIDBytes returns the byte representation of the ID
func GetMovieIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMovieIDFromBytes returns ID in uint64 format from a byte array
func GetMovieIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
