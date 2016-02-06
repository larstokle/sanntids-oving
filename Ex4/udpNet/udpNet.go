package udpNet

import (
	"fmt"
	"net"
)

func StartSender(addr string, msg chan string) {

	toAddr, err := net.ResolveUDPAddr("udp", addr)
	CheckAndPrintError(err, "ResolveUDP error")

	conn, err := net.DialUDP("udp", nil, toAddr)
	//defer conn.Close()
	CheckAndPrintError(err, "DialUDP error")

	go func() {
		for {
			strToSend := <-msg
			_, err := conn.Write([]byte(strToSend))
			CheckAndPrintError(err, "Writing error")
		}
	}()

}

func StartReciever(port string, msg chan string) {

}

func CheckAndPrintError(err error, info string) {
	if err != nil {
		fmt.Println(info, ": ", err)
		//exit(1) maybe??
	}
}
