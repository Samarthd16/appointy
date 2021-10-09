package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoField struct {
	FieldStr  string `json: "Field Str"`
	FieldInt  int    `json: "Field Int"`
	FieldBool bool   `json: "Field Bool"`
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Post struct {
	Id         string `json:"id"`
	Caption    string `json:"caption"`
	Image_URL  string `json:"Image_URL"`
	Time_Stamp string `json:"Time_Stamp"`
}

type Users []User
type Posts []Post
type MongoFields []MongoField

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func user(w http.ResponseWriter, r *http.Request) {

	users := Users{
		User{

			Id:       "Test Title",
			Name:     "Test desciption",
			Password: "Hello World",
			Email:    "samarthdsankhla16@gmail.com"},
	}

	fmt.Println("endpoint hit: All articles Endpoint")
	json.NewEncoder(w).Encode(users)

	fmt.Fprintf(w, "Welcome to the Users homepage!")
	fmt.Println("Endpoint Hit: homePage")
}

func mongofield(w http.ResponseWriter, r *http.Request) {

	mongoFields := MongoFields{
		MongoField{
			FieldStr:  "This is our first data and its very important",
			FieldInt:  826482746,
			FieldBool: true},
	}

	fmt.Println("endpoint hit: All articles Endpoint")
	json.NewEncoder(w).Encode(mongoFields)

	fmt.Fprintf(w, "Welcome to the Users homepage!")
	fmt.Println("Endpoint Hit: homePage")
}

func post(w http.ResponseWriter, r *http.Request) {

	posts := Posts{
		Post{Id: "45", Caption: "This is a beautifull world", Image_URL: "https://cdn.pixabay.com/photo/2015/04/23/22/00/tree-736885__480.jpg", Time_Stamp: "22:34"},
	}

	fmt.Println("endpoint hit: All articles Endpoint")
	json.NewEncoder(w).Encode(posts)

	fmt.Fprintf(w, "Welcome to the Users homepage!")
	fmt.Println("Endpoint Hit: homePage")
}

// func AllUserEndPoint(w http.ResponseWriter, r * http.Request) {
//     user, err: = dao.FindAll()
//     if err != nil {
//         respondWithError(w, http.StatusInternalServerError, err.Error())
//         return
//     }
//     respondWithJson(w, http.StatusOK, movies)
// }

// func FindMovieEndpoint(w http.ResponseWriter, r * http.Request) {
//     params: = mux.Vars(r)
//     test,
//     err: = dao.FindById(params["id"])
//     if err != nil {
//         respondWithError(w, http.StatusBadRequest, "Invalid User ID")
//         return
//     }
//     respondWithJson(w, http.StatusOK, user)
// }

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/user", user)
	http.HandleFunc("/post", post)
	http.HandleFunc("/mongoField", mongofield)
	//http.AllUserEndPoint("/Alluser",AllUserEndPoint)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

//__________________________________________________________________________________________________________________________

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	col := client.Database("Instagram").Collection("db3")
	fmt.Println("Collection Type: ", reflect.TypeOf(col), "\n")

	oneDoc := Users{
		User{Id: "101",
			Name:     "Samarth",
			Password: "sama@123",
			Email:    "samarthdsankhla16@gmail.com"},
	}

	fmt.Println("oneDoc Type: ", reflect.TypeOf(oneDoc), "\n")

	result, insertErr := col.InsertOne(ctx, oneDoc)
	if insertErr != nil {
		fmt.Println("InsertONE Error:", insertErr)
		os.Exit(1)
	} else {
		fmt.Println("InsertOne() result type: ", reflect.TypeOf(result))
		fmt.Println("InsertOne() api result type: ", result)

		newID := result.InsertedID
		fmt.Println("InsertedOne(), newID", newID)
		fmt.Println("InsertedOne(), newID type:", reflect.TypeOf(newID))

	}

	handleRequests()
}
