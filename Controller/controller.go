package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	model "main/Model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://GoLang:Golang123@golang.zb4dkxj.mongodb.net/?retryWrites=true&w=majority"
const dbname = "GoLang"
const colName = "watchlist"

var collection *mongo.Collection

//connect with mongodb

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection sucess")

	collection = client.Database(dbname).Collection(colName)

}

//MongoDB Helpers

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted with ID: ", inserted.InsertedID)

}
func updateOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Record updated ", result.ModifiedCount)
}

func deleteOneMovie(movieID string) (string, bool) {
	id, _ := primitive.ObjectIDFromHex(movieID)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil || result.DeletedCount == 0 {
		//log.Fatal(err)
		return err.Error(), false
	}
	fmt.Println("Deleted ", result.DeletedCount)
	return "Data Deletet Successfully", true
}

func deleteAllMovie() int64 {
	filter := bson.D{{}}
	result, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Record Deleted ", result.DeletedCount)
	return result.DeletedCount
}

func getAllMovie() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)

	}
	defer cursor.Close(context.Background())
	return movies
}

//Actual Controller

func GetAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovie()
	response := model.JsonResponseWithArray{true, "Data fetched Succesfully", allMovies}
	json.NewEncoder(w).Encode(response)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params)
}

func DeletAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	message, status := deleteOneMovie(params["id"])
	if status {
		response := model.JsonResponse{true, message}
		json.NewEncoder(w).Encode(response)
	} else {
		response := model.JsonResponse{false, message}
		json.NewEncoder(w).Encode(response)
	}

}

func DeletALLMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
