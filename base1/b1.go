package main

import (
	"fmt"
	"letsgo/base1/s1"
	"letsgo/base1/s2"
	"reflect"
)

const NMAE string = "JO"

func main() {
	var age float32 = 3.12
	var a,b,c int = 1,2,3
	var d = int(age)
	fmt.Print(a)
	fmt.Print("\n")
	fmt.Print(b)
	fmt.Print("\n")
	fmt.Print(c)
	fmt.Print("\n")
	fmt.Print(age)
	fmt.Print("\n")
	fmt.Print(d)
	fmt.Print("\n")

	var str1 = string(a)
	var str2 = string("2")
	var str3 = string(a)
	fmt.Print(reflect.TypeOf(str1))
	fmt.Print("\n")
	fmt.Print(str2)
	fmt.Print("\n")
	fmt.Print(str3)
	fmt.Print("\n")

	s1.Showcode()
	fmt.Print("\n")
	s2.Show2()
	fmt.Print("\n")
}
