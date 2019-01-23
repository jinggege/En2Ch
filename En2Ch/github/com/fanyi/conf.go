package fanyi

import (
	"strings"
)

const (
	Version  = "0.0.1"
	Author   = "jgg"
	Function = "English2Chinese"
	Prompt   = "\n[请输入英文单词]: "
)

func GetDescLine() string {
	return strings.Join([]string{"==APP:", Function, "  Version:", Version, "=="}, "")
}
