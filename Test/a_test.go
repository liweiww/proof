package main

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestName(t *testing.T) {
	filePath := "example.txt" // 文件路径
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}

	// 将文件内容输出到控制台
	fmt.Println("文件内容：", string(fileContent))
}
