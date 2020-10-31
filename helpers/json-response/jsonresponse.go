package jsonresponse

import (
	"encoding/json"
	"net"
	"net/http"
	"reflect"

	errormodel "bookstore/models/error"

	"github.com/go-errors/errors"
)

func ToJson(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func ToError(w http.ResponseWriter, errorStack error, status int) {
	w.Header().Set("Content-Type", "application/json")

	if status != 0 {
		w.WriteHeader(status)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err, ok := errorStack.(net.Error); ok {
		messages := []errormodel.Message{errormodel.Message{Pt: err.Error(), En: err.Error()}}
		errors := errormodel.Error{Code: http.StatusInternalServerError, Messages: messages}
		json.NewEncoder(w).Encode(errors)
		return
	}

	if reflect.TypeOf(errorStack).String() == "*errors.Error" {
		messages := []errormodel.Message{errormodel.Message{Pt: errorStack.(*errors.Error).ErrorStack(), En: errorStack.(*errors.Error).ErrorStack()}}
		errors := errormodel.Error{Code: http.StatusInternalServerError, Messages: messages}
		json.NewEncoder(w).Encode(errors)
	}

	if reflect.TypeOf(errorStack).String() == "*errors.errorString" {
		messages := []errormodel.Message{errormodel.Message{Pt: errorStack.Error(), En: errorStack.Error()}}
		errors := errormodel.Error{Code: http.StatusInternalServerError, Messages: messages}
		json.NewEncoder(w).Encode(errors)
	}
}

func CustomError(w http.ResponseWriter, errorM errormodel.Error, status int) {
	w.Header().Set("Content-Type", "application/json")
	if status != 0 {
		w.WriteHeader(status)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(errorM)
}
