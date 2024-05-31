package model

type CTRLCode struct {
	direct int         //帧传送方向
	reply  int         //从站应答异常标志
	follow int         // 有无后续帧
	code   controlCode // 功能码
}

type PacketData struct {
	address   string //BCD
	ctrl      CTRLCode
	length    int
	data      []byte
	checkCode int
}
