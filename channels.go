package main

import "fmt"

func main() {
	//使用make（chan-val类型）创建新频道。通道由它们所传达的值键入
	messages := make(chan string)
	//向消息通道发送“ping”
	go func() { messages <- "ping" }()
	//从消息通道中接收值
	msg := <-messages
	fmt.Println(msg)
}
