package sugar_plus

import (
	"fmt"
	"strconv"
)

func ToString(value interface{}) string {
	if value == nil {
		return ""
	}
	var result string
	switch value.(type) {
	case int:
		result = strconv.Itoa(value.(int))
	case int8:
		val := value.(int8)
		result = strconv.Itoa(int(val))
	case int32:
		val := value.(int32)
		result = strconv.FormatInt(int64(val), 10)
	case int64:
		result = strconv.FormatInt(value.(int64), 10)
	case float32:
		val := value.(float32)
		//result = strconv.FormatFloat(float64(val), 'f', 5, 32)
		result = fmt.Sprintf("%.2f", val)
	case float64:
		val := value.(float64)
		result = fmt.Sprintf("%.2f", val)
		//result = strconv.FormatFloat(value.(float64), 'f', 5, 64)
	case string:
		result = value.(string)
	}
	return result
}

type element interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~bool | ~string |
		float32 | float64
}

func StringTo[T element](strVal string) T {
	var result T
	var ti interface{} = &result
	switch v := ti.(type) {
	case *int:
		intV, _ := strconv.Atoi(strVal)
		*v = intV
	case *int32:
		val, _ := strconv.ParseFloat(strVal, 64)
		//int32V, _ := strconv.ParseInt(strVal, 0, 64)
		*v = int32(val)

	case *int64:
		val, _ := strconv.ParseFloat(strVal, 64)
		//intV, _ := strconv.ParseInt(strVal, 0, 64)
		*v = int64(val)

	case *float32:
		intV, _ := strconv.ParseFloat(strVal, 64)
		*v = float32(intV)

	case *float64:
		intV, _ := strconv.ParseFloat(strVal, 64)
		*v = intV
	}

	return result
}
