package arms

import (
	"encoding/json"
)

func JsonEncode(obj any) string {
	marshal, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func JsonEncodeFormat(obj any) string {
	marshal, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		return ""
	}
	return string(marshal)
}

func JsonDecode[T any, P string | []byte](str P) T {
	var obj T
	_ = json.Unmarshal([]byte(str), &obj)
	return obj
}
