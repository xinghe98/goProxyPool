package common

// IPGeter 是获取模块的接口
type IPGeter interface {
	GetIps() []string
}

// IPSaver 是储存模块的接口
// 包含增、删、查
type IPSaver interface {
	SaveIp(value string, score int) bool
	DeleteIp(ip string)
	GetIp() string
}

// IPDetecter 是测试模块的接口
type IPDetecter interface {
	TestIp(ip string)
}
