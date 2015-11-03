package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		cmd := exec.Command(os.Args[0], "--child")
		cmd.Start()
	} else {
		for {
			now := time.Now()
			content := []byte(now.String())
			ioutil.WriteFile("/tmp/output", content, os.ModePerm)
		}
	}
}
