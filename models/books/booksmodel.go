package booksmodel

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID        primitive.ObjectID `bson:"_id, omitempty" json:"_id,omitempty"`
	Name      string             `bson:"name" json:"name,omitempty"`
	Author    string             `bson:"author" json:"author,omitempty"`
	Genre     string             `bson:"genre" json:"genre,omitempty"`
	Price     float64            `bson:"price" json:"price,omitempty"`
	Pages     int                `bson:"pages" json:"pages,omitempty"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt,omitempty"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt,omitempty"`
}

type UpdateBook struct {
	Name      string    `bson:"name" json:"name,omitempty"`
	Author    string    `bson:"author" json:"author,omitempty"`
	Genre     string    `bson:"genre" json:"genre,omitempty"`
	Price     float64   `bson:"price" json:"price,omitempty"`
	Pages     int       `bson:"pages" json:"pages,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt" json:"updatedAt,omitempty"`
}

type BooksPaginatedList struct {
	Count int64  `bson:"count, omitempty" json:"count"`
	Books []Book `bson:"books, omitempty" json:"books"`
}

func (b Book) ValidateBookStruct() error {
	return validation.ValidateStruct(&b,
		// Name cannot be empty.
		validation.Field(&b.Name, validation.Required),
		// Author cannot be empty.
		validation.Field(&b.Author, validation.Required),
		// Genre cannot be empty.
		validation.Field(&b.Genre, validation.Required),
		// Price cannot be empty.
		validation.Field(&b.Price, validation.Required),
		// Pages cannot be empty.
		validation.Field(&b.Pages, validation.Required),
	)
}
