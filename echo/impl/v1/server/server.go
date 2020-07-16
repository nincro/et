package server

import (
	"echo/impl/common"
	"fmt"
	"net"
)



func echoHandler(conn net.Conn)  {
	defer conn.Close()

	println("[.] handling")
	buffer := make([]byte, common.SIZE_RECEIVE_BUFFER)
	len, err := conn.Read(buffer)
	if err != nil {
		println("[x] handling")
		panic(err)
	}
	msg := string(buffer[:len-1])
	println(msg)
	conn.Write([]byte(msg))

	println("[o] handling")

}

func Run(){
	addr := fmt.Sprintf("%s:%d",common.ADDR, common.PORT)
	println("[.] listening")
	listener, err := net.Listen(common.NETWORK, addr)
	if err != nil {
		println("[x] listening")
		panic(err)
	}
	println("[o] listening")

	defer listener.Close()

	for {
		println("[.] accept")
		conn, err := listener.Accept()
		if err != nil {
			println("[x] accept")
			panic(err)
		}
		fmt.Printf("remote:%s, local:%s", conn.RemoteAddr(),conn.LocalAddr())
		go echoHandler(conn)
	}
}

func init() {

}
