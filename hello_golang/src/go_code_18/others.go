package main

import (
	"fmt"
	"strconv"
)

func main(){
	test1() //type
	test2() //error
}
//type
func test1() {
	test1()
	var i1 myint
	var i2 = 100
	i1 = 200
	//i1 = i2 此时便会报错,在语法上myint和int已经是两种数据类型了
	fmt.Println(i1,i2)
	var name mystr
	name = "小花子"
	fmt.Println(name)

	fmt.Println("-------------------------")
	res1 := fun1()
	fmt.Println(res1(10,20))

	fmt.Println("-----------------------")
	var i3 myint2
	i3 = 1000
	i3 = i2 //此时就不会报错了 因为myint2实际上也是int
	fmt.Println(i3)
}
//error
/**
1.定义错误的两种方式
return errors.New("错误信息")
2.err:=fmt.Errorf("错误信息") return err
*/
func test2(){

}

//1.定义一个新的类型
type myint int
type mystr string
//2.定义函数类型
type myfunc func(int,int)(string)
func fun1() myfunc{
	fun := func(a,b int)string{
		s := strconv.Itoa(a) + strconv.Itoa(b)
		return s
	}
	return fun
}
//3.类型别名
type myint2 = int  //不是重新定义新的数据类型而是起了一个别名