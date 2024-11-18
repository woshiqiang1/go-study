package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./test_main.go", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	code := make([]byte, 1024) // 注意：切片长度决定了读取内容的长度
	n, err := file.Read(code)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d characters were successfully read\n", n)
	fmt.Printf("%s\n", code)
}
