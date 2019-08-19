//本文件程序所属包
package main

//导入依赖包
import "fmt"

//常量定义
const NAME string = "jo"

//全局变量声明和复制
var a string = "哈哈哈"

//声明一般类型的别名
type ageInt int

//结构体声明
type user struct{

}

//声明接口
type Iuser interface {
	
}

//声明函数
func studyKnowladge(){
	fmt.Print(NAME+"好好学习 天天向上\n")
}

const A1  = iota
const A2  = iota
const (
	A3=iota
	A4=iota
	A5=iota
	_
	_
	A6=iota
)

//入口函数
func main() {
	fmt.Print("\nA1=")
	fmt.Print(A1)
	fmt.Print("\nA2=")
	fmt.Print(A2)
	fmt.Print("\nA3=")
	fmt.Print(A3)
	fmt.Print("\nA4=")
	fmt.Print(A4)
	fmt.Print("\nA5=")
	fmt.Print(A5)
	fmt.Print("\nA6=")
	fmt.Print(A6)
	fmt.Print("\n")
	studyKnowladge()
}
