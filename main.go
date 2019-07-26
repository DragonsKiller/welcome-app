package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Text   string  `json:"text"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"first-name"`
	LastName  string `json:"last-name"`
}

//Init posts var as a slice Post struct
var posts []Post

// Get ALL Posts
func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Get Single Post
func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get Params
	//Loop through books and find with id
	for _, item := range posts {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Post{})
}

// Create New Post
func createPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = strconv.Itoa(rand.Intn(1000000000)) //Mock ID - not safe
	posts = append(posts, post)
	json.NewEncoder(w).Encode(post)
}

// Update Post
func updatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
      var post Post
      post.ID = item.ID
			_ = json.NewDecoder(r.Body).Decode(&post)
			posts = append(posts, post)
			json.NewEncoder(w).Encode(post)
			return
		}
	}
	json.NewEncoder(w).Encode(posts)
}

// Delete Post
func deletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range posts {
		if item.ID == params["id"] {
			posts = append(posts[:index], posts[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(posts)
}

func main() {
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	posts = append(posts, Post{ID: "1", Title: "Post 1", Text: "Just a simple text", Author: &Author{FirstName: "Igor", LastName: "Example"}})
	posts = append(posts, Post{ID: "2", Title: "Post 2", Text: "Just a simple text", Author: &Author{FirstName: "Steve", LastName: "Examplee"}})

	//Endpoints
	r.HandleFunc("/posts", getPosts).Methods("GET")
	r.HandleFunc("/posts/{id}", getPost).Methods("GET")
	r.HandleFunc("/posts", createPost).Methods("POST")
	r.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	r.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
