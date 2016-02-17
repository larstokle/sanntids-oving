package main

import (
	"fmt"
	//"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("gnome-terminal", "-x", "go", "run", "helloWorld.go") //setter opp shell command. takler ikke space
	cmd.Run()                                                                 //run venter til proc er ferdig
	//cmd.Start() //cmd.Wait()	// samme som run men venter ikke etter start
	//out, err := cmd.Output()	// ca samme som run men skal visstnok returnere noe. har ikke fått til.
	fmt.Println("Døøn")
}
