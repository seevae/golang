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
	fmt.Println()
	fmt.Println("-----------------------------")

	/**切片 --->slice
		所谓切片就是一个动态数组,定义的方式与数组也基本类似,只是定义时不需要指定长度
	*/
	//当不指定切片长度与容量时
	var s1 []int  //此时默认切片长度为0,容量为0,直接通过下标是不能赋值的
	fmt.Println(s1)
	//s1[0] = 1
	//s1[1] = 3
	s1 = append(s1, 1,2,3)  //切片的底层实现实质上是创建了一个新的数组将原先的内容拷贝了过去
	fmt.Println(s1)

	//指定切片的长度与容量
	var w2 = []int{1,2,3,4}  //此种方式切片默认就是{}中的内容
	fmt.Println(w2)
	//使用make创建切片
	//var s2 []int = make([]int,3,8)  //简写如下
	s2:=make([]int,3,8) //创建切片的三个参数 make([]type,len,cap)
	fmt.Println(s2)

	//可以使用一个数组来给切片赋值
	arr:=[]int{1,2,3,4}
	s3 := arr[1:3]  // 冒号":"前后代表所复制数组的起止位置 是一个前开后闭的区间,此数下标[1,3)
	fmt.Println(s3)
	//append的俩个用法
	//1.给切片后添加指定元素
	s3 = append(s3, 11,22,33)
	fmt.Println(s3)
	//2.将一个切片中的所有元素添加到另一个切片后面
	s3 = append(s3,arr...)
	fmt.Println(s3)
	//切片的遍历和数组相同


	a:=[10]int{1,2,3,4,5,6,7,8,9,10}
	slice1 := a[:5]
	slice2 := a[3:7]
	fmt.Println(",,",a)
	fmt.Printf("%p\n",slice1)
	fmt.Printf("%p\n",&a)
	fmt.Printf("%p\n",slice2)
	slice1 = append(slice1, 1,2,3,4)
	/**注意:如果新加元素长度小于原切片大小,则不不触发扩容而是直接在原数组中添加,所以此处a数组中的内容也会改变*/
	fmt.Println(slice1)
	fmt.Println(a)
	fmt.Printf("%p\n",slice1)
	fmt.Printf("%p\n",&a)
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
