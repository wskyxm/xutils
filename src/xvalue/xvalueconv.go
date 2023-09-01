package xvalue

import (
	"fmt"
	"strconv"
)

func S2UI(val string) uint {
	return uint(S2UI64(val))
}

func S2UI32(val string) uint32 {
	return uint32(S2UI64(val))
}

func S2UI64(val string) uint64 {
	res, _ := strconv.ParseUint(val, 10, 64)
	return res
}

func S2I(val string) int {
	return int(S2I64(val))
}

func S2I32(val string) int32 {
	return int32(S2I64(val))
}

func S2I64(val string) int64 {
	res, _ := strconv.ParseInt(val, 10, 64)
	return res
}

func S2F32(val string) float32 {
	res, _ := strconv.ParseFloat(val,64)
	return float32(res)
}

func S2F64(val string) float64 {
	res, _ := strconv.ParseFloat(val,64)
	return res
}

func S2Bool(val string) bool {
	if val == "true" {return true} else {return S2I64(val) != 0}
}

func N2S(val any) string {
	switch val.(type) {
	case uint, uint8, uint16, uint32, uint64: return fmt.Sprintf("%d", val)
	case int, int8, int16, int32, int64: return fmt.Sprintf("%d", val)
	case float32, float64: return fmt.Sprintf("%f", val)
	default: return ""
	}
}
