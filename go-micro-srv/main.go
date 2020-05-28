//__author__ = "YaoYao"
//Date: 2020/5/4
package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/client"
	grpcc "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/server"
	grpcs "github.com/micro/go-micro/v2/server/grpc"
	"github.com/micro/go-plugins/broker/kafka/v2"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	proto "k8s-micro/go-micro-srv/proto"
	"net/http"
	_ "net/http/pprof"
)

type Greeter struct {
}

func (g Greeter) Hello(ctx context.Context, request *proto.HelloRequest, response *proto.HelloResponse) error {
	panic("implement me")
}

var (
	DefaultServerPort = ":9100"
)

func main() {
	service := micro.NewService(
		micro.Name("go-micro-srv"),
		micro.Server(makeMicroRPCServer()),
		micro.Client(makeMicroRPCClient()),
		micro.Broker(makeBroker()),
		micro.Registry(makeMicroRegistry()),
	)
	service.Init()

	registryRPCHandler(service.Server())

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

	go func() { //pprof
		if err := http.ListenAndServe("", nil); err != nil {

		}
	}()
}

func makeMicroRegistry() registry.Registry {
	return kubernetes.NewRegistry()
}

func makeMicroRPCServer() server.Server {
	return grpcs.NewServer(server.Address(DefaultServerPort))
}

func makeMicroRPCClient() client.Client {
	return grpcc.NewClient()
}

func makeBroker() broker.Broker {
	return kafka.NewBroker()
}

func registryRPCHandler(s server.Server) {
	_ = proto.RegisterGreeterHandler(s, new(Greeter))
}