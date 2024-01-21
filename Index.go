package main

import (
	"fmt"
	"net"
	"time"
)

func pingServer(server string) {
	_, err := net.DialTimeout("tcp", server, time.Second*10)
	if err != nil {
		fmt.Println(server, "is down!")
		return
	}
	fmt.Println(server, "is up and running.")
}
func startBots(botCount int, server string) {
	for i := 0; i < botCount; i++ {
		go pingServer(server)
	}
}

func main() {
	var i int = 0
	for i < 10 {
		var server string = "127.0.0.1:80"
		var botCount int = 1_000

		startBots(botCount, server)

		time.Sleep(time.Second * 10)
	}
}
