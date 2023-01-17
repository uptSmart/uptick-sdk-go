package integration_test

import (
	"fmt"
	"time"

	"github.com/uptsmart/uptick-sdk-go/erc721"
	"github.com/uptsmart/uptick-sdk-go/types"
)

func (s IntegrationTestSuite) TestErc721() {

	cases := []SubTest{
		// {
		// 	"TestConvertErc721",
		// 	convertErc721,
		// },
		// {
		// 	"TestconvertNFT",
		// 	convertNFT,
		// },
		{
			"TestSend",
			send721,
		},
		//{
		//	"TestMultiSend",
		//	multiSend,
		//},
		//{
		//	"TestSimulate",
		//	simulate,
		//},
		//{
		//	"TestSendWitchSpecAccountInfo",
		//	sendWitchSpecAccountInfo,
		//},
	}

	for _, t := range cases {
		s.Run(t.testName, func() { t.testCase(s) })
	}
}

func convertErc721(s IntegrationTestSuite) {

	baseTx := types.BaseTx{
		From:               s.Account().Name,
		Gas:                40000837,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           s.Account().Password,
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}

	res, err := s.Erc721.ConvertERC721(
		"0x432FDd12f13F3D03a8429687db3102e65dBE3060",
		"1",
		s.Account().Address.String(),
		"0x7c4663d780EfA75dAF623a4E79D505df8be88CDC",
		"uptick-0x432FDd12f13F3D03a8429687db3102e65dBE3060",
		"uptick1",
		baseTx)
	fmt.Printf("xxl err %v-%v", err, res.Hash)
	s.NoError(err)
	s.NotEmpty(res.Hash)
	time.Sleep(1 * time.Second)

}

func convertNFT(s IntegrationTestSuite) {
	baseTx := types.BaseTx{
		From:               s.Account().Name,
		Gas:                40000837,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           s.Account().Password,
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}

	res, err := s.Erc721.ConvertNFT(
		"0x432FDd12f13F3D03a8429687db3102e65dBE3060",
		"1",
		"0x7c4663d780EfA75dAF623a4E79D505df8be88CDC",
		s.Account().Address.String(),
		"uptick-0x432FDd12f13F3D03a8429687db3102e65dBE3060",
		"uptick1",
		baseTx)
	fmt.Printf("xxl err %v-%v", err, res.Hash)
	s.NoError(err)
	s.NotEmpty(res.Hash)
	time.Sleep(1 * time.Second)
}

// ///
func send721(s IntegrationTestSuite) {
	coins, err := types.ParseDecCoins("10iris")
	s.NoError(err)
	to := s.GetRandAccount().Address.String()
	fmt.Println("111111111113")
	fmt.Println(to)
	fmt.Println(s.Account())
	fmt.Println("222222222222")

	ch := make(chan int)
	s.Erc721.SubscribeSendTx(s.Account().Address.String(), to, func(send erc721.EventDataMsgSend) {
		ch <- 1
	})

	baseTx := types.BaseTx{
		From:               s.Account().Name,
		Gas:                200000,
		Memo:               "TEST",
		Mode:               types.Commit,
		Password:           s.Account().Password,
		SimulateAndExecute: false,
		GasAdjustment:      1.5,
	}

	res, err := s.Erc721.Send(to, coins, baseTx)
	s.NoError(err)
	s.NotEmpty(res.Hash)
	time.Sleep(1 * time.Second)

	resp, err := s.Manager().QueryTx(res.Hash)
	s.NoError(err)
	s.Equal(resp.Result.Code, uint32(0))
	s.Equal(resp.Height, res.Height)

	<-ch
}

// func multiSend(s IntegrationTestSuite) {
// 	baseTx := types.BaseTx{
// 		From:     s.Account().Name,
// 		Gas:      2000000,
// 		Memo:     "test",
// 		Mode:     types.Commit,
// 		Password: s.Account().Password,
// 	}

// 	coins, err := types.ParseDecCoins("1000iris")
// 	s.NoError(err)

// 	accNum := 11
// 	acc := make([]string, accNum)
// 	receipts := make([]bank.Receipt, accNum)
// 	for i := 0; i < accNum; i++ {
// 		acc[i] = s.RandStringOfLength(10)
// 		addr, _, err := s.Add(acc[i], "1234567890")

// 		s.NoError(err)
// 		s.NotEmpty(addr)

// 		receipts[i] = bank.Receipt{
// 			Address: addr,
// 			Amount:  coins,
// 		}
// 	}

// 	_, err = s.Bank.MultiSend(bank.MultiSendRequest{Receipts: receipts}, baseTx)
// 	s.NoError(err)

// 	coins, err = types.ParseDecCoins("1iris")
// 	s.NoError(err)
// 	to := s.GetRandAccount().Address.String()
// 	begin := time.Now()
// 	var wait sync.WaitGroup
// 	for i := 1; i < 5; i++ {
// 		wait.Add(1)
// 		index := rand.Intn(accNum)
// 		go func() {
// 			defer wait.Done()
// 			_, err := s.Bank.Send(to, coins, types.BaseTx{
// 				From:     acc[index],
// 				Gas:      200000,
// 				Memo:     "test",
// 				Mode:     types.Commit,
// 				Password: "1234567890",
// 			})
// 			s.NoError(err)
// 		}()
// 	}
// 	wait.Wait()
// 	end := time.Now()
// 	fmt.Printf("total senconds:%s\n", end.Sub(begin).String())
// }

// func simulate(s IntegrationTestSuite) {
// 	coins, err := types.ParseDecCoins("10iris")
// 	s.NoError(err)
// 	to := s.GetRandAccount().Address.String()
// 	baseTx := types.BaseTx{
// 		From:               s.Account().Name,
// 		Password:           s.Account().Password,
// 		Gas:                200000,
// 		Memo:               "test",
// 		Mode:               types.Commit,
// 		SimulateAndExecute: true,
// 	}

// 	result, err := s.Bank.Send(to, coins, baseTx)
// 	s.NoError(err)
// 	s.Greater(result.GasWanted, int64(0))
// 	fmt.Println(result)
// }

// func sendWitchSpecAccountInfo(s IntegrationTestSuite) {
// 	for i := 0; i < 10; i++ {
// 		coins, err := types.ParseDecCoins("10iris")
// 		baseTx := types.BaseTx{
// 			From:     s.Account().Name,
// 			Gas:      200000,
// 			Fee:      coins,
// 			Memo:     "TEST",
// 			Mode:     types.Commit,
// 			Password: s.Account().Password,
// 		}

// 		curAccount, err := s.Bank.QueryAccount(s.Account().Address.String())
// 		require.NoError(s.T(), err)

// 		accountNumber := curAccount.AccountNumber
// 		sequence := curAccount.Sequence
// 		randomAddr := s.GetRandAccount().Address.String()
// 		amount, _ := types.ParseDecCoins("10iris")

// 		res, err := s.Bank.SendWitchSpecAccountInfo(randomAddr, sequence, accountNumber, amount, baseTx)
// 		require.NoError(s.T(), err)
// 		require.NotEmpty(s.T(), res.Hash)
// 	}
// }
