package goconv

import (
	"strconv"
)

// String converts any value to string, returns empty string if failed
func String(value any) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', 6, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', 6, 64)
	case bool:
		return strconv.FormatBool(v)
	case []bool:
		if len(v) > 0 {
			return strconv.FormatBool(v[0])
		}
		return ""
	case string:
		return v
	case []byte:
		return string(v)
	default:
		return ""
	}
}

// Int converts value to int, returns 0 if failed
func Int(value any) int {
	if value == nil {
		return 0
	}
	if v, ok := value.(int); ok {
		return v
	}
	return int(Int64(value))
}

// Int32 converts value to int32, returns 0 if failed
func Int32(value any) int32 {
	if value == nil {
		return 0
	}
	if v, ok := value.(int32); ok {
		return v
	}
	return int32(Int64(value))
}

// Int64 converts value to int64, returns 0 if failed
func Int64(value any) int64 {
	switch v := value.(type) {
	case int:
		return int64(v)
	case int8:
		return int64(v)
	case int16:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case uint:
		return int64(v)
	case uint8:
		return int64(v)
	case uint16:
		return int64(v)
	case uint32:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case string:
		result, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0
		}
		return result
	default:
		return 0
	}
}

// Uint converts value to uint, returns 0 if failed
func Uint(value any) uint {
	if value == nil {
		return 0
	}
	if v, ok := value.(uint); ok {
		return v
	}
	return uint(Uint64(value))
}

// Uint32 converts value to uint32, returns 0 if failed
func Uint32(value any) uint32 {
	if value == nil {
		return 0
	}
	if v, ok := value.(uint32); ok {
		return v
	}
	return uint32(Uint64(value))
}

// Uint64 converts value to uint64, returns 0 if failed
func Uint64(value any) uint64 {
	if value == nil {
		return 0
	}
	switch v := value.(type) {
	case int:
		if v < 0 {
			return 0
		}
		return uint64(v)
	case int8:
		if v < 0 {
			return 0
		}
		return uint64(v)
	case int16:
		if v < 0 {
			return 0
		}
		return uint64(v)
	case int32:
		if v < 0 {
			return 0
		}
		return uint64(v)
	case int64:
		if v < 0 {
			return 0
		}
		return uint64(v)
	case uint:
		return uint64(v)
	case uint8:
		return uint64(v)
	case uint16:
		return uint64(v)
	case uint32:
		return uint64(v)
	case uint64:
		return v
	case float32:
		if v < 0 {
			return 0
		}
		return uint64(v)
	case float64:
		if v < 0 {
			return 0
		}
		return uint64(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		result, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0
		}
		return result
	default:
		return 0
	}
}

// Float32 converts value to float32, returns 0 if failed
func Float32(value any) float32 {
	if value == nil {
		return 0
	}
	switch v := value.(type) {
	case float32:
		return v
	case float64:
		return float32(v)
	default:
		vv, _ := strconv.ParseFloat(String(value), 64)
		return float32(vv)
	}
}

// Float64 converts value to float64, returns 0 if failed
func Float64(value any) float64 {
	if value == nil {
		return 0
	}
	switch v := value.(type) {
	case float32:
		return float64(v)
	case float64:
		return v
	default:
		vv, _ := strconv.ParseFloat(String(value), 64)
		return vv
	}
}

// Bool converts value to bool, returns false if failed
func Bool(value any) bool {
	if value == nil {
		return false
	}
	switch v := value.(type) {
	case bool:
		return v
	case int, int8, int16, int32, int64:
		return Int64(v) != 0
	case uint, uint8, uint16, uint32, uint64:
		return Uint64(v) != 0
	case float32, float64:
		return Float64(v) != 0
	case string:
		result, err := strconv.ParseBool(v)
		if err != nil {
			return false
		}
		return result
	default:
		return false
	}
}
