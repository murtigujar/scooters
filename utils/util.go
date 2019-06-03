package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func JSONMarshal(data interface{}) ([]byte) {
	e, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return e
}

func HTTPResponse(w http.ResponseWriter, statusCode int, resp []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func GetParam(param string, w http.ResponseWriter, r *http.Request) (string) {

	p := r.URL.Query().Get(param)
	if p == "" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		msg := "URI parameter:[" + param + "] missing"
		w.Write(FormatErrorMessage(msg))
		return ""
	}
	return p
}

func FormatErrorMessage (msg string) ([]byte) {
	fmt.Println(msg)
	return JSONMarshal(map[string]interface{} {"error" : msg})
}

func FormatSuccessMessage (msg string) ([]byte) {
	fmt.Println(msg)
	return JSONMarshal(map[string]interface{} {"success" : msg})
}