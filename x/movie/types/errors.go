package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/movie module sentinel errors
var (
	ErrSample                         = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrDuplicationTitle               = sdkerrors.Register(ModuleName, 1101, "title duplication error")
	ErrNoSuchMoive                    = sdkerrors.Register(ModuleName, 1102, "no such a movie")
	ErrReviewDuplicationForOneAccount = sdkerrors.Register(ModuleName, 1103, "that account already created a reivew for this movie")
)
