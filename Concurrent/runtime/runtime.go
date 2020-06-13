package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	goroot := runtime.GOROOT()
	fmt.Println(goroot)

	go func (){
		for i:=0;i<100;i++{
			fmt.Println("子goroutine...",i)
		}
	}()

	for i:=0;i<100;i++{
		runtime.Gosched()  //当前goroutine让出时间片,让其他goroutine先执行
		fmt.Println("主goroutine...",i)
	}

	time.Sleep(10*time.Second)
}
