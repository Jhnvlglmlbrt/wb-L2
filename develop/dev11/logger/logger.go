package logger

import (
	"encoding/json"
	"net/http"

	"github.com/Jhnvlglmlbrt/develop/dev11/response"
)

// Общая функция для логгирования ошибок и успешных ответов
func Logger(w http.ResponseWriter, statusCode int, statusMessage string, data interface{}) {
	w.WriteHeader(statusCode)

	var details interface{}
	if statusCode >= 400 {
		// Если статус код указывает на ошибку, формируем структуру для ошибки
		details = response.Error{
			Code:   statusCode,
			Status: statusMessage,
		}
	} else {
		// Иначе формируем структуру для успешного ответа
		details = response.Response{
			Code:   statusCode,
			Status: statusMessage,
			Data:   data,
		}
	}

	response, _ := json.MarshalIndent(details, "", "\t")
	w.Write(response)
}
