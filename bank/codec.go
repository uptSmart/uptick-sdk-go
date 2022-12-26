package bank

import (
	commoncodec "github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	commoncryptocodec "github.com/irisnet/core-sdk-go/common/crypto/codec"
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/auth"

	ethTypes "github.com/irisnet/core-sdk-go/ethermint/types"
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
