package model

type controlCode int

const (
	D07_CTRL_RESV            controlCode = 0x00 // 保留
	D07_CTRL_SYNC_TIME       controlCode = 0x08 // 广播校时
	D07_CTRL_READ_DATA       controlCode = 0x11 // 读数据
	D07_CTRL_READ_AFTERDATA  controlCode = 0x12 // 读后续数据
	D07_CTRL_READ_ADDR       controlCode = 0x13 // 读通信地址
	D07_CTRL_WRITE_DATA      controlCode = 0x14 // 写数据
	D07_CTRL_WRITE_ADDR      controlCode = 0x15 // 写通信地址
	D07_CTRL_FREEZ_COMM      controlCode = 0x16 // 冻结命令
	D07_CTRL_MODIFY_BAUD     controlCode = 0x17 // 修改通信速率
	D07_CTRL_MODIFY_PASSWORD controlCode = 0x18 // 修改密码
	D07_CTRL_CLEAR_MAXDEMAND controlCode = 0x19 // 最大需量清零
	D07_CTRL_CLEAR_METER     controlCode = 0x1A // 电表清零
	D07_CTRL_CLEAR_EVENT     controlCode = 0x1B // 事件清零
	D07_CTRL_COMM            controlCode = 0x1C // 控制命令
)

const (
	E_D07_CTRL_SR_OK = iota // 从站正常应答
	E_D07_CTRL_SR_NO        // 从站异常应答
)

const (
	E_D07_CTRL_FLW_NONE = iota //无后续
	E_D07_CTRL_FLW_HAVE        //有后续
)
