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

	kalkulator := new(model.Kalkulator)

	err := rpc.Register(kalkulator)
	if err != nil {
		log.Fatal(err.Error())
	}

	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	fmt.Println("server is running and waiting for request from client...")

	err = http.Serve(l, nil)
	if err != nil {
		return
	}

}
