package main

import (
	"os"
	"os/exec"
	"fmt"
)

func main() {
	cmd := exec.Command("go","run", "helloWorld.go") //setter opp shell command. takler ikke space
	//cmd.Stdout = os.Stdout	// skulle tro denne gjør det som den under gjør, meeeen :S
    cmd.Stderr = os.Stderr		//denne linjen må være med for å få println til terminal
	cmd.Run()					//run venter til proc er ferdig
	//cmd.Start() //cmd.Wait()	// samme som run men venter ikke etter start
	//out, err := cmd.Output()	// ca samme som run men skal visstnok returnere noe. har ikke fått til.
	fmt.Println("Døøn")
}