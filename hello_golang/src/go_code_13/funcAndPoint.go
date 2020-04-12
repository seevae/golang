package main

import "fmt"

func main() {
	/**
	  函数指针 && 指针函数
	  函数指针: 函数的指针,指向了一个函数的指针
	  	go语言中 ,函数默认就看作是一个指针 , 不用带 *
	  	(map,slice,function 都是引用类型,指向的都是一个地址)
	  指针函数: 一个函数,该函数的返回值是一个指针
	*/
	arr1:=fun2()
	fmt.Printf("arr1的类型: %T,地址: %p,数值: %v\n",arr1,&arr1,arr1)

	arr2 :=fun3()
	fmt.Printf("arr2的类型: %T,地址: %p,数值: %v\n",arr2,&arr2,arr2)


	//----------- 指针作为参数 ----------------
	/*
		参数的传递: 值传递, 引用传递

	*/
	a := 10
	fmt.Println(a)
	fun4(a)  //值传递
	fmt.Println(a)

	fun5(&a)  //指针只对于值类型(基本类型)的数据使用 , 引用类型就不需要多此一举了
	fmt.Println(a)

	arr3:=[]int{4,5,6,7,8,9}
	fun6(arr3)
	fmt.Println(arr3)

}

func fun6(p1 []int){
	p1[0] = 600
	p1 = append(p1, 500)
}

func fun5(pnum *int){
	*pnum = 55
}

func fun4(num int){
	fmt.Println(num)
	num = 100
	fmt.Println(num)
}

func fun3()*[4]int{
	arr:=[4]int{1,2,3,4}
	return &arr
}

func fun2()[4]int{
	arr:=[4]int{1,2,3,4}
	return arr
}

func fun1() {
	fmt.Println("heh")
}
