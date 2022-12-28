package erc721

import (
	sdk "github.com/uptsmart/uptick-sdk-go/types"
)

const (
	ModuleName = "erc721"

	TypeMsgSend      = "send"
	TypeMsgMultiSend = "multisend"
)

var _ sdk.Msg = &MsgConvertERC721{}

// NewMsgConvertErc721Send - construct a msg to send coins from one account to another.
func NewMsgConvertErc721Send(
	contractAddress string, tokenId string, receiver string,
	sender string, classId string, nftId string) *MsgConvertERC721 {
	return &MsgConvertERC721{
		ContractAddress: contractAddress,
		TokenId:         tokenId,
		Receiver:        receiver,
		Sender:          sender,
		ClassId:         classId,
		NftId:           nftId,
	}
}

// Route Implements Msg.
func (msg MsgConvertERC721) Route() string { return ModuleName }

func (m *MsgConvertERC721) Type() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgConvertERC721) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m *MsgConvertERC721) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *MsgConvertERC721) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}
