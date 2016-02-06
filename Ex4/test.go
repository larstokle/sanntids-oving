package main

import (
	"./udpNet"
	"time"
)

func main() {
	msg := make(chan string)
	broadCastAddr := "129.241.187.255:20004"

	udpNet.StartSender(broadCastAddr, msg)
	i := 0
	for {
		msg <- "testing: " + string(i) + "\000"
		time.Sleep(time.Second * 2)
		i++
	}
}
