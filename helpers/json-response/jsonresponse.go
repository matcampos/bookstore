package jsonresponse

import (
	"encoding/json"
	"net"
	"net/http"

	errormodel "bookstore/models/error"

	"github.com/go-errors/errors"
)

// ToJSON receives the http.ResponseWriter of the request and the json content in []byte and returns the request with status 200.
func ToJSON(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// ToError is a function to return the stackTrace of an internal error of the system, it receives the http.ResponseWriter of the request, an error interface and a status which is the statusCode that will be returned on the response, if status == 0 the statusCode which will be returned is 500 http.StatusInternalServerError.
func ToError(w http.ResponseWriter, errorStack error, status int) {
	w.Header().Set("Content-Type", "application/json")

	if status != 0 {
		w.WriteHeader(status)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err, ok := errorStack.(net.Error); ok {
		messages := []errormodel.Message{{Pt: err.Error(), En: err.Error()}}
		errors := errormodel.Error{Code: http.StatusInternalServerError, Messages: messages}
		json.NewEncoder(w).Encode(errors)
		return
	}

	if _, ok := errorStack.(*errors.Error); ok {
		messages := []errormodel.Message{{Pt: errorStack.(*errors.Error).ErrorStack(), En: errorStack.(*errors.Error).ErrorStack()}}
		errors := errormodel.Error{Code: http.StatusInternalServerError, Messages: messages}
		json.NewEncoder(w).Encode(errors)
	}

}

// CustomError is a function to return a custom error message, it receives the http.ResponseWriter of the request an errormodel.Error instance and a status which is the statusCode that will be returned on the response.
func CustomError(w http.ResponseWriter, errorM errormodel.Error, status int) {
	w.Header().Set("Content-Type", "application/json")
	if status != 0 {
		w.WriteHeader(status)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(errorM)
}
