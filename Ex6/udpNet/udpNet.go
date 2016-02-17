package udpNet

import (
	"fmt"
	"net"
	"time"
)

func MakeSender(addr string, msg chan string, quit chan bool) {

	toAddr, err := net.ResolveUDPAddr("udp", addr)
	CheckAndPrintError(err, "ResolveUDP error")

	conn, err := net.DialUDP("udp", nil, toAddr)
	//defer conn.Close()
	CheckAndPrintError(err, "DialUDP error")

	go func() {
		defer conn.Close()
		for {
			select {
			case q := <-quit:
				if q {
					defer func() { quit <- false }()
					defer fmt.Println("Quiting Sender")
					return
				}
			case strToSend := <-msg:
				_, err := conn.Write([]byte(strToSend))
				CheckAndPrintError(err, "Writing error")
			}
		}
	}()
}

func MakeReciever(port string, msg chan string, quit chan bool) {

	localAddr, err := net.ResolveUDPAddr("udp", port)

	CheckAndPrintError(err, "Resolve UDP error")

	conn, err := net.ListenUDP("udp", localAddr)

	CheckAndPrintError(err, "ListenUDP error")

	buf := make([]byte, 1024)
	go func() {
		defer conn.Close()

		for {
			select {
			case q := <-quit:
				if q {
					defer func() { quit <- false }()
					defer fmt.Println("Quiting Reciever")
					return
				}
			default:
				conn.SetReadDeadline(time.Now().Add(time.Millisecond * 2000))
				n, _, err := conn.ReadFromUDP(buf)
				CheckAndPrintError(err, "ReadFromUDP error")
				if err == nil {
					msg <- string(buf[0:n])
				}

			}
		}
	}()
}

func CheckAndPrintError(err error, info string) {
	if err != nil && !err.(net.Error).Timeout() {
		fmt.Println(info, ": ", err)
		//exit(1) maybe??
	}
}
