package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	strEcho := "Halo \000"
	servAddr := "129.241.187.23:33546"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("Msg from server=", string(reply))

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("write to server = ", strEcho)

	_, err = conn.Read(reply)
	if err != nil {
		println("Reply from server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))
	conn.Close()
}
