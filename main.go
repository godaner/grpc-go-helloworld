package main

import (
	"grpc-go-stream/server"
	"time"
	"grpc-go-stream/client"
)

func main(){
	go func() {
		server.Run()
	}()
	time.Sleep(3*time.Millisecond)
	client.Run()
}
