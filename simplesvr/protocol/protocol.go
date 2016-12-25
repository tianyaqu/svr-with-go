package protocol

import (
	"fmt"
	"bytes"
	"encoding/binary"
)

type Packet struct {
	startx uint8
	len uint32
	Cmd uint8
	seq uint32
	Info   string
	endx uint8
}

func NewPacket(cmd uint8,info string) (p *Packet) {
	p = &Packet{}
	p.Cmd = cmd
	p.Info = info
	return p
}


func(p Packet) String() string {
	return "start:" + string(rune(p.startx)) + " cmd: " + fmt.Sprintf("%v",p.Cmd) + " seq: " + fmt.Sprintf("%v",p.seq) + " info: " + p.Info + " end: " + string(rune(p.endx)) 
}

func(p *Packet) Encode() []byte {
	p.startx = '('
	p.endx = ')'
	p.len =  uint32(len([]byte(p.Info))) + uint32(11) 

	infoBuffer := new(bytes.Buffer)
	binary.Write(infoBuffer,binary.LittleEndian,[]byte(p.Info))

	buffer := new(bytes.Buffer)
	binary.Write(buffer,binary.LittleEndian,p.startx)
	binary.Write(buffer,binary.LittleEndian,p.len)
	binary.Write(buffer,binary.LittleEndian,p.Cmd)
	binary.Write(buffer,binary.LittleEndian,p.seq)
	binary.Write(buffer,binary.LittleEndian,infoBuffer.Bytes())
	binary.Write(buffer,binary.LittleEndian,p.endx)
	
	return buffer.Bytes()
}

func(p *Packet) Decode(buffer []byte) {
	bufferReader := bytes.NewBuffer(buffer)

	binary.Read(bufferReader,binary.LittleEndian,&(p.startx))
	binary.Read(bufferReader,binary.LittleEndian,&(p.len))
	binary.Read(bufferReader,binary.LittleEndian,&(p.Cmd))
	binary.Read(bufferReader,binary.LittleEndian,&(p.seq))
	p.Info = (string)(bufferReader.Next(int(p.len - uint32(11))))
	binary.Read(bufferReader,binary.LittleEndian,&(p.endx))
}

/*
func main() {
	var packet Packet
	var x Packet
	packet.Cmd = 20
	packet.seq = 5
	packet.Info = "hello"

	mash := packet.Encode()
	x.Decode(mash)

}
*/
