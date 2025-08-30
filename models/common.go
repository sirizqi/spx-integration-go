package models

import (
	"encoding/json"
	"strconv"
)

type BaseResp[T any] struct {
	RetCode int    `json:"ret_code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type Header struct {
	AppID     string
	CheckSign string
	Timestamp int64
	RandomNum int64
}

type FloatString float64

func (f *FloatString) UnmarshalJSON(b []byte) error {
	var num float64
	if err := json.Unmarshal(b, &num); err == nil {
		*f = FloatString(num)
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if s == "" {
		*f = 0
		return nil
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*f = FloatString(v)
	return nil
}
