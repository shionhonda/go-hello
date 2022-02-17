package main

import (
	"fmt"
	"log"
	"reflect"
)

const billionth float64 = 1e-9

// nearlyEqual returns true iff the two values (of any numerical type)
// are within 1e-9 difference of each other.
func NearlyEqual(x, y interface{}) bool {
	return nearlyEqual(reflect.ValueOf(x), reflect.ValueOf(y))
}

func nearlyEqual(x, y reflect.Value) bool {
	xf, err := convertFloat64(x)
	if err != nil {
		log.Fatal(err)
	}
	yf, err := convertFloat64(y)
	if err != nil {
		log.Fatal(err)
	}
	return xf-yf < billionth && yf-xf < billionth
}

func convertFloat64(x reflect.Value) (float64, error) {
	switch x.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(x.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return float64(x.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return x.Float(), nil
	}
	return 0, fmt.Errorf("unsupported type %s", x.Type())
}
