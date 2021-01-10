package reflection

import (
	"reflect"
)

//walk accepts an interface and a function
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.Chan:
		//iterate through all values sent through channel until it was closed with Recv()
		// see Recv func() (x Value, ok bool)
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		//Call calls the function v with the input arguments in.
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
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
	case reflect.String:
		fn(val.String())
	}

}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	//if x is a pointer to something, retrieve the value being pointed to
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
