package assert

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// True panics if v is not true
func True(v bool) {
	if !v {
		panic(getCaller() + "Should be true")
	}
}

// False panics if v is not false
func False(v bool) {
	if v {
		panic(getCaller() + "Should be false")
	}
}

// Nil panics if v is not nil
func Nil(v interface{}) {
	switch {
	case reflect.ValueOf(v).Kind() == reflect.Ptr:
		if reflect.ValueOf(v).IsNil() {
			return
		}
	case v == nil:
		return
	}
	panic(getCaller() + fmt.Sprintf("Should be nil: recived: %#v", v))
}

// NotNil panics if v is nil
func NotNil(v interface{}) {
	if v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
		panic(getCaller() + "Should be not nil")
	}
}

func getCaller() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("Assert on %s:%d: ", file[strings.LastIndex(file, "/")+1:], line)
}
