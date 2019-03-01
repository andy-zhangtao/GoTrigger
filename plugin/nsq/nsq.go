package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"regexp"
	"strings"

	pb "github.com/andy-zhangtao/GoTrigger/pb/v1/plugin"
	"github.com/andy-zhangtao/gogather/time"
	"github.com/nsqio/go-nsq"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var _VERSION_ string
var _BUILD_ string
var (
	address = flag.String("address", "0.0.0.0", "The server listen address")
	port    = flag.Int("port", 50000, "The server port")
)

const ModuleName = "GoTrigger-NSQ-Plugin"
const PREDATETIME = "formate:"
const DEFAULTTIME = "YYYYMMDDThh:mm:ss"

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
		if strings.TrimSpace(m) != "" {
			message = customValue(strings.TrimSpace(m))
		}
	} else {
		// give me default date (YYYYMMDDThh:mm:ss)
		message = customValue(DEFAULTTIME)
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

//customValue
//create user custom value. If err, then return a now time(YYYYMMDDThh:mm:ss)
func customValue(custom string) (value string) {
	//判断是否为日志格式值
	zt := time.Ztime{}
	date, _ := zt.Now().Format("YYYYMMDDThh:mm:ss")

	if strings.HasPrefix(custom, PREDATETIME) {
		f := strings.Split(custom, PREDATETIME)
		if len(f) == 1 {
			return date
		}

		//生成特定格式的时间
		re := regexp.MustCompile(`^\$\{(.*)\}$`)
		sub := re.FindStringSubmatch(f[1])
		logrus.WithFields(logrus.Fields{"Custom-Format": sub[1]}).Debug(ModuleName)
		_date, err := zt.Now().Format(sub[1])
		if err != nil {
			return date
		}

		return _date
	}

	if custom == DEFAULTTIME {
		return date
	}

	return custom
}

func send(endpoint, topic, message string) (err error) {

	producer, err := nsq.NewProducer(endpoint, nsq.NewConfig())
	if err != nil {
		return
	}

	return producer.Publish(topic, []byte(message))
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
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
