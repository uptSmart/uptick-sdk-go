package bank

import (
	commoncodec "github.com/uptsmart/uptick-sdk-go/common/codec"
	"github.com/uptsmart/uptick-sdk-go/common/codec/types"
	commoncryptocodec "github.com/uptsmart/uptick-sdk-go/common/crypto/codec"
	sdk "github.com/uptsmart/uptick-sdk-go/types"
	"github.com/uptsmart/uptick-sdk-go/types/auth"

	ethTypes "github.com/uptsmart/uptick-sdk-go/ethermint/types"
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
		&MsgSend{},
		&MsgMultiSend{},
	)

	//registry.RegisterImplementations(
	//	(*auth.Account)(nil),
	//	&auth.BaseAccount{},
	//)

	registry.RegisterImplementations(
		(*auth.Account)(nil),
		&ethTypes.EthAccount{},
	)

}
