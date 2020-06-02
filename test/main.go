package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //只需要使用这个包中的init函数进行数据库的连接其他函数不用,所以用下划线导入
	"test/person"
	"test/pk1"
	//"fmt"
	//"test/person"
	"test/pk2"
	"test/pk3"
	//tu "test/utils/timeUtils"
	//"test/utils"  //以绝对路径的方式导入,还可以以相对路径的方式导入,如下:
	"test/utils" //以相对路径的方式导入,表示当前目录下的某个包,与上方结果一致
)

func main(){
	//test1() //包的初步学习

	//可以看到在main函数运行前包就加载了进来,并且先执行了init()函数
	//fmt.Println("hello world")
	//test2() //init()函数

	test3() //外部包的下载
}

func test3() {
	db,err := sql.Open("mysql","root:zhang@tcp(127.0.0.1:3306)/db1?charset=utf8")
	if err!=nil{
		fmt.Println("错误信息:",err)
		return
	}
	fmt.Println("连接成功",db)
}

func test2() {
	utils.Count()
}

func test1() {
	//tu.PrintTime()  //包起别名后调用
	//utils.Count()
	pk1.Mytest1()
	pk2.TestMethod()
	pk3.MyTest3()
	utils.SayHello()

	p1 := person.Person{"张三",18,"男"}
	fmt.Println(p1)
}
