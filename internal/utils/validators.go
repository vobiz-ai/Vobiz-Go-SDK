package utils

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

// MultipleValidIntegers validates that a parameter value is either an integer
// or a "<"-separated string of integers (used by multi-destination APIs).
func MultipleValidIntegers(paramName string, paramValue interface{}) {
	switch reflect.TypeOf(paramValue).Kind() {
	case reflect.String:
		for _, part := range strings.Split(paramValue.(string), "<") {
			if _, err := strconv.Atoi(part); err != nil {
				log.Fatalf("%s: all values in the '<'-separated string must be integers", paramName)
			}
		}
	case reflect.Int:
		// valid
	default:
		log.Fatalf("%s must be either string or integer", paramName)
	}
}
