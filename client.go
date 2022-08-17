package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err = ", err)
		return
	}
	defer conn.Close()
	for {
		fmt.Println("请输入信息， 回车结束输入")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err = ", err)
		}
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}
		line = strings.TrimSpace(line)
		n, err := conn.Write([]byte(line)) //n用来统计字符个数
		if err != nil {
			fmt.Println("conn.Write err = ", err)
		}
		fmt.Printf("发送了内容:%d文字\n", n)
	}
}
