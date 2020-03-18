//+build ignore
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//随机数的生成random
	num1 := rand.Int()
	fmt.Println(num1)
	//生成[0,n)的随机数
	for i:=0;i<10;i++{
		num:=rand.Intn(10) //规定生成随机数的范围
		fmt.Print(num," ")
	}
	//只有以上代码会发现每次运行生成的随机数都是相同的,原因是随机数种子是默认的,一直没有改变
	//重新设置种子
	rand.Seed(100)
	num2 :=rand.Int()  //随着种子数值的改变num2的值也会跟着改变
	fmt.Println(num2)
	//怎么样连续生成不一样的数字呢?-->让随机数种子一直变化 --> 所以让随机数种子随着时间变化就可
	//时间函数
	t1:=time.Now()  //打印现在的时间
	fmt.Println(t1)
	fmt.Printf("%T\n",t1)
	//时间戳:指定时间,距离1970年1月1日0点0分0秒之间的时间差值:秒,纳秒
	timeStamp1:=t1.Unix()   //犹豫是距上述事件的差值,所以一直都在变化,所以生成的随机数是一直变化的
	fmt.Println(timeStamp1)

	/**真正生成随机数
		1.设置种子数,可以设置成时间戳
		2.调用生成随机数的函数
	*/
	rand.Seed(time.Now().Unix())
	for i:=0;i<10;i++{
		fmt.Println("-->",rand.Intn(100))
	}
	//获取指定范围的随机数 例如[3,48] --> 那么就先获取0-45的数字再加3  [15,76] --> 76-15 -->[0,61]+15
	num3:=rand.Intn(46)+3 //Intn函数是左闭右开的
	fmt.Println(num3)

}
