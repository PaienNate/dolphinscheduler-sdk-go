package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"
const timezone = "Asia/Shanghai"

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	// TODO：找到更好的方案解决这个问题
	strData := string(data)
	if strData == "null" || strData == `""` || strData == "" {
		*t = Time(time.Time{}) // 设置为零值或你想要的默认值
		return nil
	}
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, strData, time.Local)
	*t = Time(now)
	return
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t Time) local() time.Time {
	loc, _ := time.LoadLocation(timezone)
	return time.Time(t).In(loc)
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
