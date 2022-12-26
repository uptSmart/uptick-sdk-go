package client

import (
	"context"
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	cache "github.com/irisnet/core-sdk-go/common/cache"
	commoncodec "github.com/irisnet/core-sdk-go/common/codec"
	ethXTypes "github.com/irisnet/core-sdk-go/ethermint/evm/types"
	// ethTypes "github.com/irisnet/core-sdk-go/ethermint/types"
	sdk "github.com/irisnet/core-sdk-go/types"
	// "github.com/cosmos/cosmos-sdk/x/bank"
)

// Must be used with locker, otherwise there are thread safety issues
type AccountQuery struct {
	sdk.Queries
	sdk.GRPCClient
	log.Logger
	cache.Cache
	cdc        commoncodec.Marshaler
	Km         sdk.KeyManager
	expiration time.Duration
}

func (a AccountQuery) QueryAndRefreshAccount(address string) (sdk.BaseAccount, sdk.Error) {
	account, err := a.Get(a.prefixKey(address))
	if err != nil {
		return a.refresh(address)
	}

	acc := account.(accountInfo)
	baseAcc := sdk.BaseAccount{
		Address:       address,
		AccountNumber: acc.N,
		Sequence:      acc.S + 1,
	}
	a.saveAccount(baseAcc)

	a.Debug("query account from cache", "address", address)
	return baseAcc, nil
}

func (a AccountQuery) QueryAccount(address string) (sdk.BaseAccount, sdk.Error) {
	conn, err := a.GenConn()

	fmt.Printf("---xxl 0 \n")
	if err != nil {
		return sdk.BaseAccount{}, sdk.Wrap(err)
	}

	//request := &auth.QueryAccountRequest{
	//	Address: address,
	//}
	fmt.Printf("---xxl 1 \n")
	request := ethXTypes.QueryCosmosAccountRequest{
		Address: "0x7c4663d780EfA75dAF623a4E79D505df8be88CDC",
	}

	fmt.Printf("---xxl 2 \n")
	response, err := ethXTypes.NewQueryClient(conn).CosmosAccount(context.Background(), &request)
	if err != nil {
		fmt.Printf("---xxl 3 %v \n", err)

		return sdk.BaseAccount{}, sdk.Wrap(err)
	}
	fmt.Printf("---xxl response %v \n", response)

	// bank.NewQueryClient(conn).AllBalances(context.Background(), breq)

	return sdk.BaseAccount{}, nil

	//var baseAccount auth.Account
	//var ethAccount ethTypes.EthAccount
	//if err := a.cdc.UnpackAny(response., &ethAccount); err != nil {
	//
	//	fmt.Printf("---xxl err %v \n", err)
	//	return sdk.BaseAccount{}, sdk.Wrap(err)
	//}
	//fmt.Printf("---xxl ethAccount %v \n", ethAccount)
	//
	//account := baseAccount.(*auth.BaseAccount).ConvertAccount(a.cdc).(sdk.BaseAccount)
	//
	//breq := &bank.QueryAllBalancesRequest{
	//	Address:    address,
	//	Pagination: nil,
	//}
	//balances, err := bank.NewQueryClient(conn).AllBalances(context.Background(), breq)
	//if err != nil {
	//	return sdk.BaseAccount{}, sdk.Wrap(err)
	//}
	//
	//account.Coins = balances.Balances
	//return account, nil
}

func (a AccountQuery) QueryAddress(name, password string) (sdk.AccAddress, sdk.Error) {
	addr, err := a.Get(a.prefixKey(name))
	if err == nil {
		address, err := sdk.AccAddressFromBech32(addr.(string))
		if err != nil {
			a.Debug("invalid address", "name", name)
			_ = a.Remove(a.prefixKey(name))
		} else {
			return address, nil
		}
	}

	_, address, err := a.Km.Find(name, password)
	if err != nil {
		a.Debug("can't find account", "name", name)
		return address, sdk.Wrap(err)
	}

	if err := a.SetWithExpire(a.prefixKey(name), address.String(), a.expiration); err != nil {
		a.Debug("cache user failed", "name", name)
	}
	a.Debug("query user from cache", "name", name, "address", address.String())
	return address, nil
}

func (a AccountQuery) removeCache(address string) bool {
	return a.Remove(a.prefixKey(address))
}

func (a AccountQuery) refresh(address string) (sdk.BaseAccount, sdk.Error) {
	account, err := a.QueryAccount(address)
	if err != nil {
		a.Error("update cache failed", "address", address, "errMsg", err.Error())
		return sdk.BaseAccount{}, sdk.Wrap(err)
	}

	a.saveAccount(account)
	return account, nil
}

func (a AccountQuery) saveAccount(account sdk.BaseAccount) {
	address := account.Address
	info := accountInfo{
		N: account.AccountNumber,
		S: account.Sequence,
	}
	if err := a.SetWithExpire(a.prefixKey(address), info, a.expiration); err != nil {
		a.Debug("cache user failed", "address", account.Address)
		return
	}
	a.Debug("cache account", "address", address, "expiration", a.expiration.String())
}

func (a AccountQuery) prefixKey(address string) string {
	return fmt.Sprintf("account:%s", address)
}

type accountInfo struct {
	N uint64 `json:"n"`
	S uint64 `json:"s"`
}
