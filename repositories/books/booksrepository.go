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

// Create receives the param book which is an instance of Book struct to create a book document on database, it returns an error if something wrong happens.
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

// FindAllPaginated receives two parameters skip an int64 and limit which is an int64 too, it returns a booksmodel.BooksPaginatedList instance with the found data on database or return an error if something wrong happens.
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

// count function receives a mongo collection instance and return an int64 which is the number of documents found in this collection.
func count(collection *mongo.Collection) (int64, error) {
	count, countError := collection.CountDocuments(context.Background(), bson.M{})

	if countError != nil {
		return 0, countError
	}

	return count, nil
}

// findPaginated function receives a mongo collection instance and the parameters skip and limit which both are an int64, it returns a Book struct array or an error if something wrong happens.
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

// Update function receives three parameters: id is an objectId of the document which will be updated and updateBook is an UpdateBook struct instance with the parameters sent in the body of the request, it returns Book struct with the changed book document or an error in case of none result.
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

// updateFilter function receives the updateBook parameter which is an instance of UpdateBook and return a bson.M interface object with the data of updateBook parameter.
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

// Delete function receives the parameter "id" which is the id of the document to be deleted, it returns a Book struct instance with the deleted document data or an error instance whether something wrong happens or none result was found.
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

// FindByID function receives the parameter "id" which is the id of the document to be found, it returns a Book struct instance with the found document data or an error instance whether something wrong happens or none result was found.
func FindByID(id primitive.ObjectID) (booksmodel.Book, error) {
	foundBook := booksmodel.Book{}

	conn, connErr := db.Connect()
	if connErr != nil {
		return foundBook, connErr
	}
	defer conn.Client.Disconnect(conn.Ctx)

	collection := conn.Client.Database(os.Getenv("DATABASE")).Collection("books")
	filter := bson.M{"_id": id}

	foundResult := collection.FindOne(context.Background(), filter)

	decodeErr := foundResult.Decode(&foundBook)
	if decodeErr != nil {
		return foundBook, decodeErr
	}

	return foundBook, nil
}
