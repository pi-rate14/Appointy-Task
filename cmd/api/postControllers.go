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


func CreatePost(w http.ResponseWriter, r *http.Request) (string, error) {
	setIDS()
	// post := Post {
	// 	ID: ID.postId,
	// 	Caption: "Test Caption",
	// 	ImageURL: "www.google.com",
	// 	CreatedAt: time.Now(),
	// 	UserId: ID.userId,
	// }
	var post Post
	post.UserId = ID.userId
	result, err := PostsCollection.InsertOne(Ctx, post)
	if err != nil {
		return "0", err
	}
	return 	fmt.Sprintf("%v", result.InsertedID), err
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	var post Post
	id := strings.TrimPrefix(r.URL.Path, "/posts/")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Bad Request. ID missing."
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
	objectId,err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
	}
	err = PostsCollection.
		FindOne(Ctx, bson.D{{Key:"_id",Value: objectId}}).
		Decode(&post)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(" Post by ID checkpoint hit")
	json.NewEncoder(w).Encode(post)
}

func GetPosts() ([]Post, error) {
	var post Post
	var posts []Post

	cursor, err := PostsCollection.Find(Ctx, bson.D{})
	if err != nil {
		defer cursor.Close(Ctx)
		return posts, err
	}

	for cursor.Next(Ctx) {
		err := cursor.Decode(&post)
		if err != nil {
			return posts,err
		}
		posts = append(posts, post)
	}

	return posts, nil
}



// func FindUserPosts(user_id int) ([]Post, error) {
// 	matchStage := bson.D{{"$match", bson.D{{"user_id", user_id}}}}

// 	lookupStage := bson.D{{"$lookup",
// 		bson.D{{"from", "Posts"},
// 			{"localField", "_id"},
// 			{"foreignField", "user_id"},
// 			{"as", "Posts"}}}}

// 	showLoadedCursor, err := PostsCollection.Aggregate(Ctx,
// 		mongo.Pipeline{matchStage, lookupStage})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var a []AuthorBooks
// 	if err = showLoadedCursor.All(Ctx, &a); err != nil {
// 		return nil, err

// 	}
// 	return a[0].Posts, err
// }

func FindUserPosts(w http.ResponseWriter, r *http.Request)  {
	var posts []bson.M
	id := strings.TrimPrefix(r.URL.Path, "/posts/users/")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Bad Request. ID missing."
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
	user_id,err := strconv.Atoi(id)
	if err != nil{
		log.Fatal(err)
	}
	filterCursor, err := PostsCollection.Find(Ctx, bson.M{"user_id": user_id})
if err != nil {
    log.Fatal(err)
}
//var episodesFiltered []bson.M
if err = filterCursor.All(Ctx, &posts); err != nil {
    log.Fatal(err)
}
fmt.Println(" Post by ID checkpoint hit")
	json.NewEncoder(w).Encode(posts)
}
