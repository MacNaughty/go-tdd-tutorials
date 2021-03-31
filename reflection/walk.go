package reflection

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}

	// switch val.Kind() {
	// case reflect.Struct:
	// 	for i := 0; i < val.NumField(); i++ {
	// 		walk(val.Field(i).Interface(), fn)
	// 	}
	// case reflect.Slice:
	// 	for i := 0; i < val.Len(); i++ {
	// 		walk(val.Index(i).Interface(), fn)
	// 	}
	// case reflect.String:
	// 	fn(val.String())
	// }
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
