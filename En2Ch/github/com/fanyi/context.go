package fanyi

import (
	"strings"

	"../xioxu/goreq"
)

var (
	API string = "http://fanyi.youdao.com/openapi.do?keyfrom=node-fanyi&key=110811608&type=data&doctype=json&version=1.1&q="
)

func Get(word string) string {
	url := strings.Join([]string{API, word}, "")
	url = strings.TrimSpace(url)

	req := goreq.Req(nil)
	body, _, _ := req.Get(url).Do()

	str := string(body)

	return str
}
