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
		driver.RunFloorUp()
		time.Sleep(time.Second * 5)
		driver.RunFloorDown()
		time.Sleep(time.Second * 5)
	}
}
