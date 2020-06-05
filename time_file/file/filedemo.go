package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
)

func main(){
	//test1() //fileinfo文件操作
	//test2() //创建目录(路径)操作 filePath
	//test3() //创建文件操作
	//test4() //io操作 读
	//test5() //io操作 写

	//拷贝文件
	srcFile := "E:/test/test1/newfile.txt"
	destFile := "E:/test/test1/newfile4.txt"
	//srcFile := "E:/test/test1/pic1.jpg"
	//destFile2 := "E:/test/test3/pic2.jpg"
	copyFile3(srcFile,destFile)
}

//使用ioutil包中的方法
func  copyFile3(srcFile,destFile string)(int,error){
	arr,error := ioutil.ReadFile(srcFile)
	if error != nil{
		fmt.Println("出现错误:",error)
		return 0,nil
	}
	error = ioutil.WriteFile(destFile,arr,os.ModePerm)
	if error != nil{
		return 0,nil
	}
	fmt.Println("拷贝完成")
	return len(arr),nil
}

//使用io包中的Copy方法(推荐)
func copyFile2(srcFile,destFile string)(int64,error){
	srcF,error := os.Open(srcFile)
	if error != nil{
		fmt.Println("发生错误----1")
		return 0,error
	}
	destF,error := os.OpenFile(destFile,os.O_CREATE|os.O_RDWR,os.ModePerm)
	if error != nil{
		fmt.Println("发生错误----2")
		return 0,error
	}
	defer srcF.Close()
	defer destF.Close()
	return io.Copy(destF,srcF)
}

//使用file对象的read,write进行读写
func copyFile(srcFile,destFile string)(int,error) {
	//1.打开文件
	src,error := os.Open(srcFile)
	if error != nil{
		fmt.Println("有错误,返回-->1")
		return 0,error
	}
	dest,error := os.OpenFile(destFile,os.O_RDWR|os.O_CREATE,os.ModePerm)  //注意第二个参数指定系统操作文件的权限,此处为可读可写可创建
	if error != nil{
		fmt.Println("err:",error)
		fmt.Println("有错误,返回-->2")
		return 0,error
	}
	//3.关闭连接
	defer src.Close()
	defer dest.Close()
	//2.读写
	arr := make([]byte,1024,1024)
	total := 0
	n := -1
	for{
		n,error = src.Read(arr)
		if n==0 || error == io.EOF{
			fmt.Println("拷贝完毕,准备粘贴...")
			break
		}else if error != nil{
			fmt.Println("出现错误")
			return total,error
		}
		total = total+n
		dest.Write(arr)
	}
	fmt.Println("拷贝完成")
	return total,nil
}

func test5() {
	//1.打开文件
	file,err := os.OpenFile("E:/test/test1/newfile2.txt",os.O_RDWR,os.ModePerm) //需要进行写操作,所以要使用该方法
	//3.关闭连接
	defer file.Close()
	//2.写出数据
	if err != nil{
		fmt.Println(err)
	}
	arrWrite := []byte{97,98,99,100}
	n,err2 := file.Write(arrWrite)
	if err != nil{
		fmt.Println(err2)
	}
	fmt.Println(n)
	//以字符串的形式直接输出
	n1,err3 := file.WriteString("你好啊可达鸭")
	if err3 != nil{
		fmt.Println(err3)
	}
	fmt.Println(n1)
}

func test4() {
	//1.打开文件
	file,err := os.Open("E:/test/test1/newFile.txt")
	//3.关闭连接
	defer file.Close()  //习惯上先写关闭连接使用defer关键字最后执行
	//2.读取数据
	if err != nil{
		fmt.Println(err)
		return
	}
	readArr := make([]byte,4,4)
	//n,err2 := file.Read(readArr)  //n返回值为读取到的长度,单位字节
	//if err2 != nil{
	//	fmt.Println(err2)
	//}
	//fmt.Println(n)
	//fmt.Println(string(readArr)) //没有读完
	////第二次读取
	//n1,err3 := file.Read(readArr)
	//fmt.Println(err3)
	//fmt.Println(n1)
	//fmt.Println(string(readArr))
	//...

	//正确读取的做法
	for {
		n,err := file.Read(readArr)
		if n==0 || err == io.EOF {
			fmt.Println()
			fmt.Println("读操作已完成")
			break
		}
		fmt.Print(string(readArr[:n]))
	}
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

func errorCheck(err error){
	if err != nil{
		log.Fatal(err)
	}
}