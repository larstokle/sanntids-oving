package udpNet

import (
	"fmt"
	"net"
)

func MakeSender(addr string, msg chan string) {

	toAddr, err := net.ResolveUDPAddr("udp", addr)
	CheckAndPrintError(err, "ResolveUDP error")

	conn, err := net.DialUDP("udp", nil, toAddr)
	//defer conn.Close()
	CheckAndPrintError(err, "DialUDP error")

	go func() {
		defer conn.Close()
		for {
			strToSend := <-msg
			_, err := conn.Write([]byte(strToSend))
			CheckAndPrintError(err, "Writing error")
		}
	}()
}

func MakeReciever(port string, msg chan string) {

	localAddr, err := net.ResolveUDPAddr("udp", port)

	CheckAndPrintError(err, "Resolve UDP error")

	conn, err := net.ListenUDP("udp", localAddr)

	CheckAndPrintError(err, "ListenUDP error")

	buf := make([]byte, 1024)
	go func() {
		defer conn.Close()
		for {
			n, _, err := conn.ReadFromUDP(buf)
			CheckAndPrintError(err, "ReadFromUDP error")
			msg <- string(buf[0:n])
		}
	}()
}

func CheckAndPrintError(err error, info string) {
	if err != nil {
		fmt.Println(info, ": ", err)
		//exit(1) maybe??
	}
}
