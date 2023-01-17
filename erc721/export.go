package erc721

import sdk "github.com/uptsmart/uptick-sdk-go/types"

// Client expose erc721 module api for user
type Client interface {
	sdk.Module
	ConvertERC721(contractAddress string, tokenId string, receiver string, sender string, classId string, nftId string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	ConvertNFT(contractAddress string, tokenId string, receiver string, sender string, classId string, nftId string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	SubscribeSendTx(from, to string, callback EventMsgSendCallback) sdk.Subscription
	Send(to string, amount sdk.DecCoins, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	// MultiSend(receipts MultiSendRequest, baseTx sdk.BaseTx) ([]sdk.ResultTx, sdk.Error)
	// SendWitchSpecAccountInfo(to string, sequence, accountNumber uint64, amount sdk.DecCoins, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
}

//msg := &types.MsgConvertERC721{
//ContractAddress: contractAddress,
//TokenId:         tokenID,
//Receiver:        receiver.String(),
//Sender:          from.Hex(),
//ClassId:         classID,
//NftId:           nftID,
//}
// ERC721 -> cosmosNFT

type Receipt struct {
	Address string       `json:"address"`
	Amount  sdk.DecCoins `json:"amount"`
}
type MultiSendRequest struct {
	Receipts []Receipt
}

type EventDataMsgSend struct {
	Height int64      `json:"height"`
	Hash   string     `json:"hash"`
	From   string     `json:"from"`
	To     string     `json:"to"`
	Amount []sdk.Coin `json:"amount"`
}
type EventMsgSendCallback func(EventDataMsgSend)
