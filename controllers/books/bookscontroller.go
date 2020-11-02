package bookscontroller

import (
	jsonresponse "bookstore/helpers/json-response"
	mongoerrors "bookstore/helpers/mongo-errors"
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

// Create receives the http.ResponseWriter and the *http.Request and returns by the http.ResponseWriter the message {success: true} or an error if something wrong happens.
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
		messages := []errormodel.Message{
			{
				Pt: errorMessage,
				En: errorMessage,
			},
		}
		err := errormodel.Error{Code: http.StatusBadRequest, Messages: messages}
		jsonresponse.CustomError(w, err, http.StatusBadRequest)
		return
	}

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	createBookError := booksrepository.Create(book)

	if createBookError != nil {
		mongoError, statusCode := mongoerrors.HandleMongoErrors(createBookError)
		if len(mongoError.Messages) > 0 {
			jsonresponse.CustomError(w, mongoError, statusCode)
			return
		}

		jsonresponse.ToError(w, createBookError, 0)
		return
	}

	success := successmodel.Success{Success: true}

	successJSON, err := json.Marshal(success)

	jsonresponse.ToJSON(w, successJSON)
}

// FindAllPaginated receives the http.ResponseWriter and the *http.Request and returns by the http.ResponseWriter a booksmodel.BooksPaginatedList struct.
func FindAllPaginated(w http.ResponseWriter, r *http.Request) {
	skip, _ := strconv.ParseInt(mux.Vars(r)["skip"], 10, 64)
	limit, _ := strconv.ParseInt(mux.Vars(r)["limit"], 10, 64)

	if limit > 50 {
		limit = 50
	}

	booksPaginatedList, findBooksErr := booksrepository.FindAllPaginated(skip, limit)

	if findBooksErr != nil {
		mongoError, statusCode := mongoerrors.HandleMongoErrors(findBooksErr)
		if len(mongoError.Messages) > 0 {
			jsonresponse.CustomError(w, mongoError, statusCode)
			return
		}

		jsonresponse.ToError(w, findBooksErr, 0)
		return
	}

	booksPaginatedJSON, booksPaginatedJSONErr := json.Marshal(booksPaginatedList)

	if booksPaginatedJSONErr != nil {
		jsonresponse.ToError(w, booksPaginatedJSONErr, 0)
		return
	}

	jsonresponse.ToJSON(w, booksPaginatedJSON)

}

// Update receives the http.ResponseWriter, the *http.Request and returns by the http.ResponseWriter the updated object or an errormodel.Error struct with status code 404.
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
		mongoError, statusCode := mongoerrors.HandleMongoErrors(updatedObjectErr)
		if len(mongoError.Messages) > 0 {
			jsonresponse.CustomError(w, mongoError, statusCode)
			return
		}

		jsonresponse.ToError(w, updatedObjectErr, 0)
		return
	}

	updatedBookJSON, updatedBookJSONErr := json.Marshal(updatedObject)
	if updatedBookJSONErr != nil {
		jsonresponse.ToError(w, updatedBookJSONErr, 0)
		return
	}

	jsonresponse.ToJSON(w, updatedBookJSON)
}

// Delete receives the http.ResponseWriter and the *http.Request and returns by the http.ResponseWriter the deleted object or an errormodel.Error struct with status code 404.
func Delete(w http.ResponseWriter, r *http.Request) {
	id, convertingError := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if convertingError != nil {
		jsonresponse.ToError(w, convertingError, 0)
		return
	}

	deleteObject, deleteObjectErr := booksrepository.Delete(id)
	if deleteObjectErr != nil {
		mongoError, statusCode := mongoerrors.HandleMongoErrors(deleteObjectErr)
		if len(mongoError.Messages) > 0 {
			jsonresponse.CustomError(w, mongoError, statusCode)
			return
		}

		jsonresponse.ToError(w, deleteObjectErr, 0)
		return
	}

	deletedObjectJSON, deletedObjectJSONErr := json.Marshal(deleteObject)
	if deletedObjectJSONErr != nil {
		jsonresponse.ToError(w, deletedObjectJSONErr, 0)
		return
	}

	jsonresponse.ToJSON(w, deletedObjectJSON)
}

// FindByID receives the http.ResponseWriter and the *http.Request and returns by the http.ResponseWriter the found object or an errormodel.Error struct with status code 404.
func FindByID(w http.ResponseWriter, r *http.Request) {
	id, convertingError := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if convertingError != nil {
		jsonresponse.ToError(w, convertingError, 0)
		return
	}

	foundObject, foundObjectErr := booksrepository.FindByID(id)
	if foundObjectErr != nil {
		mongoError, statusCode := mongoerrors.HandleMongoErrors(foundObjectErr)
		if len(mongoError.Messages) > 0 {
			jsonresponse.CustomError(w, mongoError, statusCode)
			return
		}

		jsonresponse.ToError(w, foundObjectErr, http.StatusNotFound)
		return
	}

	foundObjectJSON, foundObjectJSONErr := json.Marshal(foundObject)
	if foundObjectJSONErr != nil {
		jsonresponse.ToError(w, foundObjectJSONErr, 0)
		return
	}

	jsonresponse.ToJSON(w, foundObjectJSON)
}
