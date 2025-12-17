package main

import (
	"fmt"
	"io"
	"os"
)

// os.File实现了Reader和Writer接口

func main7() {
	readFile, err := os.Open("/Users/fordev/Workspace/StudyGo/oop/reader.go")
	if err != nil {
		fmt.Printf("Error opening read file: %s\n", err)
		return
	}
	// 延迟关闭文件（必须！避免资源泄漏）
	defer readFile.Close()

	writeFile, err := os.Create("/Users/fordev/Workspace/StudyGo/oop/test_write")
	if err != nil {
		fmt.Printf("Error opening write file: %s\n", err)
		return
	}
	defer writeFile.Close()

	buffer := make([]byte, 32)

	for {
		readCnt, err := readFile.Read(buffer)
		if err == io.EOF {
			break
		}

		if readCnt > 0 {
			writeCnt, wErr := writeFile.Write(buffer[:readCnt])
			if wErr != nil || writeCnt < readCnt {
				fmt.Printf("write file fail, %v, %v\n", wErr, writeCnt)
				break
			}
		}

		if err != nil {
			fmt.Printf("read file fail, %v\n", err)
			break
		}
	}
}
