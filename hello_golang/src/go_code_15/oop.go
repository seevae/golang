package main

import "fmt"

/*
	对比函数:
		A: 意义
			方法: 某个类别的行为功能, 需要指定的接收者调用
			函数: 一段独立功能的代码,可以直接调用
		B: 语法
			方法: 方法名可以相同,只要接收者不同
			函数: 命名不能冲突
		C: 变量作用域两者相同
*/

func main(){
	methodTest() //Go语言中的方法 , Go中既有方法又有函数
	//extends() //模拟继承关系
}

func methodTest() {
	//方法类似一个函数,但是他必须要有一个接收者,并且方法的调用受到限制,只能由接收者调用(包含了一个接收者的函数)
	//接收者可以是命名类型或者结构体类型的一个值或者是一个指针.所有给定类型的方法属于该类型的方法集
	w:=Worker{
		name:    "张三",
		age:     19,
		address: "西安",
	}
	w.work()  //类似于指定接收者, 方法属于指定的接收者. 这是一种值调用的形式

	//指针调用的形式
	w2:= &Worker{
		name:    "Alice",
		age:     20,
		address: "China",
	}
	w2.work()
	w2.rest()
}
type Worker struct {
	name string
	age int
	address string
}
func (w Worker) work(){
	fmt.Printf("%s 正在工作中\n",w.name)
}
//接收者指定为指针类型
func (w *Worker) rest(){  //w = w , w = w2的地址
	fmt.Println(w.name,"正在休息")
}


func extends() {
	/*
		oop:面向对象
		Java纯面向对象语言, Go不是纯面向对象语言,没有封装继承多态三大属性
	*/

	/**
	Go语言的结构体嵌套
		1.模拟继承性: is-a
			type A struct{
				field
			}
			type B struct{
				field
				A //匿名字段
			}
		2.模拟聚合关系: has -a
			type A struct{
				field
			}
			type B struct{
				field
				a A //非匿名字段, 无提升字段
			}
	*/

	stu1 := Student{
		Person:  Person{
			name: "小王",
			age: 18,
		},
		address: "北京",
	}
	fmt.Println(stu1)
	stu1.name = "小张" //提升字段结构体可以直接访问
	/*
		本来应该是 stu1.Person.name = "小张"
		但由于Person是Student中的一个匿名字段且也是一个结构体, 所以Student中的Person中的name和age叫做提升字段
		提升字段结构体对象可以直接访问, 所以写法如上所示
	*/
	fmt.Println(stu1)
}
//Go模拟实现继承 --> 使用结构体嵌套
type Person struct {
	name string
	age int
}
type Student struct {
	Person  //一个结构体中的字段是另一个结构体且该字段是匿名字段则这个字段称之为匿名字段, 匿名字段可以直接通过结构体来访问
	address string
}