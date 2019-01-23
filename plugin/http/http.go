package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/andy-zhangtao/GoTrigger/pb/v1/plugin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var (
	port = flag.Int("port", 50000, "The server port")
)

const ModuleName = "GoTrigger-Http-Plugin"

type httpPlugin struct{}

func (h httpPlugin) Invoke(t context.Context, p *pb.Trigger) (*pb.Response, error) {

	logrus.WithFields(logrus.Fields{"id": p.Id, "endpoint": p.Endpoint, "ext": p.Ext}).Info(ModuleName)

	return &pb.Response{
		Success: true,
		Error:   "",
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	h := httpPlugin{}
	pb.RegisterHttpPluginServer(server, h)

	server.Serve(lis)
}
