package errhandler

import (
	"net/http"
)

var (
	RespDBDataInsertFailure = []byte(`{"error": "db data insert failure"}`)
	RespDBDataAccessFailure = []byte(`{"error": "db data access failure"}`)
	RespDBDataUpdateFailure = []byte(`{"error": "db data update failure"}`)
	RespDBDataRemoveFailure = []byte(`{"error": "db data remove failure"}`)

	RespJSONEncodeFailure = []byte(`{"error": "json encode failure"}`)
	RespJSONDecodeFailure = []byte(`{"error": "json decode failure"}`)

	RespInvalidURLParamID  = []byte(`{"error": "invalid url param-id"}`)
	RespInvalidRequestBody = []byte(`{"error": "invalid request body"}`)
	RespUnauthorized       = []byte(`{"error": "unauthorized"}`)
)

type Error struct {
	Error string `json:"error"`
}

type Errors struct {
	Errors []string `json:"errors"`
}

func ServerError(w http.ResponseWriter, error []byte) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(error)
}

func BadRequest(w http.ResponseWriter, error []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(error)
}

func NotAcceptableRequest(w http.ResponseWriter, error []byte) {
	w.WriteHeader(http.StatusNotAcceptable)
	w.Write(error)
}

func ValidationErrors(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	w.Write(resp)
}

func Unauthorized(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(resp)
}

func Forbidden(w http.ResponseWriter, resp []byte) {
	w.WriteHeader(http.StatusForbidden)
	w.Write(resp)
}
