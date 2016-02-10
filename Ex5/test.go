package main

import (
	"./driver"
	"fmt"
	"time"
)

func main() {
	fmt.Println("driver init returned: ", driver.Init())
	time.Sleep(time.Second * 5)
	for {
		driver.RunTopFloor()
		time.Sleep(time.Second * 5)
		driver.RunBottomFloor()
		time.Sleep(time.Second * 5)
	}
}
