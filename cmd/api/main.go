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

	// CreateUser(w, r)
	// CreatePost(post)
	// CreatePost(post)
	// CreateUser(user)
	// CreatePost(post)

	// fmt.Println(GetUsers())
	// fmt.Println(GetPosts())
	//FindUserPosts(2)
}
