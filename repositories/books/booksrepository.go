package booksrepository

import (
	db "bookstore/config"
	booksmodel "bookstore/models/books"
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Create(book booksmodel.Book) error {

	conn, connErr := db.Connect()

	if connErr != nil {
		return connErr
	}
	defer conn.Client.Disconnect(conn.Ctx)

	book.ID = primitive.NewObjectID()

	collection := conn.Client.Database(os.Getenv("DATABASE")).Collection("books")

	_, insertionErr := collection.InsertOne(context.TODO(), book)

	if insertionErr != nil {
		return insertionErr
	}

	return nil
}

func FindAllPaginated(skip int64, limit int64) (booksmodel.BooksPaginatedList, error) {
	booksPaginatedList := booksmodel.BooksPaginatedList{
		Books: []booksmodel.Book{},
		Count: 0,
	}

	conn, connErr := db.Connect()
	if connErr != nil {
		return booksPaginatedList, connErr
	}
	defer conn.Client.Disconnect(conn.Ctx)

	collection := conn.Client.Database(os.Getenv("DATABASE")).Collection("books")

	count, countError := count(collection)
	if countError != nil {
		return booksPaginatedList, countError
	}

	books, booksError := findPaginated(collection, skip, limit)
	if booksError != nil {
		return booksPaginatedList, booksError
	}

	booksPaginatedList = booksmodel.BooksPaginatedList{Count: count, Books: books}
	return booksPaginatedList, nil
}

func count(collection *mongo.Collection) (int64, error) {
	count, countError := collection.CountDocuments(context.Background(), bson.M{})

	if countError != nil {
		return 0, countError
	}

	return count, nil
}

func findPaginated(collection *mongo.Collection, skip int64, limit int64) ([]booksmodel.Book, error) {
	findOptions := options.FindOptions{Skip: &skip, Limit: &limit}

	cursor, cursorError := collection.Find(context.Background(), bson.M{}, &findOptions)
	if cursorError != nil {
		return nil, cursorError
	}

	var books []booksmodel.Book = []booksmodel.Book{}

	for cursor.Next(context.Background()) {
		var book booksmodel.Book
		decodeErr := cursor.Decode(&book)
		if decodeErr != nil {
			return books, decodeErr
		}
		books = append(books, book)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(context.Background())

	return books, nil
}

func Update(id primitive.ObjectID, updateBook booksmodel.UpdateBook) (booksmodel.Book, error) {
	book := booksmodel.Book{}

	conn, connErr := db.Connect()
	if connErr != nil {
		return book, connErr
	}
	defer conn.Client.Disconnect(conn.Ctx)

	collection := conn.Client.Database(os.Getenv("DATABASE")).Collection("books")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": updateFilter(updateBook)}
	after := options.After
	updateOptions := options.FindOneAndUpdateOptions{ReturnDocument: &after}

	updateResult := collection.FindOneAndUpdate(context.Background(), filter, update, &updateOptions)
	decodeErr := updateResult.Decode(&book)
	if decodeErr != nil {
		return book, decodeErr
	}

	return book, nil
}

func updateFilter(updateBook booksmodel.UpdateBook) bson.M {
	bsonUpdate := bson.M{
		"updatedAt": time.Now(),
	}

	if updateBook.Name != "" {
		bsonUpdate["name"] = updateBook.Name
	}

	if updateBook.Author != "" {
		bsonUpdate["author"] = updateBook.Author
	}

	if updateBook.Genre != "" {
		bsonUpdate["genre"] = updateBook.Genre
	}

	if updateBook.Price != 0 {
		bsonUpdate["price"] = updateBook.Price
	}

	if updateBook.Pages != 0 {
		bsonUpdate["pages"] = updateBook.Pages
	}

	return bsonUpdate
}

func Delete(id primitive.ObjectID) (booksmodel.Book, error) {
	deletedBook := booksmodel.Book{}

	conn, connErr := db.Connect()
	if connErr != nil {
		return deletedBook, connErr
	}
	defer conn.Client.Disconnect(conn.Ctx)

	collection := conn.Client.Database(os.Getenv("DATABASE")).Collection("books")
	filter := bson.M{"_id": id}

	deleteResult := collection.FindOneAndDelete(context.Background(), filter)

	decodeErr := deleteResult.Decode(&deletedBook)
	if decodeErr != nil {
		return deletedBook, decodeErr
	}

	return deletedBook, nil
}

func FindById(id primitive.ObjectID) (booksmodel.Book, error) {
	findedBook := booksmodel.Book{}

	conn, connErr := db.Connect()
	if connErr != nil {
		return findedBook, connErr
	}
	defer conn.Client.Disconnect(conn.Ctx)

	collection := conn.Client.Database(os.Getenv("DATABASE")).Collection("books")
	filter := bson.M{"_id": id}

	findResult := collection.FindOne(context.Background(), filter)

	decodeErr := findResult.Decode(&findedBook)
	if decodeErr != nil {
		return findedBook, decodeErr
	}

	return findedBook, nil
}
