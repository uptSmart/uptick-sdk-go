package integration_test

import (
	"fmt"
	"github.com/uptsmart/uptick-sdk-go/types"
	"time"
)

func (s IntegrationTestSuite) TestErc721() {

	cases := []SubTest{
		{
			"TestConvertErc721",
			convertErc721,
		},
		//{
		//	"TestSend",
		//	send,
		//},
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
