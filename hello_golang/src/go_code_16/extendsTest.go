package main

import "fmt"

func main(){
	//创建父类对象
	s1:= Person{"张三",20}
	fmt.Println(s1)

	//创建子类对象
	s2:=Student{
		Person: Person{"李斯",22},
		school: "大同路小学",
	}
	fmt.Println(s2)

	//提升字段:如果A结构体中的匿名字段(非匿名字段为聚合关系而不是继承关系)是B结构体,那么B结构体中的字段就叫做提升字段,可以通过结构体A的对象直接访问
	var s3 Student
	s3.Person.name = "小六"
	s3.name = "小五" //两种定义方式效果都相同,因为name字段是一个提升字段
	//子类调用(继承)父类方法
	s3.eat()
	s3.study()
	//子类重写eat()方法后直接调用
	s3.eat()
	//子类重写eat*()方法后依然调用父类中的方法
	s3.Person.eat()
}

//1.定义"父类"
type Person struct {
	name string
	age int
}

//2.定义子类
type Student struct {
	Person //模拟继承结构
	school string //子类新增属性
}
//父类方法
func (p *Person) eat(){
	fmt.Println("父类的方法,吃窝窝头")
}
//子类方法
func (s *Student) study(){
	fmt.Println("子类的方法:学习")
}

//子类覆写父类方法
func (s *Student) eat(){
	fmt.Println("子类覆写父类方法,吃肉肉了")
}
