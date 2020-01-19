// +build ignore   //该命令的目的是同目录下的其它文件中存在的main方法，因为Go 像C语言一样，一个工程目录下，只能有一个main函数入口

package main

import "fmt"

func main() {
	var name string // 变量声明统一用var 定义，格式为： var varName type
	var age int
	name = "Golang"
	age = 9
	fmt.Printf("I'm %s,i am %d years old\n", name, age) //%s 为字符格式化输出，%d 为数字格式化输出

	//Go中常见的数据类型如下
	var str string = "This is one test!"
	var int_a int
	var bool_flag bool
	var byte_flag byte
	var float_a float64 = 7.6665

	fmt.Printf("%s,%v,%v,%f\n", str, int_a, bool_flag, byte_flag, float_a)
}
