package main

import (
	"fmt"
	"sync"
	"time"
)

var rlock *sync.RWMutex
var wg sync.WaitGroup
//读写锁
func main(){
	rlock = new(sync.RWMutex)
	//go readTest(1) //读操作
	//go readTest(2)
	go writeTest(1) //写操作
	go writeTest(2)
	wg.Add(2)
	wg.Wait()
}

//读操作没有顺序,并行
func readTest(i int){
	//上读锁
	rlock.RLock()
	for j:=1;j<6;j++{
		fmt.Println(i,"开始读操作...")
		time.Sleep(1*time.Second)
	}
	rlock.RUnlock()
	wg.Done()
}

//写操作拿到的是互斥锁,串行进行
func writeTest(i int){
	//上写锁
	rlock.Lock()
	for j:=1;j<6;j++{
		fmt.Println(i,"开始写操作...")
		time.Sleep(1*time.Second)
	}
	rlock.Unlock()
	wg.Done()
}