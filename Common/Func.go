package Common

import (
	"fmt"
	"reflect"
)

func Val2Float(val interface{}) float64 {
	var res float64
	switch val.(type) {
	case int64:
		res = float64(val.(int64))
	case float64:
		res = val.(float64)
	case string:
		fmt.Sscanf(val.(string), "%f", &res)
	}
	return res
}

func Val2Int(val interface{}) int64 {
	var res int64
	switch val.(type) {
	case int64:
		res = val.(int64)
	case float64:
		res = int64(val.(float64))
	case string:
		fmt.Sscanf(val.(string), "%d", &res)
	}
	return res
}

func Arithmetic(leftVal interface{}, rightVal interface{}, op Operator) interface{} {
	var res interface{}
	lT, rT := reflect.TypeOf(leftVal).Kind(), reflect.TypeOf(rightVal).Kind()
	if lT == reflect.Float64 || rT == reflect.Float64 {
		lv, rv := Val2Float(leftVal), Val2Float(rightVal)
		switch op {
		case ASTERISK:
			res = lv + rv
		case SLASH:
			res = lv / rv
		case PERCENT:
			res = nil
		case PLUS:
			res = lv + rv
		case MINUS:
			res = lv - rv
		}

	} else {
		lv, rv := Val2Int(leftVal), Val2Int(rightVal)
		switch op {
		case ASTERISK:
			res = lv * rv
		case SLASH:
			res = lv / rv
		case PERCENT:
			res = lv % rv
		case PLUS:
			res = lv + rv
		case MINUS:
			res = lv - rv
		}
	}

	return res
}
