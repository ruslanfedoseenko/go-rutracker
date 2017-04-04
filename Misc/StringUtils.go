package Misc

import (
	"fmt"
	"strings"
)

func ArrayToString(source []interface{}) (output string){
	sourceLen := len(source)
	for i := 0; i < sourceLen; i++ {
		output = fmt.Sprintf("%s, %v", output, source[i])
	}
	output = strings.TrimLeft(output, ", ")
	return
}

func ToInterfaceSlice(slice []int) (res []interface{}){
	sliceLen := len(slice)
	res = make([]interface{}, sliceLen)
	for index, item:= range slice {
		res[index] = item
	}
	return;
}