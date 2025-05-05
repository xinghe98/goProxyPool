package common

// IPGeter 是获取模块的接口
type IPGeter interface {
	GetIps() []string
}

// IPSaver 是储存模块的接口
// 包含增、删、查
type IPStorger interface {
	/**
	   *
	   * @param value 代理ip{string}
	   * @param score 分数(默认为10){number}
	   * @returns boolean
	   * 储存方法
	  首次储存的时候分数为20
	  首次测试成功增加分数至100,或测试出现问题的扣分的方法
	*/
	SaveIp(value string, score float64) bool

	// api获取代理
	GetIp() (string, error)

	GetSomeIp(start int64, end int64) []string

	GetCount() int
}

// IPDetecter 是测试模块的接口
type IPDetecter interface {
	TestIp(ip []string)
}

// IPApier 是对外调用模块的接口
type IPApier interface {
	Run()
}
