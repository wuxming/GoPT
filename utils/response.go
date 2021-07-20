package utils

import (
	"net/http"
)

func SuccessMess(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  http.StatusOK,
		"message": message,
		"data":    data,
	}
}
func ErrorMess(message string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  http.StatusInternalServerError,
		"message": message,
		"data":    data,
	}
}
