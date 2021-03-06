package main //mian函数是程序的入口 , 必须导入main包

import "fmt" // go中导包的方式

//全局变量,函数外定义,不能使用简短定义的方式
var a int = 11
var b = 22

//函数定义的方式
func main() {
	fmt.Println("Hello World") //go中一句代码完成后不需要分号
	age := 30
	var str string = "你已经成年了"
	if age >= 18 {
		fmt.Println(str)
	}

	//1. 变量的定义与取值赋值
	/**
	  go语言本质上是一种静态语言,即强类型语言..但是go语言中存在类型推断,所以在类型定义的时候可以使用var直接进行模糊定义
	  三种变量定义方式 : ->
	*/
	var num1 int = 30                                                          //直接定义好类型
	var num2 = 50                                                              //类型推断
	num3 := 60                                                                 //简短定义 , 这种简单定义的方式不能用来定义全局变量
	fmt.Printf("num1的类型是%T,num2=%d 他的类型是%T,num3=%d\n", num1, num2, num2, num3) //和c相似的打印方式

	//还可以多个变量同时定义
	var n1, f1, s1 = 100, 3.14, "大石"
	fmt.Println(n1, f1, s1)
	//还可以同时定义一组数据
	var (
		studentName = "李小花"
		high        = 165
		sex         = "女"
	)
	//fmt.Println(studentName,high,sex);
	fmt.Printf("姓名:%s,身高:%dcm,性别:%s\n", studentName, high, sex)

	//2. 常量的定义与使用の关键字 const
	const PATH string = "http:www.baidu.com"
	const PI = 3.14
	//常量在定义之后不使用是不会报错的, 常量一旦定义不能修改, 否则会报错
	//定义一组常量,一组常量中,如果某个常量没有初始值,默认和上一行一样
	const c1, c2, c3 = 100, 12, "hah"
	const (
		MALE int = 100
		UNKNOW
		FEMALE = 1
	)
	//使用常量组作为枚举类型
	const (
		SPRING  = 0
		SUNMMER = 1
		AUTUMN  = 2
		WINTER  = 3
	)

	//3. iota,特殊常量,可以认为是一个可以被编译器修改的常量
	//第一个iota等于0,后面的依次加1,例如如下 a=1,b=2,c=3
	const (
		a = iota
		b
		c
	)
	fmt.Println("---------")
	fmt.Println(a, b, c)

	//4.string类型中，''中的变量类型是int32的，""中的变量类型才是string
	fmt.Print("==============\n")
	str1 := 'A'
	str2 := "A"
	fmt.Printf("str1的类型是：%T 值是%d 内容是%c\nstr2的类型是：%T,", str1, str1, str1, str2)

	//5.go语言中的位运算符
	/**
	按位与&:对应位的值如果都为1才为1，有一个为0就为0
	按位或|：对应位的值如果都为0才为0，有一个为1就为1
	按位异或^：
		二元：a^b
			对应位的值不同为1，相同为0
		一元：^a
			按位取反
	位清空：&^
		对于a &^ b ,
			对于b上的每一个数值，
			如果为0，则取a对应位上的数值
			如果为1，则结果位就取0
	**/

	//go语言中的输入
	//使用fmt包下的scan系列，举两个例子，如下
	var x int
	var y float64
	fmt.Println("请输入一个整数，一个浮点数")
	fmt.Scanln(&x, &y)

	fmt.Scanf("%d,%f", &x, &y)

}
