package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string
	Error  *ResponseError
	Result json.RawMessage
}

type ResponseError struct {
	Code    int
	Message string
	Details []string
}

func (e *ResponseError) Error() string {
	j, err := json.Marshal(e)
	if err != nil {
		return "Response Error :" + err.Error()
	}
	return string(j)
}

func Fail(w http.ResponseWriter, errCode int, msg string, details ...string) {
	r := &Response{
		Status: "StatusFail",
		Error: &ResponseError{
			Code:    errCode,
			Message: msg,
			Details: details,
		},
	}

	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(errCode)
	w.Write(j)

}

func Send(w http.ResponseWriter, status int, result interface{}) {
	rj, err := json.Marshal(result)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	r := &Response{
		Status: "StatusOK",
		Result: rj,
	}

	j, err := json.Marshal(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(j)

}
