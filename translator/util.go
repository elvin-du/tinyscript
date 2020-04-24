package translator

import (
	"reflect"
	"strings"
)

func IsInstanceOfExpr(instance interface{}) bool {
	ival := reflect.ValueOf(instance)
	return strings.LastIndex(ival.Type().String(), "ast.Expr") != -1
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	//if vi.Kind() == reflect.Ptr {
	return vi.IsNil()
	//}
	//return false
}
