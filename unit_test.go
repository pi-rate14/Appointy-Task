package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_GetPostByUserId(t *testing.T){
	data, _ := getXML("http://localhost:8081/posts/users/8?page=1")
	expected := `[{"_id":3,"caption":"","created_at":"0001-01-01T05:53:28+05:53","image_url":"","user_id":8},{"_id":4,"caption":"","created_at":"0001-01-01T05:53:28+05:53","image_url":"","user_id":8}]`
	if data == expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			data, expected)
	}
}
func Test_GetUserById(t *testing.T){
	data, _ := getXML("http://localhost:8081/users/1")
	expected := `{"ID":1,"Name":"Apoorva Srivastava","Email":"apoorvasrivastava.14@gmail.com","Password":"53+\ufffd\ufffd3\ufffd蟼\u0011\ufffdݎ5,\u0001\ufffd\u0014\n\ufffd~\ufffd7\ufffd\ufffd\u001d\u001f\u0014\ufffd\ufffd޷\u000c\ufffd"}`
	if data == expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			data, expected)
	}
}
func Test_GetPostById(t *testing.T){
	data, err := getXML("http://localhost:8081/posts/1")
	if err!=nil {
		t.Errorf("error")
	}
	expected := `{"ID":1,"Caption":"Test Caption","ImageURL":"www.google.com","CreatedAt":"2021-10-08T23:58:26.188Z","UserId":1}`
	if data == expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			data, expected)
	}
}

func getXML(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", fmt.Errorf("GET error: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("Status error: %v", resp.StatusCode)
    }

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("Read body: %v", err)
    }

    return string(data), nil
}