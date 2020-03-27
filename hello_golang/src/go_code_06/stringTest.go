package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	/**
	string实际上是字节的集合,在go中属于基本类型
	*/
	s1 := "hello中国" //一个汉字占三个字节
	s2 := "hello world,hello go"

	len1 := len(s1)
	len2 := len(s2)
	fmt.Println(len1, len2)

	//字符串实际上是字节的集合,可以通过下标进行访问单个字符
	fmt.Printf("s1的第一个字母是 : %c \n", s1[0])
	//如果直接打印,则打印的是字母的ascii码
	fmt.Println(s1[0]) //104

	//字符串的遍历
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%c ", s2[i])
	}
	for i, v := range s2 {
		fmt.Printf("下标是:%d 字符是:%c", i, v)
	}

	//字符串是字节的集合
	slice1 := []byte{65, 66, 67, 68, 69}
	s3 := string(slice1)
	fmt.Println(s3)
	s4 := "whatwhy"
	slice2 := []byte(s4) //字符串转字节数组
	fmt.Println(slice2)

	//strings包的使用 几个常用的api
	//是否包含指定内容
	s5 := "hell"
	fmt.Println(strings.Contains(s1, s5))
	//任意包含
	fmt.Println(strings.ContainsAny(s5, "eaaa")) // 只要指定字符串中有一个包含就行 e
	//统计substr在s中出现的此数
	fmt.Println(strings.Count(s2, "he"))
	//判断字符串是否以某个字符串开头
	s6 := "2020年的特殊文件.txt"
	if strings.HasPrefix(s6, "2020") {
		fmt.Println("这是2020的文件")
	}
	//判断字符是否以某个字符串结尾
	if strings.HasSuffix(s6, ".txt") {
		fmt.Println("这是一个txt文件")
	}
	//查找字符在字符串中首次出现的位置(IndexLast 是查找最后一次出现的位置),不存在的话返回-1
	fmt.Println(strings.Index(s2,"ll"))
	//字符串拼接
	ss1:=[]string{"abc","world","hello","hi"}
	s7:=strings.Join(ss1,"&")
	fmt.Println(s7)
	//字符串切分
	s8:="123,2313,asdad,1,22"
	ss2 := strings.Split(s8,",")
	for _,v:=range ss2{
		fmt.Println(v)
	}
	//字符串自身拼接
	s9:=strings.Repeat("hello",5) //自身拼接五次
	fmt.Println(s9)
	//字符串替换 四个参数分别为: 被替换字符串,被替换字母,替换字母,替换几个,如果传入-1,则替换所有
	s10:=strings.Replace(s2,"l","*",1)
	fmt.Println(s10)
	//大小写转换,ToLower和ToUpper
	//字符串截取 和切片相同,
	fmt.Println(s2)
	s12:=s2[:5]
	fmt.Println(s12)

	/**
		字符串和基本类型的转换 strconv包
			java中可以使用+号直接连接字符串和整型 如打印 "aaa"+1 --> aaa1
			go中是不行的 , go中 +号 两边类型要一致
		strconv包提供了这样的功能 eg:Atoi-->字符串变整型, Itoa -->整型变字符串
			其他的转换功能基本都是 ParseXxx  返回值都有两个
	 */
	//string转换为bool
	str1:="true"
	b1,err := strconv.ParseBool(str1)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%t\n",b1,b1)
	//将bool类型转换回string
	b2:=strconv.FormatBool(b1)
	fmt.Printf("%T,%s\n",b2,b2)

	//整数
	str2:="100"
	i2,err := strconv.ParseInt(str2,10,64)  //base 十进制 , bitSize 最大位数
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n",i2,i2)  //转换成了int64类型
	//int转回string
	b3:=strconv.FormatInt(i2,10)
	fmt.Println(b3)

	//但最常用的string-->int间转换还是下面的 Atoi和Itoa
	//Atoi :string -> int
	ii,err:=strconv.Atoi(str2)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n",ii,ii)
	//Itoa :int -> string
	ss:= strconv.Itoa(ii)
	fmt.Printf("%T,%s\n",ss,ss)

}
