package main 

import (
	"net"
	"log"
	"fmt"
	"bufio"
	"github.com/golang/protobuf/proto"
	"github.com/tianyaqu/simplesvr/pbcodec"
	m "github.com/tianyaqu/simplesvr/proto"
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

	p := pbcodec.PbCodec{}	
	err = p.Decode(buffer[:len])
	if err != nil {
		fmt.Println(err)
		return
	}


	fmt.Println(p)

	req := &m.Request{}
	err = proto.Unmarshal(([]byte)(p.Info), req)
	if err != nil {
		log.Fatalf("Unmarshal error %v\n",err)
		return
	}

	answer := *req.Query + " success"
	//fmt.Println("answ" + answer)

	rsp := &m.Response{}
	rsp.Answer = proto.String(answer)

	data, err := proto.Marshal(rsp)
	if err != nil {
		log.Fatalf("marshalling error: %v\n", err)
		return
	}

	p.Info = string(data)	
	rspbytes := p.Encode()
	conn.Write(rspbytes)
}
