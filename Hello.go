package main // 包声明

import "fmt" // 引入 fmt 格式化IO 的包

func main() { //定义main函数，每个程序的入口；另外 { 不能单独一行
	/*
		这里是我的多行注释内容
		也称为块注释
	*/
	fmt.Println("Let's Go!") // ①Println 函数名的首字母必须大写；②每一行命令最后的分号不许手动写，Go 编辑器会自动加上
}
