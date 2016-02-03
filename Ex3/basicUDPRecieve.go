package main

import (
	"fmt"
	"net"
)

type Udp_message struct {
	Addr   string //if receiving raddr=senders address, if sending raddr should be set to "broadcast" or an ip:port
	Data   string //TODO: implement another encoding, strings are meh
	Length int    //length of received data, in #bytes // N/A for sending
}

func main() {
	//set up local recieve port
	addr, err := net.ResolveUDPAddr("udp", ":20004")

	if err != nil {
		fmt.Println(err)
	}

	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	fmt.Println("conn now got localAddr: " + conn.LocalAddr().String())
	//fmt.Println("conn now got remoteAddr: " + conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	for {
		n, newAddr, err := conn.ReadFromUDP(buf)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("recieved: ", string(buf[0:n]), "from ", newAddr)

	}

}
