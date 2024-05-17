package utils

import "encoding/json"

func ToJSONString(value interface{}) string {
	out, _ := json.Marshal(value)
	return string(out)
}
