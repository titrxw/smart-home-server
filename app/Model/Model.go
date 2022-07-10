package model

import (
	"database/sql/driver"
	"fmt"
	database "github.com/titrxw/go-framework/src/Core/Database"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

type LocalTime time.Time

func (t LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(TimeFormat))), nil
}

func (t *LocalTime) UnmarshalJSON(formatTime []byte) error {
	if formatTime != nil {
		tmp, err := time.ParseInLocation(`"`+TimeFormat+`"`, string(formatTime), time.Local)
		if err != nil {
			return err
		}
		*t = LocalTime(tmp)
	}
	return nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}
func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type Model struct {
	database.ModelAbstract
}
