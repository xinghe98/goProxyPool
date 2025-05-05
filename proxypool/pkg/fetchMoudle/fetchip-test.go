package fetchMoudle

type fetchIpTest struct {
	Url string
}

func NewFetchTest(url string) fetchIpTest {
	return fetchIpTest{
		Url: url,
	}
}

func (f *fetchIpTest) GetIps() []string {
	return []string{}
}

func (f *fetchIpTest) Run() {
}
