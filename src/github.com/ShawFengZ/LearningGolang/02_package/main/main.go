package main

import (
	"fmt"
	"github.com/ShawFengZ/LearningGolang/02_package/stringUtil"
)

/**
go build 生成.exe文件到bin文件夹下
go install 生成.a文件到pkg文件夹下
*/
func main() {
	fmt.Println(stringUtil.Reverse("!oG ,olleH"))
	fmt.Println(stringUtil.MyName)
}
