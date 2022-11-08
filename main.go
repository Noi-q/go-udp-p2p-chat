package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// UDP服务器
func main() {
	// 建立UDP服务器
	udp, err := net.ListenUDP("udp", &net.UDPAddr{Port: 9527})
	if err != nil {
		log.Panic("Failed to ListenUDP", err)
	}
	// 接收两个消息
	peers := make([]*net.UDPAddr, 2, 2)
	buf := make([]byte, 256)
	n, addr, err := udp.ReadFromUDP(buf)
	if err != nil {
		log.Panic("1 Failed to ReadFromUDP")
	}
	defer udp.Close()
	fmt.Printf("Begin Server....")
	peers[0] = addr
	fmt.Printf("1 read %d size, from %s, msg:%s\n", n, addr.String(), buf[:n])

	// 第二个
	n, addr, err = udp.ReadFromUDP(buf)
	if err != nil {
		log.Panic("2 Failed to ReadFromUDP")
	}
	peers[1] = addr
	fmt.Printf("2 read %d size, from %s, msg:%s\n", n, addr.String(), buf[:n])

	// 消息交换
	udp.WriteToUDP([]byte(peers[0].String()), peers[1])
	udp.WriteToUDP([]byte(peers[1].String()), peers[0])
	// 退出
	fmt.Println("Server will exit after 10m")
	time.Sleep(time.Second * 10) // 延迟10秒退出
	// 退出之后不会影响客户端之间通信
}
