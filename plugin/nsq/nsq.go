package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/andy-zhangtao/GoTrigger/pb/v1/plugin"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var _VERSION_ string
var _BUILD_ string
var (
	address = flag.String("address", "0.0.0.0", "The server listen address")
	port    = flag.Int("port", 50000, "The server port")
)

const ModuleName = "GoTrigger-NSQ-Plugin"

type nsqPlugin struct{}

// Invoke
// Implement the invoke logic
func (h nsqPlugin) Invoke(t context.Context, p *pb.Trigger) (*pb.Response, error) {

	logrus.WithFields(logrus.Fields{"id": p.Id, "endpoint": p.Endpoint, "ext": p.Ext}).Info(ModuleName)

	var topic string

	var message string
	if m, ok := p.Ext["topic"]; ok {
		topic = m
	}

	if m, ok := p.Ext["message"]; ok {
		message = m
	}

	err := send(p.Endpoint, topic, message)
	if err != nil {
		return &pb.Response{
			Success: false,
			Error:   err.Error(),
		}, nil
	}

	return &pb.Response{
		Success: true,
		Error:   "",
	}, nil
}

func send(endpoint, topic, message string) (err error) {

	producer, err := nsq.NewProducer(endpoint, nsq.NewConfig())
	if err != nil {
		return
	}

	return producer.Publish(topic, []byte(message))
}

func main() {
	flag.Parse()
	logrus.WithFields(logrus.Fields{"VERSION": _VERSION_, "BUILD": _BUILD_}).Info(ModuleName)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *address, *port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	n := nsqPlugin{}
	pb.RegisterHttpPluginServer(server, n)

	server.Serve(lis)
}
