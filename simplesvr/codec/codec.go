package codec


type Packer interface {
    Encode() []byte
    Decode(buffer []byte) error
    Check() error
}