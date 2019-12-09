package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Username string `json:"username"`

	Email string `json:"email,omitempty"`

	Password string `json:"password"`

	ID int `json:"id"`
}
type LUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func testSignIn() {
	url := "http://localhost:8080/v3/signIn"
	contentType := "application/json;charset=utf-8"

	var user LUser
	user.Username = "ap"
	user.Password = "ap"

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json format error:", err)
		return
	}

	body := bytes.NewBuffer(b)

	resp, err := http.Post(url, contentType, body)
	if err != nil {
		fmt.Println("Post failed:", err)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read failed:", err)
		return
	}

	fmt.Println("header:", resp.Header)
	fmt.Println("content:", string(content))
}
func testSignUp() {

	url := "http://localhost:4200/v3/signUp"
	contentType := "application/json;charset=utf-8"

	var user User
	user.Username = "ap1"
	user.Password = "ap1"
	user.Email = "api_test@qq.com"

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json format error:", err)
		return
	}

	body := bytes.NewBuffer(b)

	resp, err := http.Post(url, contentType, body)
	if err != nil {
		fmt.Println("Post failed:", err)
		return
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read failed:", err)
		return
	}

	fmt.Println("header:", resp.Header)
	fmt.Println("content:", string(content))

}

func main() {
	testSignUp()
	//testSignIn()
}
