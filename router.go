package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	FirstName, LastName string
	Age                 int
}

type UserRegistrationInfo struct {
	FirstName, LastName string
	Email               string
	Age                 int
	Password            string
}

type ErrorInfo struct {
	errors map[string]string
}

func AddHandlers() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "this is a website server by a GO HTTP server.")
		fmt.Fprintf(w, "METHOD %v", r.Method)
	})
	//get
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		user := User{"Harry", "Zhang", 43}
		b, err := json.Marshal(user)
		if err == nil {
			w.Write(b)
			w.WriteHeader(200)
		} else {
			w.WriteHeader(403)
		}
	})
	//post
	http.HandleFunc("/reg", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		//TODO: Validation
		age, _ := strconv.Atoi(r.PostForm.Get("Age"))
		registeredUser := User{r.PostForm.Get("FirstName"), r.PostForm.Get("LastName"), age}
		b, e := json.Marshal(registeredUser)
		if e != nil {
			w.WriteHeader(400)
			return
		}
		w.Write(b)
	})
}
