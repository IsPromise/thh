package p2p

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "p2p:udpClient",
		Short: "p2p udpClient",
		Run:   runP2PUdpClient,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

// 发送消息，注册自己，获取列表，列表通信
func runP2PUdpClient(_ *cobra.Command, argv []string) {
	localPort := 5050
	if len(argv) >= 1 {
		localPort = cast.ToInt(argv[0])
	}

	fmt.Println("tcp", localPort)

	//服务器IP
	remoteIP := "1.117.224.127"
	//remoteIP = "10.249.146.120"
	//服务器端口
	remotePort := 9555
	//绑定本地端口
	localAddr := net.UDPAddr{IP: net.IPv4zero, Port: localPort}
	//与服务器建立联系
	conn, err := net.DialUDP("udp", &localAddr, &net.UDPAddr{IP: net.ParseIP(remoteIP), Port: remotePort})
	fmt.Println(conn.LocalAddr())
	if err != nil {
		log.Panic("UDP拨号失败")
	}
	conn.Write([]byte("I am SEVERA")) //从服务器中获得目标地址
	buf := make([]byte, 256)
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		log.Panic("读取消息失败", err)
	}
	conn.Close()
	toAddr := parseAddr(string(buf[:n]))
	fmt.Println("目标", toAddr)
	p2p(&localAddr, &toAddr)

}

func parseAddr(addr string) net.UDPAddr {
	t := strings.Split(addr, ":")
	port, _ := strconv.Atoi(t[1])
	return net.UDPAddr{IP: net.ParseIP(t[0]), Port: port}
}

func p2p(srcAddr *net.UDPAddr, dstAddr *net.UDPAddr) {
	//请求建立联系
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println("err", err)
	}
	//发送打洞消息
	_, err = conn.Write([]byte("打洞消息"))
	if err != nil {
		fmt.Println("打洞err", err)
	}
	//启动goroutine监控标准输入
	go func() {
		buf := make([]byte, 256)
		for {
			//接受UDP消息打印
			n, add, err := conn.ReadFromUDP(buf)
			if err != nil {
				fmt.Println("接收消息", err)
			}
			fmt.Println(add)
			if n > 0 {
				fmt.Printf("收到消息:%sp2p>", buf[:n])
			}

		}
	}()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("p2p>")
		//读取标准输入，以换行为读取标志
		data, _ := reader.ReadString('\n')
		n, err := conn.Write([]byte(data))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(n)
	}
}
