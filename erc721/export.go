package erc721

import sdk "github.com/uptsmart/uptick-sdk-go/types"

// Client expose erc721 module api for user
type Client interface {
	sdk.Module
	ConvertERC721(contractAddress string, tokenId string, receiver string, sender string, classId string, nftId string, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
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
