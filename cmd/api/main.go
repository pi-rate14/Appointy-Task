package main

import (
	"sync"
	"time"
)

type GlobalID struct {
	userId int
	postId int
}

var ID GlobalID

var lock sync.Mutex

func main() {
	client := openDB()
	defer CloseClientDB(client)	
	handleRequests()
	time.Sleep(3 * time.Second)

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
	// var w http.ResponseWriter 
	// var r *http.Request
	// CreateUser(w, r)
	// CreatePost(post)
	// CreatePost(post)
	// CreateUser(user)
	// CreatePost(post)

	// fmt.Println(GetUsers())
	// fmt.Println(GetPosts())
	//FindUserPosts(2)
}
