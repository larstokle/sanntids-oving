// Go 1.2
// go run helloworld_go.go

package main

import (
    "fmt"
    "runtime"
    //"time"
)

var i = 0


func plus(ch chan int, quit chan int) {
    //var k int
   for j :=0; j < 1000000 -1; j++{
        k := <- ch
        k++
        i++
        ch <- k
        //fmt.Println("P: ", i)
   }
   quit <- 1
}

func minus(ch chan int, quit chan int) {
   for j :=0; j < 1000000; j++{
        k := <- ch
        k--
        i--
        ch <- k
        //fmt.Println("M: ", i)
   }
   quit <- -1
}

func main() {
    ch := make(chan int, 1)
    quit := make(chan int, 2)

    runtime.GOMAXPROCS(runtime.NumCPU())

    go plus(ch, quit)
    go minus(ch, quit)

    ch <-0
    
    <- quit
    <- quit

    fmt.Println("Di: ", i, " Dk: ", <-ch)
}