package p2p

import (
	"bufio"
	"fmt"
	"github.com/spf13/cast"
	"io"
	"net"
	"os"
	"strings"
	"thh/arms"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "p_2_pclient",
		Short: "",
		Run:   runP2PClient,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

func runP2PClient(_ *cobra.Command, argv []string) {
	localPort := 5050
	if len(argv) >= 1 {
		localPort = cast.ToInt(argv[0])
	}

	fmt.Println("tcp", localPort)

	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "10.249.146.120:8282")
	localTcpAddr := &net.TCPAddr{Port: localPort}
	fmt.Println("localTcpAddr", localTcpAddr)

	conn, err := net.DialTCP("tcp", localTcpAddr, tcpAddr)

	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return
	}

	defer conn.Close()

	fmt.Println(conn.LocalAddr().String() + " : Client connected!")

	go onMessageReceived(conn)

	// im
	var remoteClient string
	for {
		for data, _ := range peerMap {
			if !strings.Contains(data, cast.ToString(localPort)) {
				// 说明目标不是本机
				remoteClient = data
			}
		}
		if remoteClient != "" {
			break
		}
		fmt.Println("没有其他客户端")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("remote client ", remoteClient)
	remoteClientInfo := strings.Split(remoteClient, ":")

	remoteClientTcpAddr := &net.TCPAddr{Port: cast.ToInt(remoteClientInfo[1]), IP: net.ParseIP(remoteClientInfo[0])}
	clientConn, err := net.DialTCP("tcp", nil, remoteClientTcpAddr)
	if err != nil {
		fmt.Println("other client connect err", err)
		return
	}
	go onlyReceived(clientConn)

	reader := bufio.NewReader(os.Stdin)
	for {
		result, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		result = append(result, []byte("\n")...)
		_, err = conn.Write(result)
		if err != nil {
			return
		}
	}
}

func onlyReceived(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + "other clietn init \n")
	_, err := conn.Write(b)
	if err != nil {
		return
	}
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println(msg)
		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
	}
}

// reg client 2 server
// get other client
// im

func onMessageReceived(conn *net.TCPConn) {

	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + " init \n")
	_, err := conn.Write(b)
	if err != nil {
		return
	}
	for {
		msg, err := reader.ReadString('\n')
		peerMap = arms.JsonDecode[map[string]string](cast.ToString(msg))
		fmt.Println("server:" + msg)
		fmt.Println(conn.LocalAddr().String())

		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second * 5)

		b := []byte("getMap\n")
		_, err = conn.Write(b)

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
