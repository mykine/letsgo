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
	fmt.Print(NAME+"好好学习 天天向上")
}

//入口函数
func main() {
	studyKnowladge()
}
