package main

import (
	"belajar-rpc/model"
	"fmt"
	"log"
	"net/rpc"
)

func main() {

	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	defer func(client *rpc.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	
	args := model.Input{Pertama: 3, Kedua: 4}

	// try one of this method:
	syncWay(client, args)
	//asyncWay(client, args)

}

const KalkulatorJumlah = "Kalkulator.Jumlah"

func syncWay(client *rpc.Client, args model.Input) {

	var reply model.Output

	if err := client.Call(KalkulatorJumlah, args, &reply); err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("Arith: %d+%d=%d\n", args.Pertama, args.Kedua, reply.Hasil)
}

func asyncWay(client *rpc.Client, args model.Input) {

	var reply model.Output

	x := client.Go(KalkulatorJumlah, args, &reply, nil)

	a := <-x.Done

	if err := a.Error; err != nil {
		log.Fatal("arith error:", err)
	}

	fmt.Printf("Arith: %d+%d=%d\n", args.Pertama, args.Kedua, reply.Hasil)
}
