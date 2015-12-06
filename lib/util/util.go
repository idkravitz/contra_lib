package util

import "encoding/json"

func Qjson(v interface{}) string {
	js, _ := json.Marshal(v)
	return string(js)
}