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

type Student2 struct{
	name string
	age int
	book *Book
}

type Student struct {
	name string
	age int
	book Book
}

type Book struct {
	name string
	price string
}

type Worker struct{
	string  //匿名字段,默认使用数据类型作为名字,由于是使用类型名字区分字段名称,所以类型不能重复
	//string  //再来个string 就报错了
	int
}

func main(){
		nestStruct() //结构体嵌套
	/*
		nonameStruct() //匿名结构体 && 结构体中的匿名字段
		structPointer() //结构体指针
		structInit() //结构体的定义
	*/
}

func nestStruct() {
	student1 := Student{
		name: "小花",
		age:  18,
		book: Book{
			name:"十万个为什么",
			price:"35.5",
		},
	}
	fmt.Println(student1)
	student1.book.name = "水浒传"
	fmt.Println(student1)

	fmt.Println("---------------------------")

	book2 := Book{"哈利波特","50"}
	student2 := Student{
		name: "小飞侠",
		age:  20,
		book: book2,
	}
	fmt.Println(student2)
	/**
		结构体都是值传递, 在定义 book: book2时实际上是复制了一个副本
		所以在后面改变了student2.book.book2 的值时, student2中book的副本发生了变化但是book本身没有改变
		注意: 一般情况下结构体嵌套直接传地址, 这样可以直接改变结构体中字段的值
	*/
	student2.book.name = "名侦探柯南"
	fmt.Println(student2)
	fmt.Println(book2)

	fmt.Println("========================")
	book1:=Book{
		name: "解忧杂货铺",
		price: "35.9",
	}
	stu :=Student2{
		name: "小飞哥",
		age:  19,
		book: &book1,
	}
	fmt.Printf("学生: %s , 年龄: %d, 书籍名称: %s, 书籍价钱: %s\n",stu.name,stu.age,stu.book.name,stu.book.price)
	stu.book.name = "可可西里的美丽传说"
	fmt.Printf("学生: %s , 年龄: %d, 书籍名称: %s, 书籍价钱: %s\n",stu.name,stu.age,stu.book.name,stu.book.price)
	fmt.Printf("book1.name = %s, book1.price = %s\n",book1.name,book1.price)
}


func nonameStruct(){
	/*
		匿名结构体: 没有名字的结构体,在创建匿名结构体时,同时创建对象
		匿名字段: 一个结构体的字段没有字段名
	*/
	s1 := Student{age:18,name:"Lily"}
	fmt.Println(s1)

	//匿名函数
	func (){
		fmt.Println("我是匿名函数")
	}()  //()代表调用

	//匿名结构体
	s2 := struct {
		name string
		address string
	}{
		name:"李四",
		address:"新疆",
	}
	fmt.Println(s2.name,s2.address)

	//结构体匿名字段
	w1:=Worker{"小明",18}
	fmt.Println(w1.string,w1.int)
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