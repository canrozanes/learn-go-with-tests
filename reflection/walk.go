package main

import (
	"fmt"
	"reflect"
	"strings"
)

// walk takes a struct x and calls fn for all strings fields found inside. difficulty level: recursively.
func walk(x interface{}, fn func(input string)) {
	fmt.Println("x is: ", x)
	val := getValue(x)
	fmt.Println("Val of x is: ", val)
	fmt.Println("Val.Kind() of x is: ", val.Kind())

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

func main() {
	x := struct {
		Name string
	}{"Can"}
	walk(x, func(input string) {
		strings.ToTitle(input)
	})
}
