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

	//返回一个值的函数
	var max int = maxNum(1, 2)
	fmt.Println(max)

	//返回两个值的函数
	a, b := swap(1, "hao")
	fmt.Print(a, b)

	aa := 7
	bb := 8
	//值传递
	numswap(aa, bb)
	fmt.Print("值传递：", aa, bb)
	//址传递
	numswap2(&aa, &bb)
	fmt.Println("址传递：", aa, bb)

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

func swap(str1 int, str2 string) (string, int) {
	return str2, str1
}
