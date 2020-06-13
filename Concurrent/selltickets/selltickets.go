package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	name string
}

var tickets = 10

var mutextLock sync.Mutex //创建一把互斥锁

var wg sync.WaitGroup //创建同步等待组防止主goroutine结束影响子goroutine

func main(){
	p1 := Person{name:"黄牛A"}
	p2 := Person{name:"黄牛B"}
	p3 := Person{name:"黄牛C"}
	wg.Add(3)
	go sellTickets(p1)
	go sellTickets(p2)
	go sellTickets(p3)
	wg.Wait() //主gorotine阻塞等待子gorotine执行完毕
	fmt.Println("程序执行完毕...")
}

func sellTickets(p Person) {
	defer wg.Done()
	for{
		mutextLock.Lock()
		if tickets>0{
			fmt.Printf("%s前来买票,当前还有票: %d张\n",p.name,tickets)
			time.Sleep(1*time.Second)
			tickets--
			fmt.Printf("%s买票结束,剩余票数%d\n",p.name,tickets)
		}else{
			mutextLock.Unlock() //如果票已经卖完,此时也是带锁的,需要解锁
			fmt.Println("当前已无车票,择日购买")
			break
		}
		mutextLock.Unlock()
	}
}

