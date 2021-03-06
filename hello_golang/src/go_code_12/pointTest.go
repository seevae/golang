package main

import "fmt"

func main(){
	//指针 pointer :存储了另一个变量的内村地址的变量
	//基本用法
	a:=10
	fmt.Printf("a的地址是 : %p\n",&a)  //不同于引用类型打印引用就是引用所指向的地址,基本类型要用取地址符获取地址

	//创建一个指针用来存储a的地址
	var p1 *int
	fmt.Println(p1)
	p1=&a
	fmt.Println("p1的值是:",p1)
	fmt.Printf("p1自己的地址是:%p\n",&p1)
	fmt.Println("p1的数值是a的地址,该地址中存储的数据是:",*p1) //获取指针指向变量的数值
	//通过指针来改变变量的值
	*p1 = 200   //*p1记录的是a在内村中的地址中的值
	fmt.Println(a)

	//指针的指针  二级指针
	var p2 **int
	p2 = &p1
	fmt.Printf("p2中存放的内存地址:%p\np2自己的内存地址:%p\n",p2,&p2)
	fmt.Println("p2中存储的数值就是p1的地址,对应的*p2数据就是p1中的值,即p1记录的a的内存地址:",*p2)
	fmt.Println(**p2) //如果想通过p2拿到a的值,则需要使用二级指针获取数据的方式 **p2


	/*指针数组和数组指针
		指针数组 --> 指针的集合  --> [4]*int
		数组指针 --> 数组的指针  --> *[4]int
	*/


	//数组指针:
	//1.创建一个普通的数组
	arr1 := [4]int{1,2,3,4}
	fmt.Println(arr1)
	//2.创建一个指针,存储该数组的地址 --> 数组指针
	var p11 *[4]int = &arr1
	fmt.Printf("%p\n",&arr1)  //数组的地址
	fmt.Printf("%p\n",p11)	//指针中存放的是指向变量的地址
	fmt.Printf("%p\n",&p11)	//&p11 取出的是p11自己本身的地址
	//根据数组指针,操作数组
	(*p11)[0] = 100
	fmt.Println(arr1)
	//简化写法
	p11[0] = 200
	fmt.Println(arr1)

	//指针数组:是一个数组,里面装的都是指针类型的变量
	a,b,c,d:=1,2,3,4
	arr11:=[4]int{a,b,c,d}
	pArr1:=[4]*int{&a,&b,&c,&d}   //指针数组,数组类型为指针,存放的都是变量地址
	fmt.Println(pArr1)
	//通过指针数组修改数组中元素的值
	*pArr1[0] = 200 //pArr[0]中存放的是a的地址,通过 * 号来拿到地址中的内容并进行修改
	fmt.Println(a) //a被改变了
	//还需要注意一个现象
	//上面通过*pArr1[0] 改变了a的值,是因为0下标处存放的就是a的地址,通过地址改值值改变
	//下面修改了4的值,但是在arr11数组中存放的abcd的值是通过值传递过来的,并不会跟着发生变化
	c=400
	fmt.Println(arr11)






}
