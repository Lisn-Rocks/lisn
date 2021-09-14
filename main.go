package main

import (
	_ "github.com/lisn-rocks/lisn/configs"
	"github.com/lisn-rocks/lisn/server"
	"github.com/sharpvik/log-go/v2"
)

func init() {
	log.SetLevel(log.LevelDebug)
}

func main() {
	done := make(chan bool, 1)
	go server.New().ServeWithGrace(done)
	<-done
}
