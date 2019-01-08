/*
 * Copyright (c) 2018.
 * andy-zhangtao <ztao8607@gmail.com>
 */

package bwidow

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"time"

	"github.com/andy-zhangtao/gogather/zReflect"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/pelletier/go-toml"
)

//Write by zhangtao<ztao8607@gmail.com> . In 2018/5/23.

type BWMongoConf struct {
	Endpoint string
	DB       string
	User     string
	Password string
}

type BWMongo struct {
	session  *mgo.Session
	client   *mgo.Session
	db       *mgo.Database
	tableMap map[string]string
}

var conf *BWMongoConf

func (this *BWMongo) Map(u interface{}, name string) {
	if this.tableMap == nil {
		this.tableMap = make(map[string]string)
	}

	this.tableMap[reflect.TypeOf(u).Name()] = name
}

func (this *BWMongo) setDB() {
	this.db = this.session.Clone().DB(conf.DB)
}

func (this *BWMongo) first(u interface{}) (err error) {
	this.setDB()
	defer this.db.Session.Close()

	return this.db.C(this.tableMap[getTypeName(u)]).Find(nil).One(u)
}

func (this *BWMongo) findOne(u interface{}) (err error) {
	this.setDB()
	defer this.db.Session.Close()

	m := zReflect.ReflectStructInfo(u)

	return this.db.C(this.tableMap[getTypeName(u)]).Find(bson.M(m)).One(u)
}

func (this *BWMongo) findAll(u interface{}, a interface{}) (err error) {
	this.setDB()
	defer this.db.Session.Close()

	m := zReflect.ReflectStructInfo(u)

	return this.db.C(this.tableMap[getTypeName(u)]).Find(bson.M(m)).All(a)
}

func (this *BWMongo) findAllWithSort(u interface{}, a interface{}, sortField []string) (err error) {
	this.setDB()
	defer this.db.Session.Close()

	m := zReflect.ReflectStructInfo(u)

	return this.db.C(this.tableMap[getTypeName(u)]).Find(bson.M(m)).Sort(sortField...).All(a)
}

func (this *BWMongo) save(u interface{}) (err error) {
	this.setDB()
	defer this.db.Session.Close()
	return this.db.C(this.tableMap[getTypeName(u)]).Insert(u)
}

func (this *BWMongo) saveAll(u []interface{}) (err error) {
	typeName := reflect.TypeOf(u[0]).Name()
	this.setDB()
	defer this.db.Session.Close()

	bulk := this.db.C(this.tableMap[typeName]).Bulk()
	bulk.Insert(u...)
	_, err = bulk.Run()
	return
}

func (this *BWMongo) update(uPtr interface{}, field []string) (num int, err error) {
	this.setDB()
	defer this.db.Session.Close()

	m := zReflect.ReflectStructInfo(uPtr)

	nm := make(map[string]interface{})

	for _, f := range field {
		nm[f] = m[f]
	}

	info, err := this.db.C(this.tableMap[getTypeName(uPtr)]).UpdateAll(bson.M(nm), bson.M{"$set": bson.M(m)})
	return info.Updated, err
}

func (this *BWMongo) delete(uPtr interface{}, field []string) (num int, err error) {
	this.setDB()
	defer this.db.Session.Close()

	m := zReflect.ReflectStructInfo(uPtr)

	nm := make(map[string]interface{})

	for _, f := range field {
		nm[f] = m[f]
	}

	info, err := this.db.C(this.tableMap[getTypeName(uPtr)]).RemoveAll(bson.M(nm))

	return info.Removed, err
}

func (this *BWMongo) deleteAll(uPtr interface{}) (int, error) {
	this.setDB()
	defer this.db.Session.Close()

	info, err := this.db.C(this.tableMap[getTypeName(uPtr)]).RemoveAll(nil)

	return info.Removed, err
}

func (this *BWMongo) DriverInit() (err error) {
	if err = this.Check(); err != nil {
		return
	}

	conf = new(BWMongoConf)
	conf.User = os.Getenv(BW_MONGO_USER)
	conf.Password = os.Getenv(BW_MONGO_PASSWD)
	conf.Endpoint = os.Getenv(BW_MONGO_ENDPOINT)
	conf.DB = os.Getenv(BW_MONGO_DB)

	if conf.Endpoint == "" {
		useToml := true
		//	读取toml
		data, err := ioutil.ReadFile("bwidow_mongo.toml")
		if err != nil || len(data) == 0 {
			useToml = false
			//尝试json
			data, err = ioutil.ReadFile("bwidow_mongo.json")
			if err != nil || len(data) == 0 {
				return errors.New(fmt.Sprintf("Parse Configure Error %s", err.Error()))
			}
		}

		if useToml {
			err = toml.Unmarshal(data, conf)
			if err != nil {
				return err
			}
		} else {
			err = json.Unmarshal(data, conf)
			if err != nil {
				return err
			}
		}

	}

	var session *mgo.Session
	if conf.User == "" {
		session, err = mgo.Dial(conf.Endpoint)
		if err != nil {
			return
		}
	} else {
		dialInfo := &mgo.DialInfo{
			Addrs:    []string{conf.Endpoint},
			Database: conf.DB,
			Username: conf.User,
			Password: conf.Password,
			Timeout:  10 * time.Second,
		}

		session, err = mgo.DialWithInfo(dialInfo)
		if err != nil {
			return
		}
	}

	v, err := session.BuildInfo()
	if err != nil {
		return
	}

	fmt.Println(fmt.Sprintf("Mongo Version [%s]", v.Version))
	this.session = session
	return
}

func (this *BWMongo) Check() (err error) {

	if os.Getenv(BW_MONGO_ENDPOINT) == "" {
		// 环境变量不存在, 检查bwidow_mongo.toml文件
		if _, err := os.Stat("bwidow_mongo.toml"); os.IsNotExist(err) {
			//	toml 配置文件不存在,检查bwidow_mongo.json文件
			if _, err := os.Stat("bwidow_mongo.json"); os.IsNotExist(err) {
				return errors.New(fmt.Sprintf("Can not find Mongo configure. Env[%s]/Toml/Json are all lost!", BW_MONGO_ENDPOINT))
			}
		}

		return nil
	}

	if os.Getenv(BW_MONGO_DB) == "" {
		return errors.New(fmt.Sprintf("[%s] Empty!", BW_MONGO_DB))
	}

	return nil
}

func (this *BWMongo) checkIndex(uPtr interface{}) (err error) {
	this.setDB()
	defer this.db.Session.Close()

	m := zReflect.ReflectStructInfoWithTag(uPtr, true, "bw")

	var key []string
	for k, _ := range m {
		key = append(key, k)
	}

	index, _ := this.db.C(this.tableMap[getTypeName(uPtr)]).Indexes()
	isExist := false
	for _, idx := range index {
		if len(m) == len(idx.Key) {
			i := 0
			for k, _ := range m {
				for _, n := range idx.Key {
					if k == n {
						i++
					}
				}
			}
			if len(m) == i {
				isExist = true
				break
			}
		}
	}

	if !isExist {
		index := mgo.Index{
			Key:        key,
			Unique:     true,
			DropDups:   true,
			Background: false,
			Sparse:     true,
		}
		err = this.db.C(this.tableMap[getTypeName(uPtr)]).EnsureIndex(index)
	}

	return
}
