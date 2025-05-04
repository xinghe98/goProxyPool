package detectMoudle

type detectIP struct{}

// 初始化校验器
func NewDetect() detectIP {
	return detectIP{}
}

// 测试ip可用性
func (d *detectIP) TestIp(ip string) {
}
