package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type User struct {
	ID int `bson:"_id"`
	Name string `bson:"name"`
	Email string `bson:"email"`
	Password string `bson:"password"`
}

type Post struct {
	ID int `bson:"_id"`
	Caption string `bson:"caption"`
	ImageURL string `bson:"image_url"`
	CreatedAt time.Time `bson:"created_at"`
	UserId  int `bson:"user_id"`
}

type GlobalID struct {
	userId int
	postId int
}

var ID GlobalID

func handle404(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"404 not found")
}

func handleRequests() {
	http.HandleFunc("/", handle404)
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/", GetUser)
	http.HandleFunc("/posts", handlePosts)
	http.HandleFunc("/posts/", GetPost)
	http.HandleFunc("/posts/users/", FindUserPosts)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		CreateUser(w, r)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Bad Request. Id missing"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
}

func handlePosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		CreatePost(w, r)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		resp := make(map[string]string)
		resp["message"] = "Bad Request. Id missing"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
}

func setIDS(){
	user_count, err  := UsersCollection.CountDocuments(Ctx, bson.M{}, nil)
	if err != nil {
        log.Fatal(err)
    } 
	fmt.Println(user_count)
		ID.userId = int(user_count)+1
	post_count, err  := PostsCollection.CountDocuments(Ctx, bson.M{}, nil)
	if err != nil {
        log.Fatal(err)
    } 
	fmt.Println(post_count)
	ID.postId = int(post_count)+1
	
	// fmt.Println(UsersCollection.FindOne(Ctx, bson.M{"$natural": -1})) 
	// fmt.Println(PostsCollection.FindOne(Ctx, bson.M{"$natural": -1})) 
}

func main() {
	client := openDB()
	defer CloseClientDB(client)
	setIDS()
	handleRequests()

	// user := User {
	// 	ID: ID.userId,
	// 	Name: "Apoorva Srivastava",
	// 	Email: "apoorvasrivastava.14@gmail.com",
	// 	Password: "password",
	// }
	// post := Post {
	// 	ID: ID.postId,
	// 	Caption: "Test Caption",
	// 	ImageURL: "www.google.com",
	// 	CreatedAt: time.Now(),
	// 	UserId: ID.userId,
	// }
	
	// CreateUser(user)
	// CreatePost(post)
	// CreatePost(post)
	// CreateUser(user)
	// CreatePost(post)

	// fmt.Println(GetUsers())
	// fmt.Println(GetPosts())
	//FindUserPosts(2)
}
var (
	UsersCollection     *mongo.Collection
	PostsCollection   *mongo.Collection
	Ctx                 = context.TODO()
)



func CloseClientDB(client *mongo.Client ) {
    if client == nil {
        return
    }

    err := client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }

    // TODO optional you can log your closed MongoDB client
    fmt.Println("Connection to MongoDB closed.")
}

/*Setup opens a database connection to mongodb*/
func openDB()  *mongo.Client {
	connectionURI := "mongodb+srv://pirate:pirate2546@cluster0.9u5da.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("Appointy")
	UsersCollection = db.Collection("Users")
	PostsCollection = db.Collection("Posts")
	return client

	
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	//var user User
	user := User {
			ID: ID.userId,
			Name: "Apoorva Srivastava",
			Email: "apoorvasrivastava.14@gmail.com",
			Password: "password",
		}
	user.ID = ID.userId
	result, err := UsersCollection.InsertOne(Ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result.InsertedID)
}

func GetUser(w http.ResponseWriter, r *http.Request){
	var user User
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	objectId,err := strconv.Atoi(id)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("user id: %d", objectId)
	err = UsersCollection.
		FindOne(Ctx, bson.D{{"_id", objectId}}).
		Decode(&user)
	if err != nil {
		log.Println(err)
	}
	//return user, nil
	fmt.Println("User by id endpoint hit")
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request)  {
	var user User
	var users []User

	cursor, err := UsersCollection.Find(Ctx, bson.D{})
	if err != nil {
		defer cursor.Close(Ctx)
		//return users, err
	}

	for cursor.Next(Ctx) {
		err := cursor.Decode(&user)
		if err != nil {
			//return users,err
			fmt.Println(err)
		}
		users = append(users, user)
	}
	fmt.Println("All users endpoint hit")
	json.NewEncoder(w).Encode(users)
	//return users, nil
}

func CreatePost(w http.ResponseWriter, r *http.Request) (string, error) {
	post := Post {
		ID: ID.postId,
		Caption: "Test Caption",
		ImageURL: "www.google.com",
		CreatedAt: time.Now(),
		UserId: ID.userId,
	}
	//.ID = ID.postId
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
		//return post, err
		fmt.Println(err)
	}
	err = PostsCollection.
		FindOne(Ctx, bson.D{{"_id", objectId}}).
		Decode(&post)
	if err != nil {
		//return post, err
		fmt.Println(err)
	}
	fmt.Println(" Post by ID checkpoint hit")
	json.NewEncoder(w).Encode(post)
	//return post, nil
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

