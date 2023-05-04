package p2p

import (
	"fmt"
	"net"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "p2p:udpServer",
		Short: "p2p udpServer",
		Run:   runP2PUdpServer,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

// 接收消息，验证格式，存储列表，返回列表
func runP2PUdpServer(_ *cobra.Command, argv []string) {
	localPort := 9555
	if len(argv) >= 1 {
		localPort = cast.ToInt(argv[0])
		if localPort == 0 {
			localPort = 9555
		}
	}
	fmt.Println(localPort)
	//服务器启动侦听，定义端口
	listener, _ := net.ListenUDP("udp", &net.UDPAddr{Port: localPort})
	defer listener.Close()
	//定义切片存放udp地址
	peers := make([]*net.UDPAddr, 2, 2)
	buf := make([]byte, 256) //从两个udp消息中获得连接的地址A和B
	n, addr, _ := listener.ReadFromUDP(buf)

	fmt.Printf("read from <%s>:%s\n", addr.String(), buf[:n])
	peers[0] = addr
	n, addr, _ = listener.ReadFromUDP(buf)
	fmt.Printf("read from <%s>:%s\n", addr.String(), buf[:n])
	peers[1] = addr
	fmt.Println("begin nat")
	//将A和B分别介绍
	listener.WriteToUDP([]byte(peers[0].String()), peers[1])
	listener.WriteToUDP([]byte(peers[1].String()), peers[0])
	//睡眠，确保发送成功
	time.Sleep(time.Second * 10)
}
