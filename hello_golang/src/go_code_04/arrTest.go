package main

import (
	"fmt"
)

/**
	数组练习
		go中的数据类型
			基本类型:整数,浮点,布尔,字符串
			复合类型:array,slice,map,struct,pointer,function,channel
 */
func main(){
	//数组的定义
	var arr1 [4] int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	fmt.Println(len(arr1))  //len()求容器的长度,是容器中实际存储数据的大小
	fmt.Println(cap(arr1)) //cap()求容器的容量,是容器中能够存储数据的最大值
	for i:=0;i< len(arr1);i++{
		fmt.Print(arr1[i]," ")
	}
	fmt.Println()

	//数组的其他创建方式
	var b = [4]int{1,2,3,4}
	var s = [5]string{"a","bb","cc"}
	fmt.Println(b,s)

	f:=[...]int{7,8,6,2,1,0,3,4}
	fmt.Println(f)

	//go语言中内村地址的分配是自动的. 数组的地址就是数组第一个元素的地址,即数组首地址
	fmt.Printf("%p\n",&f)

	//数组的遍历
	//使用for....range遍历数组,他取出来的是一对数据,分别是下标和数值
	for index,value:=range f{
		fmt.Printf("数组的下标是%d , 值是%d\n",index,value)
	}
	//如果只想获得值而不要下标值,则可以使用 _ 下划线代替
	sum:=0
	for _,value:=range f{
		sum = sum+value
	}
	fmt.Println(sum)

	//go中可以用一个数组给另一个数组初始化
	f2:=f
	fmt.Print(f2," ")
	fmt.Println(f2==f)  //==比较的是值

	//多维数组
	//二维数组的本质就是一维数组的集合,一维数组的每一个下标记录的又是一个数组的地址
	bb:=[3][4]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}  //三个一维数组的集合
	fmt.Println(bb)
	fmt.Printf("二维数组的地址是%p\n",&bb)
	fmt.Printf("二维数组的长度是%d\n",len(bb))
	fmt.Printf("第一个一维数组的长度是%d\n",len(bb[0]))
	fmt.Println(bb[0])
	fmt.Printf("第二个一维数组的第二个数是%d\n",bb[1][1])
	//使用for...range遍历二维数组
	for _,arr:=range bb{      //第一次遍历取出来的对象值还是数组,所以还要用for...range遍历一次
		for _,value:=range arr{
			fmt.Print(value," ")
		}
	}

	//bubbleSort(f)
}

func bubbleSort(arr [8]int){
	for i:=0;i< len(arr);i++{
		for j:=0;j< len(arr)-i-1;j++{
			if arr[j]>arr[j+1]{
				temp:=arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	for _,v := range arr{
		fmt.Print(v," ")
	}
}
