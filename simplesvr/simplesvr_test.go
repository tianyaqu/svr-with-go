package main 

import (
	"net"
	"testing"
	"github.com/tianyaqu/protocol"
)

func TestSvr(t *testing.T){
	conn,err := net.Dial("tcp",":9876")
	defer conn.Close()

	if err != nil {
		t.Error("dia failed,exit")
		return
	}

	p := protocol.NewPacket(16,"hello")
	mash := p.Encode()
	
	conn.Write(mash)
	
	buffer := make([]byte,1500)
	len,_ := conn.Read(buffer)

	v := protocol.Packet{}
	v.Decode(buffer[:len])
	if v.Info != "ok" {
		t.Error("not pass")
	} else {
		t.Log("pass")
	}
}
