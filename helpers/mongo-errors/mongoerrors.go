package mongoerrors

import (
	errormodel "bookstore/models/error"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

// HandleMongoErrors receives an error interface and return two params, the error model and the status code.
func HandleMongoErrors(err error) (errormodel.Error, int) {
	var errorModel errormodel.Error
	var statusCode int = http.StatusInternalServerError

	switch err {
	case mongo.ErrNoDocuments:
		errorModel = errormodel.Error{
			Messages: []errormodel.Message{
				{
					Pt: "Resultado não encontrado.",
					En: "Result not found",
				},
			},
			Code: http.StatusNotFound,
		}
		statusCode = http.StatusNotFound
	}

	return errorModel, statusCode
}
