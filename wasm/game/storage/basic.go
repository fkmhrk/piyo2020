package storage

import (
	"encoding/base64"
	"encoding/json"
	"syscall/js"

	"github.com/fkmhrk/go-wasm-stg/game"
)

type basicStorage struct{}

func NewBasicStorage() game.Storage {
	return &basicStorage{}
}

func (s *basicStorage) Save(key string, data map[string]interface{}) error {
	bytesData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	str := base64.StdEncoding.EncodeToString(bytesData)
	localStorage := js.Global().Get("localStorage")
	localStorage.Call("setItem", key, str)
	return nil
}

func (s *basicStorage) Load(key string) (map[string]interface{}, error) {
	localStorage := js.Global().Get("localStorage")
	str := localStorage.Call("getItem", key)
	if str.IsNull() {
		return nil, nil
	}
	bytesData, err := base64.StdEncoding.DecodeString(str.String())
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	err = json.Unmarshal(bytesData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
