package main

import (
	"./udpNet"
	"strconv"
	"time"
)

func main() {
	sendMsg := make(chan string)
	recieveMsg := make(chan string)
	broadCastAddr := "129.241.187.255:30004"
	listenPort := ":30004"

	udpNet.MakeSender(broadCastAddr, sendMsg)
	udpNet.MakeReciever(listenPort, recieveMsg)

	go func() {
		for {
			println(<-recieveMsg)
		}
	}()

	i := 0
	for {
		sendMsg <- "testing: " + strconv.Itoa(i) + "\000"
		time.Sleep(time.Second * 2)
		i++
	}
}
