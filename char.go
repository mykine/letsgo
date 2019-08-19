package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var var1 int = 100
	var var2 = strconv.Itoa(var1)
	fmt.Print(var1)
	fmt.Print("\n")
	fmt.Print("字符串VAR2="+var2)
	fmt.Print("\n")
	var a byte = 0
	var b byte = 1
	fmt.Print(reflect.TypeOf(a))
	fmt.Print("\na=")
	fmt.Print(a)
	fmt.Print("\nb=")
	fmt.Print(b)
	fmt.Print("\na&b=")
	fmt.Print(a&b)
	fmt.Print("\na|b=")
	fmt.Print(a|b)
	fmt.Print("\na^b=")
	fmt.Print(a^b)
	fmt.Print("\nb^b=")
	fmt.Print(b^b)
	fmt.Print("\nb<<1=")
	fmt.Print(b<<1)
	fmt.Print("\nb>>1=")
	fmt.Print(b>>1)
	fmt.Print("\n")
}
