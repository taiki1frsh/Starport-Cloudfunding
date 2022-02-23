package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStopProject = "stop_project"

var _ sdk.Msg = &MsgStopProject{}

func NewMsgStopProject(creator string, id uint64) *MsgStopProject {
	return &MsgStopProject{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgStopProject) Route() string {
	return RouterKey
}

func (msg *MsgStopProject) Type() string {
	return TypeMsgStopProject
}

func (msg *MsgStopProject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStopProject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStopProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
