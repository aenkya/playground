package utils

import (
	"reflect"
	"runtime"
	"strings"
)

func FunctionName(f any) string {
	functionNameParts := strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")
	functionName := functionNameParts[len(functionNameParts)-1]

	return functionName
}
