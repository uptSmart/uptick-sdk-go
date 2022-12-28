package sdk

import (
	"github.com/tendermint/tendermint/libs/log"
	"github.com/uptsmart/uptick-sdk-go/erc721"

	"github.com/uptsmart/uptick-sdk-go/bank"
	"github.com/uptsmart/uptick-sdk-go/client"
	keys "github.com/uptsmart/uptick-sdk-go/client"
	commoncodec "github.com/uptsmart/uptick-sdk-go/common/codec"
	cryptotypes "github.com/uptsmart/uptick-sdk-go/common/codec/types"
	commoncryptocodec "github.com/uptsmart/uptick-sdk-go/common/crypto/codec"
	"github.com/uptsmart/uptick-sdk-go/types"
	txtypes "github.com/uptsmart/uptick-sdk-go/types/tx"
)

type Client struct {
	logger         log.Logger
	moduleManager  map[string]types.Module
	encodingConfig types.EncodingConfig
	types.BaseClient
	Bank   bank.Client
	Key    keys.Client
	Erc721 erc721.Client
}

func NewClient(cfg types.ClientConfig) Client {

	encodingConfig := makeEncodingConfig()

	// create a instance of baseClient
	baseClient := client.NewBaseClient(cfg, encodingConfig, nil)
	bankClient := bank.NewClient(baseClient, encodingConfig.Marshaler)
	erc721Client := erc721.NewClient(baseClient, encodingConfig.Marshaler)

	keysClient := keys.NewKeysClient(cfg, baseClient)

	client := Client{
		logger:         baseClient.Logger(),
		BaseClient:     baseClient,
		moduleManager:  make(map[string]types.Module),
		encodingConfig: encodingConfig,
		Bank:           bankClient,
		Key:            keysClient,
		Erc721:         erc721Client,
	}
	client.RegisterModule(
		bankClient,
		erc721Client,
	)

	return client
}

func (client *Client) SetLogger(logger log.Logger) {
	client.BaseClient.SetLogger(logger)
}

func (client *Client) Codec() *commoncodec.LegacyAmino {
	return client.encodingConfig.Amino
}

func (client *Client) AppCodec() commoncodec.Marshaler {
	return client.encodingConfig.Marshaler
}

func (client *Client) EncodingConfig() types.EncodingConfig {
	return client.encodingConfig
}

func (client *Client) Manager() types.BaseClient {
	return client.BaseClient
}

func (client *Client) RegisterModule(ms ...types.Module) {

	for _, m := range ms {
		m.RegisterInterfaceTypes(client.encodingConfig.InterfaceRegistry)
	}
}

func (client *Client) Module(name string) types.Module {
	return client.moduleManager[name]
}

func makeEncodingConfig() types.EncodingConfig {

	amino := commoncodec.NewLegacyAmino()
	interfaceRegistry := cryptotypes.NewInterfaceRegistry()
	marshaler := commoncodec.NewProtoCodec(interfaceRegistry)
	txCfg := txtypes.NewTxConfig(marshaler, txtypes.DefaultSignModes)

	encodingConfig := types.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          txCfg,
		Amino:             amino,
	}
	RegisterLegacyAminoCodec(encodingConfig.Amino)
	RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

// RegisterLegacyAminoCodec registers the sdk message type.
func RegisterLegacyAminoCodec(cdc *commoncodec.LegacyAmino) {

	cdc.RegisterInterface((*types.Msg)(nil), nil)
	cdc.RegisterInterface((*types.Tx)(nil), nil)
	commoncryptocodec.RegisterCrypto(cdc)
}

// RegisterInterfaces registers the sdk message type.
func RegisterInterfaces(registry cryptotypes.InterfaceRegistry) {

	registry.RegisterInterface("cosmos.v1beta1.Msg", (*types.Msg)(nil))
	txtypes.RegisterInterfaces(registry)
	commoncryptocodec.RegisterInterfaces(registry)
}
