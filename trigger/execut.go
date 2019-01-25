package trigger

import (
	"context"
	"errors"
	"github.com/andy-zhangtao/GoTrigger/db"
	"github.com/andy-zhangtao/GoTrigger/model"
	pb "github.com/andy-zhangtao/GoTrigger/pb/v1/plugin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

//execut
//invoke trigger plugin
//If it is a sync type, exeut will waiting for finish. If it is a async, execut will quit.
func execut(t *model.Trigger) (err error) {
	logrus.WithFields(logrus.Fields{"name": t.Name, "parallel": t.Parallel, "nextime": t.NextTime, "async": t.Async}).Info(model.MODULENAME)

	if t.Type.Endpoint == "" {
		err = errors.New("GT must has A trigger plugin. ")
		return
	}

	ptr := model.TriggerPlugin{
		PID: t.Type.Kind,
	}

	if err := db.FindSpecifyTriggerPlugin(&ptr); err != nil {
		return err
	}

	if _, err := invokeHttp(t, ptr); err != nil {
		return err
	}

	return nil
}

func invokeHttp(t *model.Trigger, ptr model.TriggerPlugin) (succ bool, err error) {

	conn, err := grpc.Dial(ptr.Endpoint, grpc.WithInsecure())
	if err != nil {
		return false, err
	}

	defer conn.Close()

	client := pb.NewHttpPluginClient(conn)
	response, err := client.Invoke(context.Background(), &pb.Trigger{
		Id:       t.ID.Hex(),
		Endpoint: t.Type.Endpoint,
		Ext:      t.Type.Ext,
	})
	if err != nil {
		return false, err
	}

	if response.Error != "" {
		return false, errors.New(response.Error)
	}

	return true, nil
}
