package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/ls", "-l")
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(string(buf))
}
