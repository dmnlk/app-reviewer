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
		tick := time.Tick(3 * time.Second)
		for {

			select {
			case _ = <-tick:
				now := time.Now()
				content := []byte(now.String())
				ioutil.WriteFile("/tmp/output", content, os.ModePerm)
			}
		}
	}
}

func GetAppStoreReview() {

}

func GetPlayStoreReview() {

}

func Notify() {

}
