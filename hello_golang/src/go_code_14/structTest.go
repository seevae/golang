package main

import "fmt"

/*
结构体: 是由一系列具有相同类型或不同类型的数据构成的数据集合
	结构体成员是由一系列的成员变量构成,这些成员变量也被成为字段
	成员变量如果是大写表示对外可见,小写表示对外不可见,对外可见则只要导入这个包就可以使用该成员
*/

//定义结构体
type Person struct {
	name string
	age int
	sex string
	address string
}

func main(){
	structInit() //结构体的定义
	structPointer() //结构体指针
}

func structPointer() {
	//1.证明结构体是值类型
	p1 := Person{"小王",19,"女","天津"}
	fmt.Println(p1)
	fmt.Printf("%p,%T\n",&p1,p1)

	p2 := p1
	p2.name = "李大哥"
	fmt.Println(p1,p2)  //由此可以看出 结构体是一种值类型的数据 他的传递为值传递 赋值时发生的是深拷贝

	//2.定义结构体指针
	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p,%T\n",&pp1,pp1)
	fmt.Println(*pp1)
	//如果想要改变某个属性值
	(*pp1).name = "诸葛菲菲"
	//也可以简写如下:
	pp1.name = "慕容藏藏"
	fmt.Println(*pp1,p1) //发现通过指针修改了p1对象对应的值

	//3.使用内置函数new()，go语言中专门用于创建某种类型的指针的函数
	pp2 := new(Person)
	pp2.name = "hanhan"
	pp2.age = 21
	fmt.Println(*pp2)
}

func structInit(){
	//1.结构体的初始化
	var p1 Person
	fmt.Println(p1)
	p1.name = "哈哈怪"
	p1.age = 22
	p1.sex = "男"
	p1.address = "北京"
	fmt.Println(p1)

	//2.方法二
	p2 := Person{}
	p2.address = "杭州"
	p2.age = 20
	p2.name = "翠花"
	p2.sex = "女"
	fmt.Println(p2)

	//3.方法三
	p3 :=Person{
		name:    "二狗子",
		age:     19,
		sex:     "女",
		address: "铜川",
	}
	fmt.Println(p3)

	//4.方法四
	p4 := Person{"小王",19,"女","天津"}
	fmt.Println(p4)
}