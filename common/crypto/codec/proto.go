package codec

import (
	tmcrypto "github.com/tendermint/tendermint/crypto"
	codectypes "github.com/uptsmart/uptick-sdk-go/common/codec/types"
	"github.com/uptsmart/uptick-sdk-go/common/crypto/keys/ed25519"
	"github.com/uptsmart/uptick-sdk-go/common/crypto/keys/multisig"
	"github.com/uptsmart/uptick-sdk-go/common/crypto/keys/secp256k1"
	cryptotypes "github.com/uptsmart/uptick-sdk-go/common/crypto/types"

	"github.com/uptsmart/uptick-sdk-go/common/crypto/keys/eth_secp256k1"
)

// RegisterInterfaces registers the sdk.Tx interface.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	// TODO We now register both Tendermint's PubKey and our own PubKey. In the
	// long-term, we should move away from Tendermint's PubKey, and delete
	// these lines.
	registry.RegisterInterface("tendermint.crypto.Pubkey", (*tmcrypto.PubKey)(nil))
	registry.RegisterImplementations((*tmcrypto.PubKey)(nil), &ed25519.PubKey{})
	registry.RegisterImplementations((*tmcrypto.PubKey)(nil), &secp256k1.PubKey{})
	registry.RegisterImplementations((*tmcrypto.PubKey)(nil), &multisig.LegacyAminoPubKey{})

	registry.RegisterImplementations((*tmcrypto.PubKey)(nil), &ethsecp256k1.PubKey{})

	registry.RegisterInterface("cosmos.crypto.Pubkey", (*cryptotypes.PubKey)(nil))
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ed25519.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &secp256k1.PubKey{})
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &multisig.LegacyAminoPubKey{})

	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})
}
