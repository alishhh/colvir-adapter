package mappers

import (
	"math"
	"strconv"
)

func makeStringValue(obj interface{}) string {
	switch t := obj.(type) {
	case int:
		return strconv.Itoa(t)
	case int64:
		return strconv.FormatInt(t, 10)
	case float64:
		f := strconv.FormatFloat(t, 'f', 2, 64)
		return f
	case string:
		return t
	}
	return ""
}

func makeIntValue(obj interface{}) int {
	switch t := obj.(type) {
	case int:
		return t
	case int64:
		return int(t)
	case float64:
		return int(math.Round(t))
	case string:
		r, err := strconv.Atoi(t)
		if err != nil {
			return -1
		}
		return r
	}
	return -1
}

func makeStringValuePtr(obj interface{}) *string {
	res := makeStringValue(obj)
	if res != "" {
		return &res
	}
	return nil
}

func makeString(input interface{}) string {
	ret, _ := input.(string)
	return ret
}

func makeStringPtr(input interface{}) *string {
	ret, _ := input.(string)
	if ret == "" {
		return nil
	}
	return &ret
}

func makeInt(input interface{}) int64 {
	ret, _ := input.(int64)
	return ret
}

func makeFloat(input interface{}) float64 {
	ret, _ := input.(float64)
	return ret
}
