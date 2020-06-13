package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup  //sync包下提供的同步等待组

func main(){
	//直接使用用go关键使多协程一起工作往往会因为main函数提前结束而导致其他协程也结束
	//WaitGroup可以解决该问题
	wg.Add(2)
	go fun1()
	go fun2()
	fmt.Println("main进入阻塞状态,等待waitgroup中的子goroutine结束..")
	wg.Wait()
	fmt.Println("main函数执行完毕")
}

//WaitGroup解决协程执行不完的情况
func fun1(){
	for i:=1;i<10;i++{
		fmt.Println("fun1函数中的打印",i)
	}
	wg.Done()  //给wg等待组中的count计数器-1,实际就是 wg.add(-1)
}
func fun2(){
	for i:=1;i<10;i++{
		fmt.Println("fun2函数中的打印",i)
	}
	wg.Done()
}


//Goruutine初识
func test1()  {
	go printNum()
	for i:=0;i<100;i++{
		fmt.Printf("A %d\n",i)
	}
	//此处有一个问题: 如果不加下面的休眠,有可能在main函数执行完毕时,子协程还没执行完,此时主协程关闭,子协程也就跟着关闭了
	//该种情况的解决办法就是加入协程间通信 channel. 或者让主协程先睡眠一会
	//time.Sleep(10*time.Second)
	fmt.Println("main...over...")
}
func printNum(){
	for i:=0;i<100;i++{
		fmt.Printf("%d\n",i)
	}
}

