package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	time0 := time.Now()
	output, err := exec.Command("/bin/ls").CombinedOutput()
	// output, err := exec.Command("/bin/ls", "--version").CombinedOutput()
	time1 := time.Now()
	fmt.Println(time0)
	fmt.Println(time1)
	fmt.Println(output)
	fmt.Println(err)
}
