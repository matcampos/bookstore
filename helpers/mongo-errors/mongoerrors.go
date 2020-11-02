package mongoerrors

import (
	errormodel "bookstore/models/error"
	"encoding/hex"
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
	case hex.ErrLength:
		errorModel = errormodel.Error{
			Messages: []errormodel.Message{
				{
					Pt: "O parâmetro id deve ser um ObjectId válido.",
					En: "Id parameter must be a valid ObjectId.",
				},
			},
			Code: http.StatusNotFound,
		}
		statusCode = http.StatusExpectationFailed
	}

	return errorModel, statusCode
}
