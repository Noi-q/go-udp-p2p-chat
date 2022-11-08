package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// 客户端
func main() {
	// 建立客户端
	if len(os.Args) < 3 {
		fmt.Println("./client port name")
		return
	}
	port, _ := strconv.Atoi(os.Args[1])
	name := os.Args[2]

	localaddr := net.UDPAddr{
		Port: port,
	}
	remoteAddr := net.UDPAddr{
		Port: 9527,
		IP:   net.ParseIP("192.168.10.13"),
	}
	udp, err := net.DialUDP("udp", &localaddr, &remoteAddr)
	if err != nil {
		log.Panic("Failed to DialUDP", err)
	}
	// 发送消息
	udp.Write([]byte("I am a peer: " + name))
	// 接收服务器的消息
	buf := make([]byte, 256)
	n, _, err := udp.ReadFromUDP(buf) // buf是对象的地址
	if err != nil {
		log.Panic("Failed to ReadFromUDP", err)
	}
	// 和对象建立沟通,抛弃服务器
	toAddr, err := parseIP(string(buf[:n]))
	if err != nil {
		log.Panic("Failed to parseIP", err)
	}
	fmt.Println("对象地址:", toAddr)
	udp.Close()
	// NAT 打通
	p2pchat(&localaddr, &toAddr)
}

func parseIP(addr string) (net.UDPAddr, error) {
	split := strings.Split(addr, ":")
	if len(split) != 2 {
		fmt.Println("IP addr is not valid")
		return net.UDPAddr{}, errors.New("ipaddr err")
	}
	ip := split[0]
	atoi, _ := strconv.Atoi(split[1])
	return net.UDPAddr{
		IP:   net.ParseIP(ip),
		Port: atoi,
	}, nil
}

// p2p通信
func p2pchat(fromAddr, toAddr *net.UDPAddr) {
	// 请求Dial对象
	udp, err := net.DialUDP("udp", fromAddr, toAddr)
	if err != nil {
		log.Panic("Failed to DialUDP", err)
	}
	// 发消息
	write, err := udp.Write([]byte("有人连接我了"))
	fmt.Println(write, err)
	// 启动一个协程处理网络消息 接收消息
	go func() {
		buf := make([]byte, 256)
		for {
			n, _, err2 := udp.ReadFromUDP(buf)
			if err2 != nil {
				fmt.Println("ReadFromUDP err: ", err2)
				continue
			}
			fmt.Printf("我收到了消息: %sp2p>", buf[:n])
		}
	}()
	// 处理标准输入,写给UDP对象
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("p2p>")
		readString, err := reader.ReadString('\n')
		if err != nil {
			log.Panic("read err: ", err)
		}
		udp.Write([]byte(readString))
	}
}
