package main

import "fmt"

func main() {
	/**
	1.defer语句用于延迟一个函数或者方法的执行,当同时有多个defer函数,则先defer的后执行
	应用场景:
		A.例如在读写文件(数据库操作中也常用)的时候我们常常会忘记关闭连接,
		那么使用defer修饰.close方法就可以在整个程序执行完的最后自动关闭连接
		B.go语言中关于异常的处理,使用panic()和recover()
			panic函数用于引发恐慌,导致程序中断执行
			recover函数用于恢复程序的执行,recover语法上要求必须在defer中执行
	2.defer函数传递参数时:defer函数调用时,就已经传递了参数数据,只是暂时不执行代码,即延迟的是代码的执行而不是函数的调用
	3.当一个函数调用另一个函数使用defer语句时,会在return前就执行defer语句
	*/
	defer fmt.Println("66666666") //defer是一种栈操作,如果有多个defer函数,则支持压栈出栈原则
	fun1("gege")
	fmt.Println("123456")
	defer fun1("hehehe") // 在主函数(外围函数)中的内容全部执行后在执行
	fmt.Println("打个")
	//没有defer则按顺序运行,如果有defer则 hehehe 最后打印

	//2.
	a:=2
	fmt.Println(a)
	defer fun2(2) //此时已经发生参数传递而不是等主函数执行完才调用函数传递参数
	a++
	fmt.Println(a)

	//3.
	res:= fun3()
	fmt.Println(res)
}

func fun3()int{
	fmt.Println("fun3开始执行...")
	//这里程序的执行被延迟,但往下走发现是return,所以执行完毕后在执行renturn操作
	//程序结果: 先嘿嘿嘿在打印0
	defer fun1("嘿嘿嘿")
	return 0;
}

func fun2(n int){
	fmt.Println(n)  //此处打印2 说明defer处已经发生参数传递,但函数执行被延迟
}

func fun1(s string) {
	fmt.Println("11111")
	defer fmt.Println(s) //外围函数为fun1中另外的内容,最后执行
	fmt.Println("22222")
}
