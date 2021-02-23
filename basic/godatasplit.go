package basic

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

//golang 官网的bufio库中提供了Scanner 用于数据的分割

/*
1 创建scanner
2 自定义分割方式
3 注入自定义分割方式
4 调用分割
*/

func DataSplitTest() {
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))

	defineSplit := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		//token 是分割后的值，调用scanner.Text() 获取
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}

	//注入回调函数
	scanner.Split(defineSplit)

	for scanner.Scan() {
		fmt.Printf("%s \n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invaild input: %s", err)
	}
}
