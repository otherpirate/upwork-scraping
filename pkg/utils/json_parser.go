package utils

import "encoding/json"

func ToJSON(obj interface{}) ([]byte, error) {
	json, err := json.Marshal(obj)
	return json, err
}
