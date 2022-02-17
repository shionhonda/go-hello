package main

import (
	"reflect"
	"unsafe"
)

type ptr struct {
	x unsafe.Pointer
	t reflect.Type
}

// IsCycle determines if the argument has a cycle data structure or not
func IsCycle(x interface{}) bool {
	seen := make(map[ptr]bool)
	return isCycle(reflect.ValueOf(x), seen)
}

func isCycle(x reflect.Value, seen map[ptr]bool) bool {
	if x.CanAddr() {
		xptr := ptr{unsafe.Pointer(x.UnsafeAddr()), x.Type()}
		if seen[xptr] {
			return true
		}
		seen[xptr] = true
	}
	switch x.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if isCycle(x.Index(i), seen) {
				return true
			}
		}
	case reflect.Struct:
		for i := 0; i < x.NumField(); i++ {
			if isCycle(x.Field(i), seen) {
				return true
			}
		}
	case reflect.Map:
		for _, key := range x.MapKeys() {
			if isCycle(x.MapIndex(key), seen) {
				return true
			}
		}
	case reflect.Ptr, reflect.Interface:
		return isCycle(x.Elem(), seen)
	}
	return false
}
