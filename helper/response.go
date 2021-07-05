package helper

import (
	"encoding/json"
	"net/http"
)

//ResponseData struct
type ResponseDate struct {
	Date interface{} `json:"data"` //interface{} 任何值都可以赋值给这个类型的变量
	Meta interface{} `json:"meta"`
}
//Message return map data
func Message(status int,message string) map[string]interface{} {
	return map[string]interface{}{
		"status":status,
		"message":message,
	}
}
func Respond(w http.ResponseWriter,data map[string]interface{})  {
	w.Header().Add("Content-Type","application/json")
	if data["status"]==http.StatusInternalServerError {
		w.WriteHeader(data["status"].(int))//请求头是500
		json.NewEncoder(w).Encode(data)
	}else {
		json.NewEncoder(w).Encode(data)
	}
}
func SuccessMess(message string,data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":http.StatusOK,
		"message":message,
		"data":data,
	}
}
func ErrorMess(message string,data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":http.StatusOK,
		"message":message,
		"data":data,
	}
}
