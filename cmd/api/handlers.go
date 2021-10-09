package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*Handler for 404 errors*/
func handle404(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"404 not found")
}

/*Handle main routing requests*/
func handleRequests() {
	lock.Lock()
    defer lock.Unlock()
	http.HandleFunc("/", handle404)
	http.HandleFunc("/users", handleUsers)
	http.HandleFunc("/users/", GetUser)
	http.HandleFunc("/posts", handlePosts)
	http.HandleFunc("/posts/", GetPost)
	http.HandleFunc("/posts/users/", FindUserPosts)
	log.Fatal(http.ListenAndServe(":8081", nil))
	time.Sleep(1 * time.Microsecond)
}

/*Handler to handle POST request /users */
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

/*Handler to handle POST request /posts */
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
