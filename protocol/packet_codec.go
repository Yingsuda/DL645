package protocol

import (
	"dl645/protocol/model"
	"sync"
)

type PacketCodec interface {
	Decode([]byte) (*model.PacketData, error)
}

type DL645PacketCodec struct{}

var dl808PacketCodec *DL645PacketCodec
var codecOnce sync.Once

func NewJT808PacketCodec() *DL645PacketCodec {
	codecOnce.Do(func() {
		dl808PacketCodec = &DL645PacketCodec{}
	})
	return dl808PacketCodec
}

// 反转义 -> 校验 -> 反序列化
func (c *DL645PacketCodec) Decode(data []byte) (*model.PacketData, error) {
	return nil, nil
}
