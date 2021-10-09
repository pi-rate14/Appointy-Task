package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func setIDS(){
	user_count, err  := UsersCollection.CountDocuments(Ctx, bson.M{}, nil)
	if err != nil {
        log.Fatal(err)
    } 
		ID.userId = int(user_count)+1
	post_count, err  := PostsCollection.CountDocuments(Ctx, bson.M{}, nil)
	if err != nil {
        log.Fatal(err)
    } 
	ID.postId = int(post_count)+1
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}