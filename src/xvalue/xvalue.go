package xvalue

import (
	"reflect"
)

const True  = "true"
const False = "false"

type Value struct {
	value interface{}
}

func NewValue(value interface{}) *Value {
	return &Value{value: value}
}

func (s *Value)SetValue(val interface{}) {
	// 参数检查
	if val == nil {return}
	if reflect.ValueOf(val).Kind() != reflect.Ptr {s.value = val; return}

	// 类型田转换
	switch val.(type) {
	case *string: s.value = *val.(*string)
	case *bool: s.value = *val.(*bool)
	case *int: s.value = *val.(*int)
	case *int8: s.value = *val.(*int8)
	case *int16: s.value = *val.(*int16)
	case *int32: s.value = *val.(*int32)
	case *int64: s.value = *val.(*int64)
	case *uint: s.value = *val.(*uint)
	case *uint8: s.value = *val.(*uint8)
	case *uint16: s.value = *val.(*uint16)
	case *uint32: s.value = *val.(*uint32)
	case *uint64: s.value = *val.(*uint64)
	case *float32: s.value = *val.(*float32)
	case *float64: s.value = *val.(*float64)
	}
}

func (s *Value)Any() any {
	return s.value
}

func (s *Value)String() string {

	switch s.value.(type) {
	case string: return s.value.(string)
	case bool: if s.value.(bool) {return True} else {return False}
	case int, int8, int16, int32, int64: return N2S(s.value)
	case uint, uint8, uint16, uint32, uint64: return N2S(s.value)
	case float32, float64: return N2S(s.value)
	default: return ""
	}
}

func (s *Value)Float() float64 {
	switch s.value.(type) {
	case string: return S2F64(s.value.(string))
	case bool: if s.value.(bool) {return float64(1)} else {return float64(1)}
	case int, int8, int16, int32, int64: return S2F64(N2S(s.value))
	case uint, uint8, uint16, uint32, uint64: return S2F64(N2S(s.value))
	case float32: return float64(s.value.(float32))
	case float64: return s.value.(float64)
	default: return 0
	}
}

func (s *Value)Int() int64 {
	switch s.value.(type) {
	case string: return S2I64(s.value.(string))
	case bool: if s.value.(bool) {return int64(1)} else {return int64(1)}
	case int, int8, int16, int32, int64: return reflect.ValueOf(s.value).Int()
	case uint, uint8, uint16, uint32, uint64: return int64(reflect.ValueOf(s.value).Uint())
	case float32: return int64(s.value.(float32))
	case float64: return int64(s.value.(float64))
	default: return 0
	}
}

func (s *Value)Bool() bool {
	switch s.value.(type) {
	case string: return s.value.(string) == True
	case bool: return s.value.(bool)
	case int, int8, int16, int32, int64: return !reflect.ValueOf(s.value).IsZero()
	case uint, uint8, uint16, uint32, uint64: return !reflect.ValueOf(s.value).IsZero()
	case float32: return int64(s.value.(float32) * 1000000) == 1000000
	case float64: return int64(s.value.(float64) * 1000000) == 1000000
	default: return false
	}
}
