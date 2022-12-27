package integration_test

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/uptsmart/uptick-sdk-go"
	"github.com/uptsmart/uptick-sdk-go/common/crypto"
	"github.com/uptsmart/uptick-sdk-go/common/log"
	"github.com/uptsmart/uptick-sdk-go/types"
	"github.com/uptsmart/uptick-sdk-go/types/store"

	ethTypes "github.com/uptsmart/uptick-sdk-go/ethermint/types"
)

const (
	nodeURI  = "tcp://localhost:26657"
	grpcAddr = "localhost:9090"
	chainID  = "uptick_7000-1"
	charset  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	addr     = "uptick103rx84uqa7n4mtmz8f88n4g9m7973rxutrtn7d"
)

type IntegrationTestSuite struct {
	suite.Suite
	sdk.Client
	r            *rand.Rand
	rootAccount  MockAccount
	randAccounts []MockAccount
}

type SubTest struct {
	testName string
	testCase func(s IntegrationTestSuite)
}

// MockAccount define a account for test
type MockAccount struct {
	Name, Password string
	Address        types.AccAddress
	test           types.BaseAccount
	tes2           ethTypes.EthAccount
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {

	bech32AddressPrefix := types.AddrPrefixCfg{
		AccountAddr:   "uptick",
		ValidatorAddr: "uptickvaloper",
		ConsensusAddr: "uptickvalcons",
		AccountPub:    "uptickpub",
		ValidatorPub:  "uptickvaloperpub",
		ConsensusPub:  "uptickvalconspub",
	}

	options := []types.Option{
		types.KeyDAOOption(store.NewMemory(nil)),
		types.TimeoutOption(10),
		types.TokenManagerOption(TokenManager{}),
		types.KeyManagerOption(crypto.NewKeyManager()),
		types.Bech32AddressPrefixOption(&bech32AddressPrefix),
		types.BIP44PathOption(""),
	}

	cfg, err := types.NewClientConfig(nodeURI, grpcAddr, chainID, options...)
	if err != nil {
		panic(err)
	}

	s.Client = sdk.NewClient(cfg)

	s.r = rand.New(rand.NewSource(time.Now().UnixNano()))

	s.rootAccount = MockAccount{
		Name:     "val1",
		Password: "1234567890",
		Address:  types.MustAccAddressFromBech32(addr),
	}
	s.SetLogger(log.NewLogger(log.Config{
		Format: log.FormatJSON,
		Level:  log.DebugLevel,
	}))
	s.initAccount()
}

func (s *IntegrationTestSuite) initAccount() {

	_, err := s.Import(
		s.Account().Name,
		s.Account().Password,
		string(getPrivKeyArmor()),
	)
	if err != nil {
		panic(err)
	}

	//var receipts bank.Receipts
	for i := 0; i < 5; i++ {
		name := s.RandStringOfLength(10)
		pwd := s.RandStringOfLength(16)
		address, _, err := s.Add(name, "11111111")
		if err != nil {
			panic("generate test account failed")
		}

		s.randAccounts = append(s.randAccounts, MockAccount{
			Name:     name,
			Password: pwd,
			Address:  types.MustAccAddressFromBech32(address),
		})
	}

}

// RandStringOfLength return a random string
func (s *IntegrationTestSuite) RandStringOfLength(l int) string {
	var result []byte
	bytes := []byte(charset)
	for i := 0; i < l; i++ {
		result = append(result, bytes[s.r.Intn(len(bytes))])
	}
	return string(result)
}

// GetRandAccount return a random test account
func (s *IntegrationTestSuite) GetRandAccount() MockAccount {
	return s.randAccounts[s.r.Intn(len(s.randAccounts))]
}

// Account return a test account
func (s *IntegrationTestSuite) Account() MockAccount {
	return s.rootAccount
}

func getPrivKeyArmor() []byte {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path = filepath.Dir(path)
	path = filepath.Join(path, "integration_test/scripts/priv.key")
	bz, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bz
}

type TokenManager struct{}

////
func (TokenManager TokenManager) QueryToken(denom string) (types.Token, error) {
	return types.Token{}, nil
}

func (TokenManager TokenManager) SaveTokens(tokens ...types.Token) {
	return
}

func (TokenManager TokenManager) ToMinCoin(coins ...types.DecCoin) (types.Coins, types.Error) {
	for i := range coins {
		if coins[i].Denom == "uptick" {
			coins[i].Denom = "auptick"
			coins[i].Amount = coins[i].Amount.MulInt(types.NewIntWithDecimal(1, 6))
		}
	}
	ucoins, _ := types.DecCoins(coins).TruncateDecimal()
	return ucoins, nil
}

func (TokenManager TokenManager) ToMainCoin(coins ...types.Coin) (types.DecCoins, types.Error) {
	decCoins := make(types.DecCoins, len(coins), 0)
	for _, coin := range coins {
		if coin.Denom == "auptick" {
			amtount := types.NewDecFromInt(coin.Amount).Mul(types.NewDecWithPrec(1, 6))
			decCoins = append(decCoins, types.NewDecCoinFromDec("uptick", amtount))
		}
	}
	return decCoins, nil
}
