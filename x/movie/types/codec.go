package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateMovie{}, "movie/CreateMovie", nil)
	cdc.RegisterConcrete(&MsgUpdateMovie{}, "movie/UpdateMovie", nil)
	cdc.RegisterConcrete(&MsgDeleteMovie{}, "movie/DeleteMovie", nil)
	cdc.RegisterConcrete(&MsgCreateReview{}, "movie/CreateReview", nil)
	cdc.RegisterConcrete(&MsgUpdateReview{}, "movie/UpdateReview", nil)
	cdc.RegisterConcrete(&MsgDeleteReview{}, "movie/DeleteReview", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateMovie{},
		&MsgUpdateMovie{},
		&MsgDeleteMovie{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateReview{},
		&MsgUpdateReview{},
		&MsgDeleteReview{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
