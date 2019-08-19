package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 3
	if a >1 {
		fmt.Print("a>1\n")
	}else if a<5 {
		fmt.Print("a<5\n")
	}else{
		fmt.Print("a不在1~5之间\n")
	}

	var param1 interface{}
	param1 = 123
	switch param1.(type) {
		case int:
			fmt.Print("int类型\n")
			break

		case string:
			fmt.Print("string类型\n")
			break
	}

	for i:=0;i<10;i++{
		fmt.Print(strconv.Itoa(i)+",")
	}
	var arr [10]int = [10]int {10,9,8,7,6,5,4,3,2,1}
	fmt.Print(arr)
	for key,value:=range arr{
		fmt.Print("\narr["+strconv.Itoa(key)+"]=")
		fmt.Print(value)
	}


}
