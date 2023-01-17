package erc721

import (
	"fmt"

	commoncodec "github.com/uptsmart/uptick-sdk-go/common/codec"
	"github.com/uptsmart/uptick-sdk-go/common/codec/types"
	sdk "github.com/uptsmart/uptick-sdk-go/types"
)

type erc721Client struct {
	sdk.BaseClient
	commoncodec.Marshaler
}

func (e erc721Client) Name() string {
	//TODO implement me
	panic("implement me")
}

func (e erc721Client) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

// bank NewClient
func NewClient(bc sdk.BaseClient, cdc commoncodec.Marshaler) Client {
	return erc721Client{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (e erc721Client) ConvertERC721(contractAddress string,
	tokenId string, receiver string, sender string,
	classId string, nftId string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {

	fmt.Println("xx 0")
	msg := NewMsgConvertErc721Send(contractAddress, tokenId, receiver, sender, classId, nftId)
	fmt.Printf("xx 1 %v \n", msg)

	return e.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (e erc721Client) ConvertNFT(contractAddress string,
	tokenId string, receiver string, sender string,
	classId string, nftId string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {

	fmt.Println("xx 0")
	msg := NewMsgConvertNFTSend(contractAddress, tokenId, receiver, sender, classId, nftId)
	fmt.Printf("xx 1 %v \n", msg)

	return e.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

// Send is responsible for transferring tokens from `From` to `to` account
func (b erc721Client) Send(to string, amount sdk.DecCoins, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := b.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrapf("%s not found", baseTx.From)
	}

	amt, err := b.ToMinCoin(amount...)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	outAddr, err := sdk.AccAddressFromBech32(to)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrapf(fmt.Sprintf("%s invalid address", to))
	}

	msg := NewMsgSend(sender, outAddr, amt)
	return b.BuildAndSend([]sdk.Msg{msg}, baseTx)
}
