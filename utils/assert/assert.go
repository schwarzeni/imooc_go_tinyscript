package assert

import (
	"fmt"
	"reflect"
)

// Eq 工具函数，断言
func Eq(v1 interface{}, v2 interface{}) {
	if !reflect.DeepEqual(v1, v2) {
		panic(fmt.Sprintf("%v != %v", v1, v2))
	}
}

// Nil 判断是否为 nil
func Nil(v interface{}) {
	if v != nil {
		panic("v is not nil")
	}
}

// NotNil 判断是不为 nil
func NotNil(v interface{}) {
	if v == nil {
		panic("v is nil")
	}
}
