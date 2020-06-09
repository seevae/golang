package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	利用缓冲区提高读写效率
	没有缓冲区的时候,文件的读写都是程序从磁盘中进行的,并且每一次读写都要访问磁盘,磁盘io效率是较低的
	有了缓冲区每次将读写的内容先放入缓冲区,
		如果是写操作则等到所有写操作完成后,一次性的写入磁盘,
		如果是读操作则下一次的读就不需要从磁盘进行了,直接访问缓冲区就可以了
*/

func main(){
	//bufio Reader读文件
	//filePath := "E:/test/test1/newfile4.txt"
	//file,err := os.OpenFile(filePath,os.O_RDONLY,os.ModePerm)
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//defer file.Close()
	//
	////创建Reader对象
	//p := make([]byte,1024)
	//b1 := bufio.NewReader(file)
	//n1,err := b1.Read(p)
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(n1)
	//fmt.Println(string(p[:n1]))

	//bufio Writer写文件
	filePath2 := "E:/test/test1/newfile5.txt"
	file2,err := os.OpenFile(filePath2,os.O_RDWR|os.O_CREATE,os.ModePerm)
	if err != nil{
		fmt.Println(err)
		return
	}
	w1 := bufio.NewWriter(file2)
	w1.WriteString("大家好,我是vae")
	w1.Flush() //刷新缓冲区, 当缓冲区的大小比写入数据大时,数据会先放在缓冲区而不是直接写入文件,此时就需要刷新缓冲区让其进入文件
	fmt.Println("写入完成")

}
