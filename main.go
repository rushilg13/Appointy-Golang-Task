package main

import (
	"appointy/users"
	"appointy/posts"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {

	//POST users
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var newUser users.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		fmt.Println(newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(users.CreateUser(newUser).InsertedID)
	})

	//GET users/{id}
	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/users/")
		fmt.Println(id)
		if id != "" {
			json.NewEncoder(w).Encode(users.GetUser(id))
		}

	})

	//POST posts
	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var newPost posts.Post
		err := json.NewDecoder(r.Body).Decode(&newPost)
		fmt.Println(newPost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(posts.CreatePost(newPost).InsertedID)
	})

	//GET posts/{id}
	http.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/posts/")
		fmt.Println(id)
		if id != "" {
			json.NewEncoder(w).Encode(posts.GetPost(id))
		}
	})

	log.Fatal(http.ListenAndServe(":8085", nil))
}