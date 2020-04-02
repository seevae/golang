package main

import "fmt"

func main() {
	/**匿名函数
	作用: 1.将匿名函数作为另一个函数的参数 --->回调函数
	      2.将匿名函数作为另一个函数的返回值 --->可以形成闭包结构
	 */
	//函数赋值并调用
	fun2 := fun1
	fun2()

	//匿名函数:定义时不加函数名,在函数体后直接跟括号进行调用,通常函数只能使用一次
	func() {
		fmt.Println("我是一个匿名函数")
	}()
	//匿名函数可以当作变量来给其他函数赋值
	fun3 := func() {
		fmt.Println("我也是一个匿名函数")
	}
	fun3()

	//定义带返回值的匿名函数
	res1:=func(a int, b int) int {
		return a+b
	}(10,20)   //如果进行了函数调用则返回的值是函数的执行结果
	fmt.Println(res1)

	res2:=func(a int,b int)int{
		return a*b
	}               //如果没有进行函数调用返回的值是函数引用的地址值,可以通过这个引用来调用函数
	fmt.Println(res2)
	fmt.Println(res2(10,20))  //200
}

func fun1() {
	fmt.Println("我是一个正常函数")
}
