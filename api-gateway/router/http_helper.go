package router

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type IHttpHelper interface {
	JsonParse(body io.Reader, bindStruct any) error
	JsonResponse(w http.ResponseWriter, statusCode int, response H)
	JsonResponseMessage(w http.ResponseWriter, statusCode int, message any)
}

type HttpHelper struct{}

func NewHttpHelper() HttpHelper {
	return HttpHelper{}
}

func (helper HttpHelper) JsonParse(body io.Reader, bindStruct any) error {
	read, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	return json.Unmarshal(read, bindStruct)
}

func (helper HttpHelper) JsonResponse(w http.ResponseWriter, statusCode int, response H) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func (helper HttpHelper) JsonResponseMessage(w http.ResponseWriter, statusCode int, message any) {
	helper.JsonResponse(w, statusCode, H{
		"message": message,
	})
}
