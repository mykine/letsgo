package main

import (
	"fmt"
	"strconv"
)

//普通函数
func Add(a int,b int ) (int){
	return a+b
}

//多个返回值的hanshu
func Add_Sub(a,b int) (int,int){
	return a+b,a-b
}

/*
"类的"方法:带有接收者的函数，Go 语言不像其它面相对象语言一样可以写个类，然后在类里面写一堆方法，
但其实Go语言的方法很巧妙的实现了这种效果：
我们只需要在普通函数前面加个接受者（receiver，写在函数名前面的括号里面），
这样编译器就知道这个函数（方法）属于哪个struct了。
*/
type User struct {
	Name string
	Age int
}

func ( A User) showName(){
	fmt.Print("\n name="+A.Name)
}

func ( A User) showAge(){
	fmt.Print("\n age="+strconv.Itoa(A.Age))
}

func main() {
	var a,b int = 1,2
	var x1 = Add(a,b)
	fmt.Print("\nAdd函数加法结果="+strconv.Itoa(x1))
	var x2,x3 = Add_Sub(a,b)
	fmt.Print("\nAdd_Sub函数加减法结果X2="+strconv.Itoa(x2)+",x3="+strconv.Itoa(x3))

	var user = new(User)
	user.Name="Jo"
	user.showName()
	user.showAge()

}
