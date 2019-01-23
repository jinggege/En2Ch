package fanyi

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Wweb struct {
	web []interface{}
}

type Basic struct {
	//us-phonetic string
	Phonetic string `json:"phonetic"`
	//uk-phonetic string
	Explains []string `json:"explains"`
}

type FyStruct struct {
	Translation []string      `json:"translation"`
	Bas         Basic         `json:"basic"`
	Query       string        `json:"query"`
	Ecode       int           `json:"errorCode"`
	Web         []interface{} `json:"web"`
}

func Format(str string) string {

	var fs FyStruct
	err := json.Unmarshal([]byte(str), &fs)

	if nil != err {
		fmt.Println("=format error: ", err.Error())
		return ""
	}

	var errorCode int = fs.Ecode

	if 0 != errorCode {
		fmt.Println("= req error error code=", errorCode)
		return ""
	}

	var queryStr string = ""

	query := fs.Query
	basic := fs.Bas
	exp := basic.Explains

	fmt.Println("\n")

	queryStr = strings.Join([]string{"----------------------start--------------------", "\n"}, "")

	queryStr = strings.Join([]string{queryStr, query, "    ", "[", basic.Phonetic, "]", "\n"}, "")

	for index, v := range exp {
		expStr := strings.Join([]string{"[", fmt.Sprint(index), "]", " ", v, "\n"}, "")
		queryStr = strings.Join([]string{queryStr, expStr}, "")
	}

	queryStr = strings.Join([]string{queryStr, "----------------------end--------------------"}, "")

	return queryStr

}
