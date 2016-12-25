package main 

import (
	"net"
	"testing"
	"github.com/tianyaqu/simplesvr/pbcodec"
	"github.com/golang/protobuf/proto"
	m "github.com/tianyaqu/simplesvr/proto"


)

func TestSvr(t *testing.T){
	conn,err := net.Dial("tcp",":9876")
	defer conn.Close()

	if err != nil {
		t.Error("dia failed,exit")
		return
	}


	req := &m.Request{}
	req.Query = proto.String("hello")
	data, err := proto.Marshal(req)
	if err != nil {
		t.Error("marshalling error")
	}
	

	p := pbcodec.New(16,string(data))
	mash := p.Encode()
	
	conn.Write(mash)
	
	buffer := make([]byte,1500)
	len,_ := conn.Read(buffer)

	v := pbcodec.PbCodec{}
	err = v.Decode(buffer[:len])
	if err != nil {
		t.Error("decode failed")
	}


	rsp := &m.Response{}
	err = proto.Unmarshal(([]byte)(v.Info), rsp)

	if(*rsp.Answer == "hello success"){
		t.Log("pass")
	} else {
		t.Error("not pass")
	}
}
