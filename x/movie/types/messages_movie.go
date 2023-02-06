package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateMovie = "create_movie"
	TypeMsgUpdateMovie = "update_movie"
	TypeMsgDeleteMovie = "delete_movie"
)

var _ sdk.Msg = &MsgCreateMovie{}

func NewMsgCreateMovie(creator string, title string, description string, year uint64) *MsgCreateMovie {
	return &MsgCreateMovie{
		Creator:     creator,
		Title:       title,
		Description: description,
		Year:        year,
	}
}

func (msg *MsgCreateMovie) Route() string {
	return RouterKey
}

func (msg *MsgCreateMovie) Type() string {
	return TypeMsgCreateMovie
}

func (msg *MsgCreateMovie) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateMovie) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateMovie) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMovie{}

func NewMsgUpdateMovie(creator string, id uint64, title string, description string, year uint64) *MsgUpdateMovie {
	return &MsgUpdateMovie{
		Id:          id,
		Creator:     creator,
		Title:       title,
		Description: description,
		Year:        year,
	}
}

func (msg *MsgUpdateMovie) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMovie) Type() string {
	return TypeMsgUpdateMovie
}

func (msg *MsgUpdateMovie) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMovie) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMovie) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMovie{}

func NewMsgDeleteMovie(creator string, id uint64) *MsgDeleteMovie {
	return &MsgDeleteMovie{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteMovie) Route() string {
	return RouterKey
}

func (msg *MsgDeleteMovie) Type() string {
	return TypeMsgDeleteMovie
}

func (msg *MsgDeleteMovie) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteMovie) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteMovie) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
