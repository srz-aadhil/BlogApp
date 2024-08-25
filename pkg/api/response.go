package api

import (
	"encoding/json"
	"net/http"
)

const (
	StatusOK   = "ok"
	StatusFail = "not ok"
)

type Response struct {
	Status string          `json:"status"`
	Error  *ResponseError  `json:"error,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
}

type ResponseError struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func (e *ResponseError) Error() string {
	j, err := json.Marshal(e)
	if err != nil {
		return "Response Error :" + err.Error()
	}
	return string(j)
}

// Fail sends an unsuccessful JSON with the standard failure format
func Fail(w http.ResponseWriter, errCode int, msg string, details ...string) {
	r := &Response{
		Status: StatusFail,
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

// Success sends a successful JSON response using the standard success format
func Success(w http.ResponseWriter, status int, result interface{}) {
	rj, err := json.Marshal(result)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	r := &Response{
		Status: StatusOK,
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
