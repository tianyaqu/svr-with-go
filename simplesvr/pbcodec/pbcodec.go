package pbcodec

import (
    "fmt"
    "bytes"
    "encoding/binary"
    "errors"
)


type PbCodec struct {
    start uint8
    len uint32
    Cmd uint8
    Seq uint32
    Info   string
    end uint8
}

func New(cmd uint8,info string) (p *PbCodec) {
    p = &PbCodec{}
    p.Cmd = cmd
    p.Info = info
    return p
}


func(p PbCodec) String() string {
    return "start:" + string(rune(p.start)) + " cmd: " + fmt.Sprintf("%v",p.Cmd) + " seq: " + fmt.Sprintf("%v",p.Seq) + " info: " + p.Info + " end: " + string(rune(p.end)) 
}

func(p *PbCodec) Encode() []byte {
    p.start = '('
    p.end = ')'
    p.len =  uint32(len([]byte(p.Info))) + uint32(11) 

    infoBuffer := new(bytes.Buffer)
    binary.Write(infoBuffer,binary.LittleEndian,[]byte(p.Info))

    buffer := new(bytes.Buffer)
    binary.Write(buffer,binary.LittleEndian,p.start)
    binary.Write(buffer,binary.LittleEndian,p.len)
    binary.Write(buffer,binary.LittleEndian,p.Cmd)
    binary.Write(buffer,binary.LittleEndian,p.Seq)
    binary.Write(buffer,binary.LittleEndian,infoBuffer.Bytes())
    binary.Write(buffer,binary.LittleEndian,p.end)
    
    return buffer.Bytes()
}

func(p *PbCodec) Decode(buffer []byte) error {
    bufferReader := bytes.NewBuffer(buffer)

    binary.Read(bufferReader,binary.LittleEndian,&(p.start))
    binary.Read(bufferReader,binary.LittleEndian,&(p.len))
    binary.Read(bufferReader,binary.LittleEndian,&(p.Cmd))
    binary.Read(bufferReader,binary.LittleEndian,&(p.Seq))
    p.Info = (string)(bufferReader.Next(int(p.len - uint32(11))))
    binary.Read(bufferReader,binary.LittleEndian,&(p.end))

    return p.Check()
}

func(p *PbCodec) Check() error {
    if p.start == '(' && p.end == ')' {
        return nil
    } else {
        return errors.New("format incorrect")
    }
}
