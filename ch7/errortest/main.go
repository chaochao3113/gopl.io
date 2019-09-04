package main

import (
	"errors"
	"fmt"
	"syscall"
)

//创建error的二种方式
func main() {
	//第一种,直接使用errors.New(string) error
	err1 :=  errors.New("test")
	fmt.Println(err1)

	//第二种,使用fmt.Errorf(string) error
	fmt.Println(fmt.Errorf("test1"))

	//syscall---Go语言底层系统调用API
	err := syscall.Errno(2)
	fmt.Println(err.Error())
	fmt.Println(err)
}
