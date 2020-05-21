package main // 包声明

import (
	"fmt"
) // 引入 fmt 格式化IO 的包
//go语言的数据类型
func main() {

	// 按类别分：布尔型、数字类型、字符串类型、派生类型:
	//其中派生类型包括：
	/*
		(a) 指针类型（Pointer）
		(b) 数组类型
		(c) 结构化类型(struct)
		(d) Channel 类型
		(e) 函数类型
		(f) 切片类型
		(g) 接口类型（interface）
		(h) Map 类型
	*/

	//数字类型按照架构又可以细分为：
	/*
		uint8
		无符号 8 位整型 (0 到 255)
		 uint16
		无符号 16 位整型 (0 到 65535)
		 uint32
		无符号 32 位整型 (0 到 4294967295)
		 uint64
		无符号 64 位整型 (0 到 18446744073709551615)
		int8
		有符号 8 位整型 (-128 到 127)
		 int16
		有符号 16 位整型 (-32768 到 32767)
		int32
		有符号 32 位整型 (-2147483648 到 2147483647)
		 int64
		有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)
	*/

	//浮点型：
	/*
		loat32
		IEEE-754 32位浮点型数
		float64
		IEEE-754 64位浮点型数
		complex64
		32 位实数和虚数
		complex128
		64 位实数和虚数
	*/

	//其它数值类型
	/*
		byte
		类似 uint8
		rune
		类似 int32
		uint
		32 或 64 位
		int
		与 uint 一样大小
		uintptr
		无符号整型，用于存放一个指针
	*/

	//变量的申明格式 var 变量名 类型 ，之间必须使用空格隔开
	var bl bool         //bool 类型 默认为false
	var age int         //int 类型默认为 0
	var score32 float32 //默认为 0
	var score64 float64 //默认为 0

	var name string // 默认空字符串

	fmt.Println(bl)
	fmt.Println(age)
	fmt.Println(score32)
	fmt.Println(score64)
	fmt.Println(name + "---")
}
