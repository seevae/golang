package utils

import "fmt"

func SayHello(){
	fmt.Println("hello miss Wang")
	Count() //在同一个包下方法可以直接调用
}
