package erc721

import (
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

	msg := NewMsgConvertErc721Send(contractAddress, tokenId, receiver, sender, classId, nftId)

	return e.BuildAndSend([]sdk.Msg{msg}, baseTx)
}
