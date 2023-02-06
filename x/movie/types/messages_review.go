package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateReview = "create_review"
	TypeMsgUpdateReview = "update_review"
	TypeMsgDeleteReview = "delete_review"
)

var _ sdk.Msg = &MsgCreateReview{}

func NewMsgCreateReview(creator string, movieId uint64, ratingUint string, description string) *MsgCreateReview {
	return &MsgCreateReview{
		Creator:     creator,
		MovieId:     movieId,
		RatingUint:  ratingUint,
		Description: description,
	}
}

func (msg *MsgCreateReview) Route() string {
	return RouterKey
}

func (msg *MsgCreateReview) Type() string {
	return TypeMsgCreateReview
}

func (msg *MsgCreateReview) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateReview) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateReview) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateReview{}

func NewMsgUpdateReview(creator string, id uint64, movieId uint64, ratingUint string, description string) *MsgUpdateReview {
	return &MsgUpdateReview{
		Id:          id,
		Creator:     creator,
		MovieId:     movieId,
		RatingUint:  ratingUint,
		Description: description,
	}
}

func (msg *MsgUpdateReview) Route() string {
	return RouterKey
}

func (msg *MsgUpdateReview) Type() string {
	return TypeMsgUpdateReview
}

func (msg *MsgUpdateReview) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateReview) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateReview) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteReview{}

func NewMsgDeleteReview(creator string, id uint64) *MsgDeleteReview {
	return &MsgDeleteReview{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteReview) Route() string {
	return RouterKey
}

func (msg *MsgDeleteReview) Type() string {
	return TypeMsgDeleteReview
}

func (msg *MsgDeleteReview) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteReview) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteReview) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
