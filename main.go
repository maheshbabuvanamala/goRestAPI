package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tkanos/gonfig"
)

type Post struct {
	ID   string `json:"id"`
	Name string `json:"Name"`
}
type Configuration struct {
	db struct {
		port     string
		username string
		password string
		host     string
	}
}

func main() {
	configuration := Configuration{}
	err := gonfig.GetConf("config.json", &configuration)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(configuration.db.host, configuration.db.username, configuration.db.password, configuration.db.port)
	router := mux.NewRouter()

	router.HandleFunc("/posts", getPosts).Methods("GET")
	Getconnection("root", "root@1234", "logparser")
	http.ListenAndServe(":8000", router)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := db.Query("SELECT ID,NAME from TEST1")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer result.Close()
	var posts []Post
	for result.Next() {
		var post Post
		err := result.Scan(&post.ID, &post.Name)
		if err != nil {
			panic(err.Error())
		}
		posts = append(posts, post)
	}
	json.NewEncoder(w).Encode(posts)
}
