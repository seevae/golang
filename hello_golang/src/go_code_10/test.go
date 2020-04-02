package main

import "fmt"

func main(){
	/**
		高阶函数:
			根据go语言的数据类型特点,可以将一个函数作为另一个函数的参数

		fun1() fun2()
			将fun1()作为fun2()函数的参数
				fun2()函数: 叫做高阶函数,接收了一个函数作为参数的函数
				fun1()函数: 叫做回调函数,作为另一个函数的参数的函数
	*/
	//eg:设计一个函数,用于求两个整数的加减乘除运算
	fmt.Printf("%T\n",add)  //普通函数类型  func(int, int) int
	fmt.Printf("%T\n",oper)  //高阶函数类型  func(int, int, func(int, int) int) int
	//高阶函数的调用
	res1:=add(1,4)
	fmt.Println(res1)

	res2:=oper(10,20,add)
	fmt.Println(res2)
	fmt.Println(add)
	/*
		高阶函数的调用过程
			10,20和函数add分别作为参数传给了oper函数的参数 a,b,fun
			然后在oper函数内部调用了fun函数-->即add函数, 此时他俩引用的指向已经相同,见下方打印
			调用oper函数后计算出的返回值在作为oper的最终结果,便可以得到两数相加的结果
	*/

	//高阶函数调用匿名函数
	res3:= oper(20,50,func(a,b int)int{   //匿名函数的整体作为oper中fun参数,然后在oper中进行回调
		return a*b
	})
	fmt.Println(res3)

}
//加法运算
func add(a int,b int)int{
	return a+b
}
//真正实现加减乘除运算的函数
func oper(a,b int,fun func(int,int)int)int{
	fmt.Println(a,b,fun)  //先打印一下三个参数
	res:=fun(a,b)
	return res
}

