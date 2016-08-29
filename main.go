package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	if len(os.Args) > 1 {
		command := os.Args[1:]
		cmd := exec.Command(command[0], command[1:]...)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(out)
		}

		fmt.Print(string(bytes.TrimRight(out, "\n")))
	}
}
