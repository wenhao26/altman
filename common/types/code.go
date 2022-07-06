package types

type Codes struct {
	ERR     int
	OK      int
	Message map[int]string
}

var RespCode = &Codes{
	ERR: -1,
	OK:  0,
}

func init() {
	RespCode.Message = map[int]string{
		RespCode.ERR: "请求失败",
		RespCode.OK:  "请求成功",
	}
}

func (c *Codes) GetMsg(code int) string {
	message, ok := c.Message[code]
	if !ok {
		return ""
	}
	return message
}
