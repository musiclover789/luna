package luna_request

import (
	"encoding/json"
	"strings"
)

type HttpMap struct {
	values []interface{}
	keys   []interface{}
}


func (receiver *HttpMap) Push(key, value interface{}) {
	if  strings.EqualFold(value.(string),""){
		return
	}
	bool := true
	for k, v := range receiver.keys {
		dataStr1, _ := json.Marshal(v)
		dataStr2, _ := json.Marshal(key)
		if string(dataStr1) == string(dataStr2) {
			receiver.values[k] = value
			bool=false
		}
	}
	if bool {
		receiver.keys = append(receiver.keys, key)
		receiver.values = append(receiver.values, value)
	}
}

func (receiver *HttpMap) Remove(key interface{}) {
	for index, v := range receiver.keys {
		dataStr1, _ := json.Marshal(v)
		dataStr2, _ := json.Marshal(key)
		if string(dataStr1) == string(dataStr2) {
			receiver.keys = append(receiver.keys[:index], receiver.keys[index+1:]...)
			receiver.values = append(receiver.values[:index], receiver.values[index+1:]...)
		}
	}

}

func (receiver *HttpMap) Get(key interface{}) interface{} {
	for k, v := range receiver.keys {
		dataStr1, _ := json.Marshal(v)
		dataStr2, _ := json.Marshal(key)
		if string(dataStr1) == string(dataStr2) {
			return receiver.values[k]
		}
	}
	return nil
}

func (receiver *HttpMap) Size() int {

	return len(receiver.keys)
}

func (receiver *HttpMap) List() ([]interface{},[]interface{}) {

	return receiver.keys ,receiver.values
}


