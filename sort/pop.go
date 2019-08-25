package main

import (
	"fmt"
)

func PopSort(arr *[]int){
	//fmt.Print(strconv.Itoa(len(arr)-1))
	for i:=0;i<len(*arr)-1;i++{
		for j:=0;j<len(*arr)-1-i;j++{
			if (*arr)[j]>(*arr)[j+1]{
				tmp:=(*arr)[j]
				(*arr)[j]=(*arr)[j+1]
				(*arr)[j+1]=tmp
			}
		}
	}
}


type Array []int
func (arr * Array) PopSort1(){
	for i:=0;i<len(*arr)-1;i++{
		isOK := true
		for j:=0;j<len(*arr)-i-1;j++{
			if (*arr)[j]>(*arr)[j+1]{
				isOK = false
				tmp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
		}
		if true == isOK{
			break
		}
	}
}

func (arr * Array) PopSort2()  {
	switchMaxIndex := len(*arr) -1
	lastSwitchIndex := 0
	for i:=0;i<len(*arr)-1;i++{
		isOK := true
		for j:=0;j<switchMaxIndex;j++{
			if (*arr)[j]>(*arr)[j+1]{
				isOK =false
				lastSwitchIndex = j
				tmp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
		}
		switchMaxIndex = lastSwitchIndex
		if true == isOK {
			break
		}
	}
}

func main() {
	//arr1 := []int{6,5,4,3,2,1}
	//fmt.Print("arr1排序前输出:")
	//fmt.Print(arr1)
	//fmt.Print("\narr1排序后输出:")
	//PopSort1(&arr1)
	//fmt.Print(arr1)

	arr11 := Array{6,5,4,3,2,1}
	fmt.Print(arr11)
	fmt.Print("\n")
	arr11.PopSort2()
	fmt.Print(arr11)

}
