package main

import "time"

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