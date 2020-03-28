package main

import "fmt"

/*
	函数也是一种类型,相同类型函数可以直接赋值例如定义好了fun1 可以直接 fun2:=fun1,fun2就有了和fun1相同的功能
		函数名():将函数进行调用.函数中的代码全部执行,然后将return的结果返回给调用处
		函数名:指向函数体的内存地址
	fmt.Printf("%T",fun1) 是可以打印出来函数的类型的
	fun1实际上是一个引用,他指向了函数在内存中的地址 所以直接fmt.Println(fun1)可以打印出来函数内村地址
    调用一个函数就是找到他在内村中的位置,运行其中的代码 .
*/

//全局变量的定义不支持简短定义,即: a:=1
var a int = 100 //全局变量所有函数都可以使用

func main() {
	getSum(10) //函数的初步认识
	///求a,b之和
	getadd(11, 19)
	/**可变参数
	语法:
		参数名... 参数类型
	对于函数,可变参数就相当于一个切片,调用函数时可以传入多个参数
	printXx系列函数,append函数 实际都是可变参数函数
	注意:~~~
		1.如果一个函数的参数是可变参数,同时还有其他参数,那么可变参数要放在参数列表的最后
		2.一个函数的参数列表中最多只能有一个可变参数列表
	*/
	res := getAllSum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(res)
	//可变参数列表可以直接计算一个切片中的所有值的和
	arr := []int{1, 2, 3, 4, 5}
	getAllSum(arr...) //切片后跟 ... 便可以当作多参数传入可变参数列表
	fun1(7.22, 1.22, "hehe", "heihei")

	//参数传递: 值传递(基本类型的参数传递都是值传递除非传地址), 引用传递
	//数组的值传递
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("调用函数前数组中的内容:", arr1)
	arrTest(arr1)
	fmt.Println("调用函数后前数组中的内容:", arr1)
	fmt.Println("------切片-----------")
	s1 := make([]int, 0, 0)
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println("调用函数前切片中的内容:", s1)
	sliceTest(s1)
	fmt.Println("调用函数后切片中的内容:", s1)

	fmt.Println("--------函数的多返回值-------------")
	//求矩形面积和周长
	area, perimeter := rectangle(10, 52) //如果只想要返回值其中之一就使用 _ 代替不想要的那个
	fmt.Printf("矩形的周长为:%d\n矩形的面积为:%d\n", perimeter, area)

	fmt.Println("********** 递归函数 ***********")
	//斐波那契
	fmt.Println(fibonic(10))

}

func fibonic(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fibonic(n-1) + fibonic(n-2)
}

//递归求1-5的和
func getRes(num int) int {
	if num == 1 {
		return 1
	}
	return getRes(num-1) + num
}

func rectangle(length, width int) (int, int) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

func sliceTest(s2 []int) {
	s2[2] = 789
	fmt.Println("改变切片中的值后切片中的内容", s2)
}

//数组:值传递
func arrTest(arr2 [4]int) { //实际是一个新的数组赋值了arr1的内容,两个数组已经是不同的地址
	arr2[2] = 18
	fmt.Println("改变数组中的值后数组的内容:", arr2)
}

//可变参数要放在参数列表的最后
func fun1(a float64, b float64, s ...string) {
	fmt.Println(a, b, s)
}

//可变参数
func getAllSum(nums ...int) int {
	fmt.Printf("%T\n", nums) //打印结果可知 可变参数列表实质就是一个切片
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	fmt.Println(sum)
	return sum
}

//返回值还有一种写法
func getAllSum2(nums ...int) (num int) {
	fmt.Println(nums)
	return //此种方式不需要在return后加返回的对象,因为在函数返回值处已经定义,默认就返回num
}

//两数之和
func getadd(a int, b int) {
	sum := a + b
	fmt.Printf("%d+%d=%d\n", a, b, sum)
}

//求和函数
func getSum(count int) {
	var sum int = 0
	for i := count; i > 0; i-- {
		sum += i
	}
	fmt.Printf("从1-%d的和是%d", count, sum)
}
