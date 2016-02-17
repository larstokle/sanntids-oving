package main

import (
	"./udpNet"
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	master := false

	send := make(chan string)
	qSend := make(chan bool, 1)
	recieve := make(chan string)
	qRecieve := make(chan bool, 1)

	slavePort := ":30004"
	masterPort := ":31004"
	udpNet.MakeSender(masterPort, send, qSend)
	udpNet.MakeReciever(slavePort, recieve, qRecieve)

	fmt.Println("slaving away")
	lastRecieveTime := time.Now()
	backupNum := 0
	state := 2
	for !master {
		select {
		case msg := <-recieve:
			if msg == "starting" && state == 2 {
				state = 0
			} else if msg == "done" && state == 0 {
				state = 1
			} else if state == 1 {
				backupNum, _ = strconv.Atoi(msg)
				lastRecieveTime = time.Now()
				state = 2
			}
		default:
			if time.Since(lastRecieveTime).Seconds() > 3.0 {
				fmt.Println("becomming master")
				master = true
			}
		}
	}

	qSend <- true
	qRecieve <- true
	fmt.Println("Quiting network")
	time.Sleep(time.Second * 5)

	udpNet.MakeSender(slavePort, send, qSend)
	udpNet.MakeReciever(masterPort, recieve, qRecieve)
	time.Sleep(time.Second * 1)

	cmd := exec.Command("gnome-terminal", "-x", "go", "run", "safeCounter.go")
	cmd.Start()
	time.Sleep(time.Second * 1)

	num := backupNum
	for {
		if state == 2 {
			send <- "starting"
			//check slave...
			state = 0
		} else if state == 0 {
			fmt.Println(num)
			send <- "done"
			//check slave...
			state = 1
		} else if state == 1 {
			num++
			send <- strconv.Itoa(num)
			//check slave
			state = 2
			time.Sleep(time.Second * 1)
		}
	}
}
