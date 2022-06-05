package helpers

import (
	"strconv"
)

//ToString 转换成字符串
func ToString(a interface{}) string {
	if value, ok := a.([]uint8); ok {
		return string(value)
	}

	if value, ok := a.(bool); ok {
		if value {
			return "1"
		} else {
			return "0"
		}
	}

	if value, ok := a.([]byte); ok {
		return string(value)
	}

	if value, ok := a.(string); ok {
		return value
	}

	if value, ok := a.(int); ok {
		return strconv.Itoa(value)
	}

	if value, ok := a.(int8); ok {
		return strconv.Itoa(int(value))
	}

	if value, ok := a.(int16); ok {
		return strconv.Itoa(int(value))
	}

	if value, ok := a.(int32); ok {
		return strconv.Itoa(int(value))
	}

	if value, ok := a.(int64); ok {
		return strconv.FormatInt(value, 10)
	}

	if value, ok := a.(uint); ok {
		return strconv.FormatUint(uint64(value), 10)
	}

	if value, ok := a.(uint8); ok {
		return strconv.FormatUint(uint64(value), 10)
	}

	if value, ok := a.(uint16); ok {
		return strconv.FormatUint(uint64(value), 10)
	}

	if value, ok := a.(uint32); ok {
		return strconv.FormatUint(uint64(value), 10)
	}

	if value, ok := a.(uint64); ok {
		return strconv.FormatUint(value, 10)
	}

	if value, ok := a.(float32); ok {
		return strconv.FormatFloat(float64(value), 'f', -1, 32)
	}

	if value, ok := a.(float64); ok {
		return strconv.FormatFloat(value, 'f', -1, 64)
	}

	return ""
}

//ToInt 转换成整数
func ToInt(a interface{}) int {
	if value, ok := a.(int); ok {
		return value
	}

	if value, ok := a.(float64); ok {
		return int(value)
	}

	if value, ok := a.(float32); ok {
		return int(value)
	}

	if value, ok := a.(int64); ok {
		return int(value)
	}

	if value, ok := a.(string); ok {
		val, err := strconv.Atoi(value)
		if err == nil {
			return val
		}
	}

	return 0
}

//GetSpace 获取空格
func GetSpace() string {
	return string(rune(32))
}
