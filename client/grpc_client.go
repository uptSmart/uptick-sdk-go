package client

import (
	grpc1 "github.com/gogo/protobuf/grpc"
	log "github.com/sirupsen/logrus"
	"github.com/uptsmart/uptick-sdk-go/types"
	"google.golang.org/grpc"
)

type grpcClient struct {
	clientConn grpc1.ClientConn
}

func NewGRPCClient(url string, options ...grpc.DialOption) types.GRPCClient {

	if options == nil {
		options = []grpc.DialOption{grpc.WithInsecure()}
	}

	clientConn, err := grpc.Dial(url, options...)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
	return &grpcClient{clientConn: grpc1.ClientConn(clientConn)}
}

func (g grpcClient) GenConn() (grpc1.ClientConn, error) {
	return g.clientConn, nil
}
