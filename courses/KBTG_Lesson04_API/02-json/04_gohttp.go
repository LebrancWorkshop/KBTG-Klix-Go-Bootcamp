package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{ID: 1, Name: "AnuchitO", Age: 18},
}

func usersHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		log.Println("POST")
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		u := User{}
		err = json.Unmarshal(body, &u)
		if err != nil {
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		users = append(users, u)
		fmt.Printf("% #v\n", users)

		fmt.Fprintf(w, "hello %s created users", "POST")
		return
	}

	if req.Method == "GET" {
		log.Println("GET")
		b, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

func main() {
	http.HandleFunc("/users", usersHandler)

	log.Println("Server started at :2565")
	log.Fatal(http.ListenAndServe(":2565", nil))
	log.Println("bye bye!")
}
