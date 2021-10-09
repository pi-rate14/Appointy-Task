package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

/*POST request 		Create an User*/
func CreateUser(w http.ResponseWriter, r *http.Request) {
	/*setting the user ID*/
	setIDS()

	/* temporary user entry to seed database */

	//var user User
	// user := User {
	// 		ID: ID.userId,
	// 		Name: "Apoorva Srivastava",
	// 		Email: "apoorvasrivastava.14@gmail.com",
	// 		Password: "password",
	// 	}

	var user User
	user.ID = ID.userId
	/*encryptng user password*/
	md5Key := createHash(user.Email)
	hashedPassword := encrypt([]byte(user.Password),md5Key)
	user.Password = string(hashedPassword)
	
	result, err := UsersCollection.InsertOne(Ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result.InsertedID)
}

/* GET Request		Get one User using ID*/
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	objectId,err := strconv.Atoi(id)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Bad Request. ID invalid."
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
	err = UsersCollection.
		FindOne(Ctx, bson.D{{Key:"_id",Value: objectId}}).
		Decode(&user)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(" User by id endpoint hit")
	
	json.NewEncoder(w).Encode(user)
}

/*function to get all the users from the database */

// func GetUsers(w http.ResponseWriter, r *http.Request)  {
// 	var user User
// 	var users []User

// 	cursor, err := UsersCollection.Find(Ctx, bson.D{})
// 	if err != nil {
// 		defer cursor.Close(Ctx)
// 		//return users, err
// 	}

// 	for cursor.Next(Ctx) {
// 		err := cursor.Decode(&user)
// 		if err != nil {
// 			//return users,err
// 			fmt.Println(err)
// 		}
// 		users = append(users, user)
// 	}
// 	fmt.Println("All users endpoint hit")
// 	json.NewEncoder(w).Encode(users)
// 	//return users, nil
// }
