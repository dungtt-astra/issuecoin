package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMintcoin = "mintcoin"

var _ sdk.Msg = &MsgMintcoin{}

func NewMsgMintcoin(creator string, coin *sdk.Coin, to string) *MsgMintcoin {
	return &MsgMintcoin{
		Creator: creator,
		Coin:    coin,
		To:      to,
	}
}

func (msg *MsgMintcoin) Route() string {
	return RouterKey
}

func (msg *MsgMintcoin) Type() string {
	return TypeMsgMintcoin
}

func (msg *MsgMintcoin) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMintcoin) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintcoin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
