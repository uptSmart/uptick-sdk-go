package erc721

import (
	commoncodec "github.com/uptsmart/uptick-sdk-go/common/codec"
	"github.com/uptsmart/uptick-sdk-go/common/codec/types"
	commoncryptocodec "github.com/uptsmart/uptick-sdk-go/common/crypto/codec"
	sdk "github.com/uptsmart/uptick-sdk-go/types"
)

var (
	amino     = commoncodec.NewLegacyAmino()
	ModuleCdc = commoncodec.NewAminoCodec(amino)
)

func init() {
	commoncryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// No duplicate registration
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgConvertERC721{},
	)

}
