package cycle

import (
	"reflect"
	"unsafe"
)

func IsCycle(i interface{}) bool {
	var seen []unsafe.Pointer
	return isCycle(reflect.ValueOf(i), seen)
}

func isCycle(v reflect.Value, seen []unsafe.Pointer) bool {
	if v.CanAddr() {
		p := unsafe.Pointer(v.UnsafeAddr())

		for _, ptr := range seen {
			if ptr == p {
				return true
			}
		}
		seen = append(seen, p)
	}

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		if isCycle(v.Elem(), seen) {
			return true
		}

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if isCycle(v.Index(i), seen) {
				return true
			}
		}

	case reflect.Map:
		for _, k := range v.MapKeys() {
			if isCycle(v.MapIndex(k), seen) {
				return true
			}
		}

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if isCycle(v.Field(i), seen) {
				return true
			}
		}

	}

	return false
}
