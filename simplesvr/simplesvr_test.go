package main 

import (
	"net"
	"testing"
	"github.com/tianyaqu/simplesvr/pbcodec"
)

func TestSvr(t *testing.T){
	conn,err := net.Dial("tcp",":9876")
	defer conn.Close()

	if err != nil {
		t.Error("dia failed,exit")
		return
	}

	p := pbcodec.New(16,"hello")
	mash := p.Encode()
	
	conn.Write(mash)
	
	buffer := make([]byte,1500)
	len,_ := conn.Read(buffer)

	v := pbcodec.PbCodec{}
	v.Decode(buffer[:len])
	if v.Info != "ok" {
		t.Error("not pass")
	} else {
		t.Log("pass")
	}
}
