package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main(){
	//test1() //fileinfo文件操作
	//test2() //创建目录(路径)操作 filePath
	//test3() //创建文件操作
	test4() //io操作
}

func test4() {

}

func test3() {
	//create: 如果当前文件不存在,直接创建文件;如果当前文件存在,创建并覆盖原有重名文件
	file,err := os.Create("E:/test/test1/newFile2.txt")  //以绝对路径的方式创建文件
 	if err != nil {
		fmt.Println("创建文件失败",err)
		return
	}
	fmt.Println("创建文件成功")
	fmt.Printf("%T\n",file)
	//以相对路径的方式创建文件,以工程目录为参照
	file2,err2 := os.Create("bb.txt")
	if err2 != nil{
		fmt.Println("创建文件失败",err2)
	}
	fmt.Println(file2)

	//打开文件
	//open 以只读模式打开文件,只能读文件,不能进行其他操作
	file3,err3 := os.Open("E:/test/test1/newFile2.txt")
	if err3 != nil{
		fmt.Println(err3)
		return
	}
	fmt.Println("文件打开成功",file3)
	//openfile:可读可写 ,三个参数 : 文件路径,打开模式,文件权限
	file4, err4 := os.OpenFile("E:/test/test1/newFile.txt",os.O_RDWR,os.ModePerm)
	if err4 != nil{
		fmt.Println(err4)
	}
	fmt.Println("文件打开成功",file4)

	//关闭文件连接: 所有打开的文件都和程序建立了连接,使用完毕后应当关闭所有文件连接
	file3.Close()
	file4.Close()

	//删除文件或文件夹
	//err5 := os.Remove("E:/test/test3/demo/hello.txt")
	//if err5 != nil {
	//	fmt.Println(err5)
	//}
	//fmt.Println("删除文件成功")

	//删除空目录也可以使用Remove方法
	err6 := os.Remove("E:/test/test2")
	if err6 != nil {
		fmt.Println(err6)
	}
	fmt.Println("目录删除成功")
	//同时删除多级目录(非空目录) RemoveAll , 不包括最上级的父目录,eg:/test
	err7 := os.RemoveAll("E:/test/test3/demo")
	if err6 != nil {
		fmt.Println(err7)
	}
	fmt.Println("全部删除成功")

}

func test2() {
	/*
	文件路径
		1.相对路径
			相对于当前目录而言的
		2.绝对路径
			目录全路径
			E:/test/test1/newFile.txt
			E:\\test\\test1\\newFile.txt
	*/
	fileName1 := "E:/test/test1/newFile.txt"
	fileName2 := "ab.txt"
	//判断是不是绝对路径
	fmt.Println(filepath.IsAbs(fileName1))
	fmt.Println(filepath.IsAbs(fileName2))
	//返回绝对路径
	fmt.Println(filepath.Abs(fileName1))
	fmt.Println(filepath.Abs(fileName2)) //E:\go\src\time_pk\ab.txt
	//获取父目录
	fmt.Println("获取父目录:",path.Join(fileName1,".."))

	//创建目录
	err := os.Mkdir("E:/test/test2",os.ModePerm)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("目录创建成功")
	//创建多级目录
	err2 := os.MkdirAll("E:/test/test3/demo",os.ModePerm) //递归创建多级目录
	if err != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("多级目录创建成功")
}

func test1() {
	fileInfo, err := os.Stat("E:/test/test1/newFile.txt")
	if err!= nil{
		fmt.Println(err)
	}
	fmt.Printf("%T\n",fileInfo)
	//有了fileInfo对象就可以操作该对象对应的文件了
	fmt.Println("文件名称是",fileInfo.Name())
	fmt.Println("文件大小是",fileInfo.Size())  //文件大小,字节为单位
	fmt.Println("修改时间是",fileInfo.ModTime())
	fmt.Println("权限：",fileInfo.Mode())
	fileInfo.Size()
	if fileInfo.IsDir() {
		fmt.Println("是一个目录")
	}else{
		fmt.Println("是一个文件")
	}
}

