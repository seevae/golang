package main

import (
	"fmt"
)

func main() {
	/*
		深浅拷贝
		深拷贝:拷贝的是数据本身 -->值类型的数据,默认都是深拷贝:array,int,float,string,bool,stuct
		浅拷贝:拷贝的是数据的地址 -->导致多个变量指向同一块内村,引用类型的数据都是浅拷贝,slice,map
				因为切片是引用类型,直接拷贝的是地址
	*/
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := make([]int, 0, 0)
	for i := 0; i < len(s1); i++ {
		s2 = append(s2, s1[i])
	}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("%p,%p\n", s1, &s1) //第一个是引用的地址, 第二个是数组的真实地址
	fmt.Printf("%p,%p\n", s2, &s2)

	//copy函数
	s3 := []int{11, 12, 13}
	fmt.Println(s3)
	copy(s3, s2) //第一个参数是目标数组,第二个参数是被拷贝的数组
	fmt.Println(s3)
	copy(s3[1:], s2[2:])
	fmt.Println(",,", s3)

	//map的初步使用
	//创建map的三种方式
	var map1 map[int]string //nil map,无法直接使用  需要使用make重新定义
	var map2 = make(map[int]string)
	var map3 = map[string]int{"Go": 98, "Java": 98, "Python": 98}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)
	//只定义的map为nil 不能直接使用 ,需要使用make初始化一下
	if map1 == nil {
		map1 = make(map[int]string)
	}
	map1[1] = "hehe"
	map1[2] = "haha"
	map1[3] = "heihei"
	map1[4] = ""
	fmt.Println(map1)

	//根据key值获取value直接map1[下标]就好  如果下标不存在则默认返回value类型的默认值 ,如string类型返回空字符串
	/**如何区分返回的值是默认的值还是map中已经存放的值, 例如map1中放入一个空字符串,此时取map1[40]还是返回一个空字符串
		那究竟是哪个字符串呢,此时就要用到两个返回值来做判断
	**/
	v1, ok := map1[40]
	if ok == true {
		fmt.Println("key存在,value=", v1)
	} else {
		fmt.Println("key不存在,返回的是默认值:", v1)
	}
	//map的长度计算  len(map) --> 求出来的就是map的长度及对数
	len1 := len(map1)
	fmt.Println(len1)
	//删除数据
	delete(map1, 4) //key值存在直接删除.key值不存在删除失败
	fmt.Println(map1)

	//map的遍历
	//for..range遍历: 数组中获取的是下标和对应的值,map中获取的是key和对应的value
	for key, value := range map1 {
		fmt.Println(key, value) //打印结果是无序的.
	}
	//有序的遍历map,for循环
	for i := 1; i <= len(map1); i++ {
		fmt.Println(i, map1[i])
	}

	/**
	1.创建map用于存储人的信息
		name,age,sex,address
	2.每个map存储一个人的信息
	3.将这些map放入一个slice中
	4.打印遍历输出
	*/
	peopleSlice := make([]map[string]string, 0, 0)
	map11 := make(map[string]string)
	map11["name"] = "张益达"
	map11["age"] = "12"
	map11["sex"] = "男"
	map11["address"] = "北京"

	map12 := make(map[string]string)
	map12["name"] = "李菲菲"
	map12["age"] = "12"
	map12["sex"] = "男"
	map12["address"] = "北京"

	map13 := make(map[string]string)
	map13["name"] = "赵小雨"
	map13["age"] = "13"
	map13["sex"] = "女"
	map13["address"] = "北京"

	peopleSlice = append(peopleSlice, map11,map12,map13)
	for i,value:=range peopleSlice{
		fmt.Printf("第 %d 个人的信息是:\n",i+1)
		fmt.Printf("\t姓名:%s\n",value["name"])
		fmt.Printf("\t年龄:%s\n",value["age"])
		fmt.Printf("\t性别:%s\n",value["sex"])
		fmt.Printf("\t地址:%s\n",value["address"])
	}

	//map也是引用类型


}
