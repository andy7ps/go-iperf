package main

import (
	"fmt"
	"os"
	"time"

	"github.com/andy7ps/go-iperf"
)

func main() {
	s := iperf.NewServer()
	err := s.Start()
	if err != nil {
		fmt.Println("failed to start server")
		os.Exit(-1)
	}

	for s.Running {
		time.Sleep(1 * time.Second)
	}

	fmt.Println("server has exited")
}
