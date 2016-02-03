package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type Udp_message struct {
	Raddr  string //if receiving raddr=senders address, if sending raddr should be set to "broadcast" or an ip:port
	Data   string //TODO: implement another encoding, strings are meh
	Length int    //length of received data, in #bytes // N/A for sending
}

func main() {
	//set up where to send to
	serverAddr, err := net.ResolveUDPAddr("udp", "129.241.187.143:20002")

	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	defer conn.Close()
	//laddr, err := net.ResolveUDPAddr("udp", conn.LocalAddr().String())
	fmt.Println("conn now got localAddr: " + conn.LocalAddr().String())
	fmt.Println("conn now got remoteAddr: " + conn.RemoteAddr().String())
	//fmt.Println("laddr: ", laddr)

	if err != nil {
		fmt.Println(err)
	}

	//buf := make([]byte, 1024)
	i := 0
	for {
		_, err := conn.Write([]byte("test send nr: " + strconv.Itoa(i)))
		if err != nil {
			fmt.Println(err)
		}
		i++
		time.Sleep(time.Second * 2)
	}

}
