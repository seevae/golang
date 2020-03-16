package main

import "fmt"

func main() {
	/**
		go语言中只有for循环  标准写法如下
	**/
	for i := 0; i < 10; i++ {
		fmt.Print("i ")
	}

	//for循环代替while循环的写法--》省略1,3表达式，只写判断条件
	num := 1
	for num >= 1 { //如果fo后面什么条件也不加那么默认为true 死循环
		println(num)
		num++
		if num == 5 {
			break
		}
	}

	//for练习题
	//1.打印58-23数字
	for j := 58; j >= 23; j-- {
		fmt.Print(j, " ")
	}
	//2.求1-100的和
	var sum int = 0
	for k := 1; k <= 100; k++ {
		sum = sum + k
	}
	fmt.Printf("1-100的和是: %d\n", sum)
	//3.打印1-100内,能够被3整除,但是不能被5整除的数字,统计被打印数字的个数,每行打印五个
	var count int = 0
	for q := 1; q <= 100; q++ {
		if q%3 == 0 && q%5 != 0 {
			fmt.Print(q, " ")
			count++
			if count%5 == 0 {
				fmt.Print("\n")
			}
		}
	}
	fmt.Print("\n")
	fmt.Printf("总共有%d个符合要求的数字\n", count)

	printPattern()
}

//打印菱形
func printPattern() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5-i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	//反过来打印另一边
	for i := 1; i <= 4; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 9-2*i; k++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
