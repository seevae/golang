package main

import (
	"fmt"
	"strconv"
	"time"
)

func main(){

}

//如何创建一个channel
func test1(){
	var ch chan int
	if ch == nil{
		fmt.Println("通道为空,需要先创建通道...")
		ch = make(chan int)
		fmt.Printf("%T,%v",ch,ch)
	}
	fmt.Println("通道已经存在..")
	checkChannel(ch)
}
//channel的调用
func checkChannel(ch chan int){
	fmt.Printf("%T,%v",ch,ch)
}

//通道间数据传递
func test2(){
	var ch chan bool
	ch = make(chan bool)
	go func() {
		for i:=0;i<11;i++{
			fmt.Println("子goroutine...",i)
		}
		ch <- true
	}()

	data := <-ch
	fmt.Println("main函数读到通道内数据...",data)
	fmt.Println("main函数执行完毕...")
}

//关闭通道和通道上范围循环
func test3(){
	ch := make(chan int)
	go sendData(ch)
	for {
		data,ok := <-ch //如果从通道中读到数据则返回true
		time.Sleep(1*time.Second)
		fmt.Println("数据已读..")
		if !ok{
			fmt.Println("数据已经读完",ok)
			break
		}
		fmt.Println("读到了数据",data,ok)
	}
	fmt.Println("main函数执行完毕...")
} //普通循环的方式
func sendData(ch chan int){
	for i:=0;i<10;i++{
		fmt.Println("数据写入通道..")
		ch <- i
	}
	close(ch)  //将ch通道关闭 , 否则会发生死锁
}
//for...range的方式使用通道
func test4(){
	ch := make(chan int)
	go sendData2(ch)
	for v := range ch{  //就是从通道中读取数据  相当于 v <- ch
		fmt.Println("通道中读取内容:",v)
	}
	fmt.Println("main...over...")
}
func sendData2(ch chan int){
	for i:=0;i<10;i++{
		ch<-i
	}
	close(ch)
}

//带缓冲区的通道
func sendData3(ch chan string){
	for i:=0;i<10;i++{
		ch <- "数据"+strconv.Itoa(i)
		fmt.Printf("子goroutine写入第%d个数据\n",i)
	}
	close(ch)
}
func test5(){
	ch := make(chan string,5)
	go sendData3(ch)
	for{
		data,ok := <-ch
		if !ok{
			fmt.Println("通道内已读完")
			break
		}
		fmt.Println("\t主goroutine从通道中读到的内容:",data)
	}
	fmt.Println("main...over.")
}