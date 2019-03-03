package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/andy-zhangtao/GoTrigger/pb/v1/plugin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

var _VERSION_ string
var _BUILD_ string
var (
	address = flag.String("address", "0.0.0.0", "The server listen address")
	port    = flag.Int("port", 50000, "The server port")
)

const ModuleName = "GoTrigger-Http-Plugin"

type httpPlugin struct{}

// Invoke
// Implement the invoke logic
func (h httpPlugin) Invoke(t context.Context, p *pb.Trigger) (*pb.Response, error) {

	logrus.WithFields(logrus.Fields{"id": p.Id, "endpoint": p.Endpoint, "ext": p.Ext}).Info(ModuleName)

	method := http.MethodGet

	if m, ok := p.Ext["method"]; ok {
		method = m
	}

	err := send(method, p.Endpoint, nil, nil)
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

func send(method, endpoint string, query map[string]string, header map[string]string) (error) {

	_query := ""
	for key, value := range query {
		_query += fmt.Sprintf("%s=%s&", key, value)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, endpoint, strings.NewReader(_query))
	if err != nil {
		return err
	}

	for key, value := range header {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	logrus.WithFields(logrus.Fields{"VERSION": _VERSION_, "BUILD": _BUILD_}).Info(ModuleName)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *address, *port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	h := httpPlugin{}
	pb.RegisterHttpPluginServer(server, h)

	server.Serve(lis)
}
