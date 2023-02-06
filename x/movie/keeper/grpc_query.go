package keeper

import (
	"movie/x/movie/types"
)

var _ types.QueryServer = Keeper{}
