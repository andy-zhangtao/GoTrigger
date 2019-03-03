package util

import (
	"errors"
	zt "github.com/andy-zhangtao/gogather/time"
	"strconv"
	"strings"
)

func NextTime(interval int) (time int) {
	t := zt.Ztime{}
	_t, _ := strconv.Atoi(t.Now().UnixNano(10))

	return _t + interval
}

/**
	minutely → *-*-* *:*:00
      hourly → *-*-* *:00:00
       daily → *-*-* 00:00:00
     monthly → *-*-01 00:00:00
      yearly → *-01-01 00:00:00
 */
func ParseInterval(interval string) (intervalTime int, err error) {
	ints := strings.Split(interval, " ")
	if len(ints) != 2 {
		return intervalTime, errors.New("Invalid Interval Formate")
	}

	dayTime := strings.Split(ints[1], ":")
	if len(dayTime) != 3 {
		return intervalTime, errors.New("Invalid Daytime Formate")
	}

	dt := parseDayTime(dayTime)

	return dt, nil
}

// parseYearTime
// return the days. Mini 1 day
//func parseYearTime(yeartimes []string) (dt int) {
//	if yeartimes[2] == "*" {
//		return 1
//	}
//
//	t, err := strconv.Atoi(yeartimes[2])
//	if err != nil {
//		return 1
//	}
//
//	if t > 60 {
//		t = 60
//	}
//
//	dt = t
//
//	if daytimes[1] == "*" {
//		return dt
//	}
//
//	t, err = strconv.Atoi(daytimes[1])
//	if err != nil {
//		return dt
//	}
//
//	if t > 60 {
//		t = 60
//	}
//
//	dt += t * 60
//
//	if daytimes[0] == "*" {
//		return dt
//	}
//
//	t, err = strconv.Atoi(daytimes[0])
//	if err != nil {
//		return dt
//	}
//
//	if t > 24 {
//		t = 24
//	}
//
//	dt += t * 60 * 60
//
//	return dt
//}

// parseDayTime
// return the day interval time. Mini 1 second
func parseDayTime(daytimes []string) (dt int) {
	if daytimes[2] == "*" {
		return 1
	}

	t, err := strconv.Atoi(daytimes[2])
	if err != nil {
		return 1
	}

	if t > 60 {
		t = 60
	}

	dt = t

	if daytimes[1] == "*" {
		return dt
	}

	t, err = strconv.Atoi(daytimes[1])
	if err != nil {
		return dt
	}

	if t > 60 {
		t = 60
	}

	dt += t * 60

	if daytimes[0] == "*" {
		return dt
	}

	t, err = strconv.Atoi(daytimes[0])
	if err != nil {
		return dt
	}

	if t > 24 {
		t = 24
	}

	dt += t * 60 * 60

	return dt
}
