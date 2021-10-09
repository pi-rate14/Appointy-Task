package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*POST request 		Create a Post*/
func CreatePost(w http.ResponseWriter, r *http.Request) (string, error) {
	/*setting the Post ID*/
	setIDS()

	/* temporary user entry to seed database */

	// post := Post {
	// 	ID: ID.postId,
	// 	Caption: "Test Caption",
	// 	ImageURL: "www.google.com",
	// 	CreatedAt: time.Now(),
	// 	UserId: ID.userId,
	// }

	var post Post
	post.ID = ID.postId
	post.UserId = ID.userId
	post.CreatedAt = time.Now()
	result, err := PostsCollection.InsertOne(Ctx, post)
	if err != nil {
		return "0", err
	}
	return 	fmt.Sprintf("%v", result.InsertedID), err
}

/* GET Request		Get one Post using ID*/
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

/*GET request		Get all posts by a user using user ID */
func FindUserPosts(w http.ResponseWriter, r *http.Request)  {
	query := r.URL.Query()
	
	var page int
	filters,ok := query["page"] 
    if len(filters) == 0 || !ok{
		query.Set("page","1")
		filters = query["page"] 
    }
	findOptions := options.Find()
	page,err := strconv.Atoi(filters[0]) 
	var limit int64 = 2
	if err != nil{
		fmt.Println(err)
	}
	findOptions.SetSkip((int64(page) - 1) * limit)
	findOptions.SetLimit(limit)
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
	filterCursor, err := PostsCollection.Find(Ctx, bson.M{"user_id": user_id}, findOptions)
if err != nil {
    log.Fatal(err)
}
if err = filterCursor.All(Ctx, &posts); err != nil {
    log.Fatal(err)
}
fmt.Println(" Post by User ID checkpoint hit")
	json.NewEncoder(w).Encode(posts)
}

/* Function to get all posts

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
*/