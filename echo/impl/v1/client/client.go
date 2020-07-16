package client

import (
	"echo/impl/common"
	"fmt"
	"log"
	"net/rpc"
)

func Run() error{
	log.Println("[.] Dialing")
	client, err := rpc.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln("[x] Dialing:", err)
		return err
	}
	log.Println("[o] Dialing")

	log.Println("[.] Calling")
	args := common.Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", &args, &reply)
	if err != nil {
		log.Fatalln("[x] Calling err:", err)
		return err
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
	log.Println("[o] Calling")
	return nil;
}


