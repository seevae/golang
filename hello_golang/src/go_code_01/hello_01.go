package main

import "fmt"

func main() {
	//if分支语句
	a := 12
	if a > 12 {
		fmt.Printf("%d", a)
	} else {
		fmt.Printf("%d\n", a-2)
	}
	//if..else if
	b := 22
	if b < 10 {
		fmt.Println(b + 1)
	} else if b == 22 {
		fmt.Println(b)
	} else {
		fmt.Println(b - 1)
	}
	//go中if语句的一种特殊用法
	/**
		if statement;condition{

		}
		解释： 可以在if后先定义变量，然后在写判断语句。该变量的作用域仅为该if语句块
	**/
	if num := 55; num > 0 {
		fmt.Printf("num 大于0 他的值=%d", num)
	}

	//go中的switch语句
	/**
		Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case
		而是跳出整个switch，但是可以使用fallthrought强制执行后面的case代码
	注意： switch可以作用在其他类型上，case后的数值必须和switch作用的变量类型一致，例如下面例子，
				switch作用的变量为int类型，所以case后也要跟int类型
	**/
	grade := "X"
	marks := 90
	switch marks {
	case 90, 85, 99:
		grade = "A"
		fmt.Println("1111111")
		//break //switch中的break用于强制中断case语句，下面的222不再被打印
		fmt.Println("222222")
		fallthrough //fallthrough一般写在case的最后一行，这个case后的下一个case直接执行，也叫做case穿透
	case 80, 75, 70:
		grade = "B"
		fmt.Println("fallthrough导致的case穿透")
	default:
		println(grade)
	}

	//go语言中的循环 只有for循环

}
