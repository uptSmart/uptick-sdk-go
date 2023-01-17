package erc721

import (
	github_com_uptsmart_uptick_sdk_go_types "github.com/uptsmart/uptick-sdk-go/types"
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

func NewMsgConvertNFTSend(
	contractAddress string, tokenId string, receiver string,
	sender string, classId string, nftId string) *MsgConvertNFT {
	return &MsgConvertNFT{
		ContractAddress: contractAddress,
		TokenId:         tokenId,
		Receiver:        receiver,
		Sender:          sender,
		ClassId:         classId,
		NftId:           nftId,
	}
}

// Route Implements Msg.
func (msg MsgConvertNFT) Route() string { return ModuleName }

func (m *MsgConvertNFT) Type() string {
	//TODO implement me
	panic("implement me")
}

func (m *MsgConvertNFT) ValidateBasic() error {
	//TODO implement me
	panic("implement me")
}

func (m *MsgConvertNFT) GetSignBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *MsgConvertNFT) GetSigners() []sdk.AccAddress {
	//TODO implement me
	panic("implement me")
}

// NewMsgSend - construct a msg to send coins from one account to another.
//
//nolint:interfacer
func NewMsgSend(fromAddr, toAddr sdk.AccAddress, amount sdk.Coins) *MsgSend {
	return &MsgSend{
		FromAddress: fromAddr.String(),
		ToAddress:   toAddr.String(),
		Amount:      amount,
	}
}

// MsgSend represents a message to send coins from one account to another.
type MsgSend struct {
	FromAddress string                                        `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	ToAddress   string                                        `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	Amount      github_com_uptsmart_uptick_sdk_go_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/uptsmart/uptick-sdk-go/types.Coins" json:"amount"`
}
