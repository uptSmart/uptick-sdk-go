package integration_test

func (s IntegrationTestSuite) TestErc721() {

	cases := []SubTest{
		{
			"TestQueryAccount",
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

	//

}
