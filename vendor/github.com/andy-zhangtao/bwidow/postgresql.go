/*
 * Copyright (c) 2018.
 * andy-zhangtao <ztao8607@gmail.com>
 */

package bwidow

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/andy-zhangtao/gogather/zReflect"
	_ "github.com/lib/pq"
	"github.com/pelletier/go-toml"
)

type BWPostgresConf struct {
	Endpoint string
	DB       string
	User     string
	Password string
}

type BWPostgresql struct {
	db       *sql.DB
	tableMap map[string]string
}

var pqconf *BWPostgresConf

func (this *BWPostgresql) setDB() {

}
func (this *BWPostgresql) Check() error {
	if os.Getenv(BW_PQ_ENDPOINT) == "" {
		if _, err := os.Stat("bwidow_pq.toml"); os.IsNotExist(err) {
			//	toml 配置文件不存在,检查bwidow_mongo.json文件
			if _, err := os.Stat("bwidow_pq.json"); os.IsNotExist(err) {
				return errors.New(fmt.Sprintf("Can not find Mongo configure. Env[%s]/Toml/Json are all lost!", BW_PQ_ENDPOINT))
			}
		}

		return nil
	}

	if os.Getenv(BW_PQ_DB) == "" {
		return errors.New(fmt.Sprintf("[%s] Empty!", BW_PQ_DB))
	}

	return nil
}

func (this *BWPostgresql) DriverInit() error {
	if err := this.Check(); err != nil {
		return err
	}

	pqconf = new(BWPostgresConf)
	pqconf.User = os.Getenv(BW_PQ_USER)
	pqconf.Password = os.Getenv(BW_PQ_PASSWD)
	pqconf.Endpoint = os.Getenv(BW_PQ_ENDPOINT)
	pqconf.DB = os.Getenv(BW_PQ_DB)

	if pqconf.Endpoint == "" {
		useToml := true
		//	读取toml
		data, err := ioutil.ReadFile("bwidow_pq.toml")
		if err != nil || len(data) == 0 {
			useToml = false
			//尝试json
			data, err = ioutil.ReadFile("bwidow_pq.json")
			if err != nil || len(data) == 0 {
				return errors.New(fmt.Sprintf("Parse Configure Error %s", err.Error()))
			}
		}

		if useToml {
			err = toml.Unmarshal(data, pqconf)
			if err != nil {
				return err
			}
		} else {
			err = json.Unmarshal(data, pqconf)
			if err != nil {
				return err
			}
		}
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", pqconf.User, pqconf.Password, pqconf.Endpoint, pqconf.DB)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return errors.New(fmt.Sprintf("Ping [%s] Error [%s]", connStr, err))
	}

	this.db = db
	fmt.Printf("Postgresql [%s] connect success! \n", pqconf.Endpoint)

	return nil
}

func (this *BWPostgresql) Map(u interface{}, name string) {
	if this.tableMap == nil {
		this.tableMap = make(map[string]string)
	}

	this.tableMap[reflect.TypeOf(u).Name()] = name
	return
}

func (this *BWPostgresql) checkIndex(u interface{}) error {
	m := zReflect.ReflectStructInfoWithTag(u, true, "bw")

	var key []string
	for k, _ := range m {
		key = append(key, k)
	}

	table := this.tableMap[getTypeName(u)]

	sql := fmt.Sprintf("CREATE UNIQUE INDEX IF NOT EXISTS %s_unique on %s ( %s )", table, table, strings.Join(key, ","))

	_, err := this.db.Exec(sql)

	return err
}

func (this *BWPostgresql) first(uPtr interface{}) error {

	table := this.tableMap[getTypeName(uPtr)]

	uStruct := zReflect.ReflectStructInfoWithTag(uPtr, true, "pq")

	var columns []string

	for key, _ := range uStruct {
		columns = append(columns, key)
	}

	valPtr, err := zReflect.ExtractValuePtrFromStruct(uPtr, columns)
	if err != nil {
		return err
	}

	sql := fmt.Sprintf("SELECT %s FROM %s LIMIT 1", strings.Join(columns, ","), table)

	rows, err := this.db.Query(sql)
	if err != nil {
		return errors.New(fmt.Sprintf("Query First Rows Error [%s]", err.Error()))
	}

	for rows.Next() {
		err = rows.Scan(valPtr...)
		if err != nil {
			return err
		}
	}

	rows.Close()
	return nil
}

func (this *BWPostgresql) findOne(uPtr interface{}, fields ...string) error {
	table := this.tableMap[getTypeName(uPtr)]
	var columns []string
	var filter []string

	uValues := zReflect.ReflectStructInfoWithTag(uPtr, false, "pq")

	_uStruct := zReflect.ReflectStructInfoWithTag(uPtr, true, "pq")

	uStruct := make(map[string]interface{})

	if len(fields) == 0 {
		uStruct = _uStruct
	} else {
		for _, f := range fields {
			uStruct[f] = _uStruct[f]
		}
	}

	for key, value := range uValues {
		filter = append(filter, fmt.Sprintf(" %s='%v'", key, value))
	}

	for key, _ := range uStruct {
		columns = append(columns, key)
	}

	valPtr, err := zReflect.ExtractValuePtrFromStruct(uPtr, columns)
	if err != nil {
		return err
	}

	sql := fmt.Sprintf("SELECT %s FROM %s WHERE %s LIMIT 1", strings.Join(columns, ","), table, strings.Join(filter, " AND "))

	//fmt.Println(sql)
	rows, err := this.db.Query(sql)
	if err != nil {
		return errors.New(fmt.Sprintf("Query Rows Error [%s]", err.Error()))
	}

	for rows.Next() {
		err = rows.Scan(valPtr...)
		if err != nil {
			return err
		}
	}

	rows.Close()
	return nil
}
func (this *BWPostgresql) findAll(uPtr interface{}, aPtr interface{}, fields ...string) error {
	table := this.tableMap[getTypeName(uPtr)]

	var columns []string
	var filter []string

	uValues := zReflect.ReflectStructInfoWithTag(uPtr, false, "pq")

	uStruct := zReflect.ReflectStructInfoWithTag(uPtr, true, "pq")

	for key, value := range uValues {
		filter = append(filter, fmt.Sprintf(" %s='%v'", key, value))
	}

	if len(fields) > 0 {
		for _, f := range fields {
			delete(uStruct, f)
		}
	}

	for key, _ := range uStruct {
		columns = append(columns, key)
	}

	var sql string
	if len(filter) > 0 {
		sql = fmt.Sprintf("SELECT %s FROM %s WHERE %s ", strings.Join(columns, ","), table, strings.Join(filter, " AND "))
	} else {
		sql = fmt.Sprintf("SELECT %s FROM %s ", strings.Join(columns, ","), table)
	}

	rows, err := this.db.Query(sql)
	if err != nil {
		return errors.New(fmt.Sprintf("Query Rows Error [%s] SQL[%s]", err.Error(), sql))
	}

	valuePtr := reflect.ValueOf(aPtr)
	arrayValue := valuePtr.Elem()

	for rows.Next() {
		_valPtr := reflect.New(reflect.TypeOf(uPtr).Elem()).Interface()
		valPtr, err := zReflect.ExtractValuePtrFromStruct(_valPtr, columns)
		if err != nil {
			return err
		}
		err = rows.Scan(valPtr...)
		if err != nil {
			return err
		}

		arrayValue.Set(reflect.Append(arrayValue, reflect.ValueOf(_valPtr)))
	}

	rows.Close()
	return nil
}
func (this *BWPostgresql) findAllWithSort(uPtr interface{}, aPtr interface{}, sortField []string, fields ...string) error {

	table := this.tableMap[getTypeName(uPtr)]

	var columns []string
	var filter []string

	uValues := zReflect.ReflectStructInfoWithTag(uPtr, false, "pq")

	uStruct := zReflect.ReflectStructInfoWithTag(uPtr, true, "pq")

	if len(fields) > 0 {
		for _, f := range fields {
			delete(uStruct, f)
		}
	}

	for key, value := range uValues {
		filter = append(filter, fmt.Sprintf(" %s='%v'", key, value))
	}

	for key, _ := range uStruct {
		columns = append(columns, key)
	}

	var _sort []string
	for _, s := range sortField {
		if strings.HasPrefix(s, "-") {
			_sort = append(_sort, fmt.Sprintf(" %s DESC", s[1:len(s)]))
		} else if strings.HasPrefix(s, "+") {
			_sort = append(_sort, fmt.Sprintf(" %s ASC", s[1:len(s)]))
		} else {
			_sort = append(_sort, fmt.Sprintf(" %s ASC", s))
		}
	}

	var sql string
	if len(filter) > 0 {
		sql = fmt.Sprintf("SELECT %s FROM %s WHERE %s ORDER BY %s ", strings.Join(columns, ","), table, strings.Join(filter, " AND "), strings.Join(_sort, ","))
	} else {
		sql = fmt.Sprintf("SELECT %s FROM %s ORDER BY %s ", strings.Join(columns, ","), table, strings.Join(_sort, ","))
	}

	//fmt.Println(sql)
	rows, err := this.db.Query(sql)
	if err != nil {
		return errors.New(fmt.Sprintf("Query Rows Error [%s]", err.Error()))
	}

	valuePtr := reflect.ValueOf(aPtr)
	arrayValue := valuePtr.Elem()

	for rows.Next() {
		_valPtr := reflect.New(reflect.TypeOf(uPtr).Elem()).Interface()
		valPtr, err := zReflect.ExtractValuePtrFromStruct(_valPtr, columns)
		if err != nil {
			return err
		}
		err = rows.Scan(valPtr...)
		if err != nil {
			return err
		}

		arrayValue.Set(reflect.Append(arrayValue, reflect.ValueOf(_valPtr)))
	}

	rows.Close()
	return nil

}
func (this *BWPostgresql) save(uPtr interface{}) error {
	table := this.tableMap[getTypeName(uPtr)]
	var columns []string
	var value []interface{}

	uValues := zReflect.ReflectStructInfoWithTag(uPtr, false, "pq")

	for key, val := range uValues {
		columns = append(columns, key)
		value = append(value, val)
	}

	var _value string

	for _, v := range value {
		_value += fmt.Sprintf("'%v',", v)
	}

	_value = _value[:len(_value)-1]
	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", table, strings.Join(columns, ","), _value)

	//fmt.Println(sql)
	_, err := this.db.Exec(sql)
	if err != nil {
		return errors.New(fmt.Sprintf("Insert Row Error [%s]", err.Error()))
	}

	return nil
}
func (this *BWPostgresql) saveAll(uArray []interface{}) error {
	for _, u := range uArray {
		if err := this.save(u); err != nil {
			return err
		}
	}
	return nil
}
func (this *BWPostgresql) update(uPtr interface{}, field []string) (int, error) {

	table := this.tableMap[getTypeName(uPtr)]
	var filters []string
	var updates []string
	uStruct := zReflect.ReflectStructInfoWithTag(uPtr, false, "pq")

	if len(field) > 0 {
		for _, f := range field {
			if v, ok := uStruct[f]; ok {
				filters = append(filters, fmt.Sprintf("%s='%v'", f, v))
			}
		}
	} else {
		for key, value := range uStruct {
			filters = append(filters, fmt.Sprintf("%s='%v'", key, value))
		}
	}

	for key, value := range uStruct {
		updates = append(updates, fmt.Sprintf("%s='%v'", key, value))
	}

	if len(filters) == 0 {
		return 0, errors.New("bw: invalid update filter params")
	}

	sql := fmt.Sprintf("UPDATE %s SET %s WHERE %s ", table, strings.Join(updates, ","), strings.Join(filters, " AND "))
	fmt.Println(sql)
	result, err := this.db.Exec(sql)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rows), nil
}

func (this *BWPostgresql) delete(uPtr interface{}, field []string) (int, error) {
	table := this.tableMap[getTypeName(uPtr)]
	var columns []string
	uStruct := zReflect.ReflectStructInfoWithTag(uPtr, false, "pq")

	if len(field) > 0 {
		for _, f := range field {
			if v, ok := uStruct[f]; ok {
				columns = append(columns, fmt.Sprintf("%s='%v'", f, v))
			}
		}
	} else {
		for key, value := range uStruct {
			columns = append(columns, fmt.Sprintf("%s='%v'", key, value))
		}
	}

	if len(columns) == 0 {
		return 0, errors.New("bw: invalid delete filter params")
	}

	sql := fmt.Sprintf("DELETE FROM %s WHERE %s ", table, strings.Join(columns, " AND "))
	fmt.Println(sql)
	result, err := this.db.Exec(sql)
	if err != nil {
		return 0, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(rows), nil
}

func (this *BWPostgresql) deleteAll(uPtr interface{}) (int, error) {
	table := this.tableMap[getTypeName(uPtr)]

	sql := fmt.Sprintf("DELETE FROM %s ", table)
	_, err := this.db.Exec(sql)

	return 0, err
}
func (this *BWPostgresql) count(uPtr interface{}) (int, error) {
	var count int
	table := this.tableMap[getTypeName(uPtr)]

	sql := fmt.Sprintf("SELECT COUNT(*) FROM %s ", table)
	rows, err := this.db.Query(sql)
	if err != nil {
		return count, err
	}

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return count, err
		}
	}

	return count, nil
}
