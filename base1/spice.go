package main

import (
	"fmt"
	"reflect"
)
type stack struct {
	top int
	bottom int
	dataArr []int64
}
func main() {

	var s = []int64{}
	s = append(s, 1,2,3)
	fmt.Print(s)

	var i *int
	i = new(int)
	*i=1
	fmt.Print(*i)
	fmt.Print("\n")

	sta := new(stack)
	sta.top=0
	sta.bottom=0

	fmt.Print(reflect.TypeOf(sta.dataArr))
	sta.dataArr = append(sta.dataArr,1,2,3)
	fmt.Print("\n")
	fmt.Print(sta)
	fmt.Print("\n")
	sta.dataArr = append(sta.dataArr[:0],sta.dataArr[0+1:]...)
	fmt.Print("\n")
	fmt.Print(sta)
	fmt.Print("\n")
	data := sta.dataArr[1]
	fmt.Print(data)
}
