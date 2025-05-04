package fetchMoudle

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type fetchIp struct {
	Url string
}

func NewFetch(url string) fetchIp {
	return fetchIp{
		Url: url,
	}
}

func (f *fetchIp) GetIps() []string {
	resp, err := http.Get(f.Url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return strings.Split(string(body), "\n")
}
