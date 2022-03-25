package main

import (
	"fmt"
)

func checkResponse(responseRaw map[string]interface{}) (string, bool) {
	_code, ok := responseRaw["cod"]
	if ok {
		code := int(_code.(float64))

		if code == 200 {
			return "", true
		} else if code == 401 {
			return fmt.Sprintf("Error: code %v. Your API key is invalid.", code), false
		} else {
			return fmt.Sprintf("Error: code %v", code), false
		}
	} else {
		return "", true
	}	
}