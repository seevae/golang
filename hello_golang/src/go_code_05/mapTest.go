package main

import "fmt"

func main(){
	/*
		深浅拷贝
		深拷贝:拷贝的是数据本身 -->值类型的数据,默认都是深拷贝:array,int,float,string,bool,stuct
		浅拷贝:拷贝的是数据的地址 -->导致多个变量指向同一块内村,引用类型的数据都是浅拷贝,slice,map
				因为切片是引用类型,直接拷贝的是地址
	*/
	s1:=[]int{1,2,3,4,5,6}
	s2 := make([]int,0,0)
	for i:=0;i< len(s1);i++{
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("%p,%p\n",s1,&s1)  //第一个是引用的地址, 第二个是数组的真实地址
	fmt.Printf("%p,%p\n",s2,&s2)

	//copy函数
	s3:=[]int{11,12,13}
	fmt.Println(s3)
	copy(s3,s2)  //第一个参数是目标数组,第二个参数是被拷贝的数组
	fmt.Println(s3)
	copy(s3[1:],s2[2:])
	fmt.Println(",,",s3)

	//map的初步学习
	
}
