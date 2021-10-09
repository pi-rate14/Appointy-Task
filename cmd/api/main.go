package main

import (
	"sync"
	"time"
)

/* Global ID struct to assign User and Post IDs*/

type GlobalID struct {
	userId int
	postId int
}

var ID GlobalID

// sync.Mutex type object to make server thread safe
var lock sync.Mutex

func main() {
	/*Initiating Mongo DB connection*/
	client := openDB()
	defer CloseClientDB(client)	
	/*handling requests*/
	handleRequests()
	time.Sleep(3 * time.Second)
}
