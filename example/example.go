package main

import (
	"fmt"
	"os"
	"time"

	"logging"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	logging.SetRotate("hour")

	err := logging.SetOutputByName("example.log")
	checkError(err)

	timer := time.NewTicker(time.Duration(10) * time.Second)
	for {
		select {
		case <-timer.C:
			logging.Debug(time.Now().String())
			logging.Info(time.Now().String())
			logging.Warningf("%s", time.Now().String())
			logging.Errorf("%s", time.Now().String())
		}
	}
}
