package main

import (
	"belajar-rpc/model"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {

	err := rpc.Register(&model.Kalkulator{})
	if err != nil {
		log.Fatal(err.Error())
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}

	fmt.Println("server is running and waiting for request from client...")

	//rpc.Accept(listener)

	err = http.Serve(listener, nil)
	if err != nil {
		return
	}

}
