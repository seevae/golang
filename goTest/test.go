package main

import "fmt"

func main() {
	// var a string = "good"
	// var w = "hhe"
	// fmt.Println(a + " " + w)

	// var c int = 1
	// var b int
	// fmt.Println(b, c)
	// //还可以按照c的形式打印

	// e := 22
	// var (
	// 	f int
	// 	g string = "peace"
	// )
	// fmt.Println(e, f, g)

	// //定义一个常量是用
	// const r string = "nihaozhongguo"
	// g = "haha"
	// fmt.Println(r, g)

	// //有点问题
	// var number int = 101
	// fmt.Printf("%d\n", number)

	// //循环
	// for i := 0; i < 10; i++ {
	// 	fmt.Print(i, " ")
	// }

	// //相当于while循环
	// for e <= 25 {
	// 	fmt.Println("呵 ")
	// 	e++
	// }

	// //定义数组
	// strings := []string{"aa", "bb"}
	// length := len(strings)
	// fmt.Println("------", length)

	// for i := 0; i < 2; i++ {
	// 	fmt.Println(strings[i])
	// }

	// for i := range strings {
	// 	fmt.Println(strings[i])
	// }

	// //返回一个值的函数
	// var max int = maxNum(1, 2)
	// fmt.Println(max)

	// //返回两个值的函数
	// a, b := swap(1, "hao")
	// fmt.Print(a, b)

	// aa := 7
	// bb := 8
	// //值传递
	// numswap(aa, bb)
	// fmt.Print("值传递：", aa, bb)
	// //址传递
	// numswap2(&aa, &bb)
	// fmt.Println("址传递：", aa, bb)

	//指针
	var number int = 101
	var p *int
	p = &number
	fmt.Printf("number的地址是 %x\n", &number)
	fmt.Printf("number的地址是 %x\n", p)
	fmt.Printf("number地址中的值是 %d\n", *p)
	fmt.Println("=========")

	//二级指针
	var k = 300
	var p1 *int
	var p2 **int
	p1 = &k
	p2 = &p1
	fmt.Printf("k的地址是 %x", &k)
	fmt.Printf("k的地址是 %x, 即p1", p1)
	fmt.Printf("p1的地址是 %x", p2)
	fmt.Print("---------")
	fmt.Printf("k的值为%d,%d", *p1, *(*p2))
	fmt.Print("---------\n")

	//Go语言指针作为函数参数
	var a1 int = 100
	var b1 int = 200
	swapvalue(&a1, &b1) //相当于： a1,b1 = b1,a1
	fmt.Printf("a1=%d,b1=%d", a1, b1)
	/*
	   补充 a1,b1 = b1,a1
	   实际上是编译器为我们创建了一个临时变量，在编译的时候执行的还是正常步骤的交换
	   temp = a1
	   a1 = b1
	   b1 = temp
	*/

}

func swapvalue(a *int, b *int) {
	var temp = *a
	*a = *b
	*b = temp
}

func numswap(a, b int) {
	var temp = a
	a = b
	b = temp
}

func numswap2(a *int, b *int) {
	var temp = *a
	*a = *b
	*b = temp
}

func maxNum(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

//go中函数定义--》 函数名形参的定义与Java相同，不同的是返回值
//返回值可以返回多个，返回时应该对应形参的返回顺序，且写在形参定义之后
func swap(str1 int, str2 string) (string, int) {
	return str2, str1
}
