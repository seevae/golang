package main

import "fmt"

func main() {
	//go语言支持函数式编程的最大特点就是函数可以当作参数

	//调用increment函数
	res:= increment() //函数返回值是一个函数 即 res=fun
	resNum1:= res();
	fmt.Println(resNum1) //1
	resNum2:= res();
	fmt.Println(resNum2) //2

	res2:=increment()
	resNum3:=res2()
	fmt.Println(resNum3)
	/**
		闭包(closure)
			一个外层函数中(increment),有内层函数(fun),
			该内层函数中,会操作外层函数的局部变量(外层函数中的参数,或者外层函数中直接定义的变量-->i),
			并且该外层函数的返回值就是这个内层函数
		这个内层函数和外层函数的局部变量,统称为闭包结构
	*/
	/*
		注意:
			一般的局部变量会随着函数的调用而产生,函数的结束而销毁
			但是闭包结构中的外层函数的局部变量不会随着外层函数的结束而销毁,因为内层函数还要继续使用
	*/

}
func increment()(func()int){
	i := 0
	fun := func() int {
		i++
		return i
	}
	/**3.返回该匿名函数
	3.1 直接return该函数
	3.2 将函数赋值给一个变量将变量返回
	*/
	return fun
}
