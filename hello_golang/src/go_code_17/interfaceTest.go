package main

import "fmt"

/*
Go语言中的接口:
定义使用 interface关键字 ,与Java相比不同之处在于Go中接口和实现类的关系是非侵入式的
例如 Java中实现一个类需要显示的使用implements关键字指明关系
class A implements InterfaceA{}
而Go中不需要使用什么关键字来显示指明,一个类实现了接口的所有功能就代表实现了该接口
*/

func main(){
	test1() //接口的初步认识
	test2() //空接口
	test3() //接口的嵌套
}

func test3() {
	//如果想要实现接口D,则也要实现接口B,C因为B,C在D中嵌套着
	c := Cat{}
	c.testB()
}

func test2() {
	m := Mouse{name:"小鼠标"}
	k := Keyboard{name:"小键盘"}
	s := "小字符串"
	n := 200
	testI(m) //由于testI所需参数是空接口类型所以任何类型都可以当作参数
	testI(k)
	testI(s)
	testI(n)

	//制作一个可以存放任何类型的数组
	arr := []interface{}{}  //可以存放任何类型的数组
	arr = append(arr, 1)
	arr = append(arr,"字符串")
	fmt.Println(arr)


}

//接口的初步认识
func test1() {
	var m = Mouse{"罗技炫彩"}
	m.light() //该方法是Mouse实例特有的, 所实现的接口类型是不具有该方法的
	//u1.start()
	//u1.end()
	testUSB(m)
	var k = Keyboard{"好键盘"}
	testUSB(k)

	//也可以创建单独的接口类型
	var u USB
	u = m
	u.start()
}

//定义接口
type USB interface {
	start()
	end()
}

//定义结构体
type Mouse struct{
	name string
}
type Keyboard struct {
	name string
}
//结构体实现接口
func (m Mouse) start(){
	fmt.Println(m.name+"开始工作")
}
func (m Mouse) end(){
	fmt.Println(m.name+"工作结束")
}
func (m Mouse) light(){
	fmt.Println(m.name+"开始发光")
}

func (k Keyboard) start(){
	fmt.Println(k.name+"开始工作")
}
func (k Keyboard) end(){
	fmt.Println(k.name+"结束工作")
}

//测试方法
func testUSB(usb USB) {
	usb.start()
	usb.end()
}

//-------------空接口---------------------------
//定义一个空接口 , 由于是一个空接口里面没有方法,所以可以默认所有的类型都实现了他
type A interface {

}
func testI(a A){
	fmt.Println(a)
}

//------------接口的嵌套------------
//接口的嵌套实际也就是模拟实现"多接口的继承"
type B interface {
	testB()
}
type C interface {
	testC()
}
type D interface {
	B
	C
	testD()
}
type Cat struct {

}
//只有把D接口和他所嵌套的接口方法都实现才是真正实现了D接口
func (c Cat) testB(){
	fmt.Println("testB()....")
}
func (c Cat) testC(){
	fmt.Println("testC()....")
}
func (c Cat) testD(){
	fmt.Println("testD()....")
}

