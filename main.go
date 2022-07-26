package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)

	http.ListenAndServe("127.0.0.1:8080", nil)
}

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	u := User{
		Name: "Harish",
		Age:  23,
	}

	err := json.NewEncoder(w).Encode(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "Application/json")

}

func postHandler(w http.ResponseWriter, r *http.Request) {

	u := User{}

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		fmt.Println(err)
	}

	fmt.Println(u)

	if err := json.NewEncoder(w).Encode(u); err != nil {
		fmt.Println(err)
	}
}
