package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type StoreJSON struct {
	Path string
}

func (s *StoreJSON) Save(name string, obj interface{}) error {
	json, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	file := fmt.Sprintf("%s/%s.json", s.Path, name)
	err = ioutil.WriteFile(file, json, 0644)
	return err
}
