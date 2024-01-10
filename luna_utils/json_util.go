package luna_utils

import (
	"encoding/json"
	"fmt"
)

func ToJSON(obj interface{}) ([]byte, error) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return bytes, err
	}
	return bytes, nil
}


func FormatJSONAsString(resp map[string]interface{}) string{
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("FormattedJSON Error:", err)
		return ""
	}
	return string(jsonData)
}

func FormatJSONAsBytes(resp map[string]interface{}) []byte{
	jsonData, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("FormattedJSON Error:", err)
		return nil
	}
	return jsonData
}