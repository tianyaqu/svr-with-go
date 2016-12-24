package main 

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"github.com/tianyaqu/protocol"
)

func main(){
	svr,err := net.Listen("tcp",":9876")
	if err != nil {
		log.Fatalln("listen faied,exit")
		return
	}

	for{
		conn,err := svr.Accept()
		if err != nil {
			log.Fatalln("accept error,continue")
			continue
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte,1500)
	b := bufio.NewReader(conn)
	len,err := b.Read(buffer)
	if err != nil {
		fmt.Println("NIL")	
		return
	}

	p := protocol.Packet{}
	p.Decode(buffer[:len])
	fmt.Println(p)
	p.Info = "ok"
	rsp := p.Encode()
	conn.Write(rsp)
}