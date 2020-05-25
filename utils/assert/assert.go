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
