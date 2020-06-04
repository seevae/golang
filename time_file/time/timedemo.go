package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	//获取当前时间
	t1:=time.Now()
	fmt.Printf("%T\n",t1)
	fmt.Println(t1)

	//获取指定时间
	t2 := time.Date(2020,6,3,9,52,28,0,time.Local)
	fmt.Println(t2)

	//time--->string之间的转换
	s1 := t1.Format("2006年1月2日 15:04:05")  //go语言在转换时间时必须是从该日期算起,该日期据说是为了纪念go的诞生
	s2 := t1.Format("2006-1-2 15:04:05")
	fmt.Println(s1,s2)
	//如果只想获得年月日
	s3 := t1.Format("2006/01/02")
	fmt.Println(s3)

	//字符串转日期string--->time
	s4 := "1999年10月10日"
	t3,err := time.Parse("2006年1月2日",s4)
	if err!=nil{
		fmt.Println("错误信息:",err)
	}
	fmt.Printf("%T\n",t3)
	fmt.Println(t3)

	//直接将时间类型转换为string
	fmt.Println(t1.String())

	//根据当前时间获取指定的内容
	//根据当前时间获取年月日
	year,month,day := time.Now().Date()
	fmt.Println(year,month,day)
	//根据当前时间获取时分秒
	hour,min,second := time.Now().Clock()
	fmt.Println(hour,min,second)
	//也可以当个获取
	year2 := time.Now().Year()
	fmt.Println(year2)  //其他同理
	fmt.Println(time.Now().YearDay()) //表示今年一共过了多少天

	//时间戳: 指定的日期,距离1970年1月1日0点0时0分0秒的时间差值:秒,纳秒
	t4 := time.Date(1970,1,1,1,0,0,0,time.UTC)
	timeStamp1 := t4.Unix() //秒的差值
	fmt.Println(timeStamp1)

	//睡眠
	time.Sleep(3 *time.Second) //让当前程序进入睡眠状态,代表让程序睡眠3秒钟
	fmt.Println("main...over")
	//使用随机数睡眠1-10秒
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10)+1
	fmt.Println(randNum)
	time.Sleep(time.Duration(randNum)*time.Second)
	fmt.Println("over睡醒了")
}
