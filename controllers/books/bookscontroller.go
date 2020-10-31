package bookscontroller

import (
	jsonresponse "bookstore/helpers/json-response"
	booksmodel "bookstore/models/books"
	errormodel "bookstore/models/error"
	successmodel "bookstore/models/success"
	booksrepository "bookstore/repositories/books"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	book := booksmodel.Book{}
	err := decoder.Decode(&book)

	if err != nil {
		jsonresponse.ToError(w, err, 0)
		return
	}

	invalidBook := book.ValidateBookStruct()

	if invalidBook != nil {
		errorMessage := invalidBook.Error()
		messages := []errormodel.Message{errormodel.Message{Pt: errorMessage, En: errorMessage}}
		err := errormodel.Error{Code: http.StatusBadRequest, Messages: messages}
		jsonresponse.CustomError(w, err, http.StatusBadRequest)
		return
	}

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	createBookError := booksrepository.Create(book)

	if createBookError != nil {
		jsonresponse.ToError(w, createBookError, 0)
		return
	}

	success := successmodel.Success{Success: true}

	successJSON, err := json.Marshal(success)

	jsonresponse.ToJson(w, successJSON)
}

func FindAllPaginated(w http.ResponseWriter, r *http.Request) {
	skip, _ := strconv.ParseInt(mux.Vars(r)["skip"], 10, 64)
	limit, _ := strconv.ParseInt(mux.Vars(r)["limit"], 10, 64)

	if limit > 50 {
		limit = 50
	}

	books, findBooksErr := booksrepository.FindAllPaginated(skip, limit)

	if findBooksErr != nil {
		jsonresponse.ToError(w, findBooksErr, 0)
		return
	}

	booksJSON, booksJSONErr := json.Marshal(books)

	if booksJSONErr != nil {
		jsonresponse.ToError(w, booksJSONErr, 0)
		return
	}

	jsonresponse.ToJson(w, booksJSON)

}

func Update(w http.ResponseWriter, r *http.Request) {
	id, convertingError := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if convertingError != nil {
		jsonresponse.ToError(w, convertingError, 0)
		return
	}

	decoder := json.NewDecoder(r.Body)
	updateBook := booksmodel.UpdateBook{}
	err := decoder.Decode(&updateBook)
	if err != nil {
		jsonresponse.ToError(w, err, 0)
		return
	}

	updatedObject, updatedObjectErr := booksrepository.Update(id, updateBook)
	if updatedObjectErr != nil {
		messages := []errormodel.Message{errormodel.Message{Pt: updatedObjectErr.Error(), En: updatedObjectErr.Error()}}
		errorModel := errormodel.Error{Code: http.StatusNotFound, Messages: messages}
		jsonresponse.CustomError(w, errorModel, http.StatusNotFound)
		return
	}

	updatedBookJSON, updatedBookJSONErr := json.Marshal(updatedObject)
	if updatedBookJSONErr != nil {
		jsonresponse.ToError(w, updatedBookJSONErr, 0)
		return
	}

	jsonresponse.ToJson(w, updatedBookJSON)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, convertingError := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if convertingError != nil {
		jsonresponse.ToError(w, convertingError, 0)
		return
	}

	deleteObject, deleteObjectErr := booksrepository.Delete(id)
	if deleteObjectErr != nil {
		messages := []errormodel.Message{errormodel.Message{Pt: deleteObjectErr.Error(), En: deleteObjectErr.Error()}}
		errorModel := errormodel.Error{Code: http.StatusNotFound, Messages: messages}
		jsonresponse.CustomError(w, errorModel, http.StatusNotFound)
		return
	}

	deletedObjectJSON, deletedObjectJSONErr := json.Marshal(deleteObject)
	if deletedObjectJSONErr != nil {
		jsonresponse.ToError(w, deletedObjectJSONErr, 0)
		return
	}

	jsonresponse.ToJson(w, deletedObjectJSON)
}

func FindById(w http.ResponseWriter, r *http.Request) {
	id, convertingError := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if convertingError != nil {
		jsonresponse.ToError(w, convertingError, 0)
		return
	}

	findedObject, findedObjectErr := booksrepository.FindById(id)
	if findedObjectErr != nil {
		messages := []errormodel.Message{errormodel.Message{Pt: findedObjectErr.Error(), En: findedObjectErr.Error()}}
		errorModel := errormodel.Error{Code: http.StatusNotFound, Messages: messages}
		jsonresponse.CustomError(w, errorModel, http.StatusNotFound)
		return
	}

	findedObjectJSON, findedObjectJSONErr := json.Marshal(findedObject)
	if findedObjectJSONErr != nil {
		jsonresponse.ToError(w, findedObjectJSONErr, 0)
		return
	}

	jsonresponse.ToJson(w, findedObjectJSON)
}
