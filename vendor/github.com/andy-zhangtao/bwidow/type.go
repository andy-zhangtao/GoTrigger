/*
 * Copyright (c) 2018.
 * andy-zhangtao <ztao8607@gmail.com>
 */

package bwidow

import "reflect"

// getTypeName获取给定的指针真正的数据类型
// u 必须是指针
func getTypeName(u interface{}) (name string) {
	rv := reflect.ValueOf(&u)
	for rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}

	return rv.Type().Name()
}
