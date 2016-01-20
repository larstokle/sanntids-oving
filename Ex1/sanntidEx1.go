// Go 1.2
// go run helloworld_go.go

package main

import (
    "fmt"
    "runtime"
    "time"
)

var i = 0

func plus() {
   for j :=0; j < 1000000; j++{
        i++
        //fmt.Println("P: ", i)
   }
}

func minus() {
   for j :=0; j < 1000000; j++{
        i--
        //fmt.Println("M: ", i)
   }
}

func main() {
    //runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
    runtime.GOMAXPROCS(2)                                       // Try doing the exercise both with and without it!
    go plus()
    go minus()
                         // This spawns someGoroutine() as a goroutine

    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(2000*time.Millisecond)
    fmt.Println("D: ", i)
}