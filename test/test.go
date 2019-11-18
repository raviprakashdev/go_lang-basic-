package main

import (
	"fmt"
	"net/http"
	"log"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login",login).Methods("POST")
	router.HandleFunc("/logout",logout).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000",router))
	fmt.Printf("hello, world\n")
}

func login(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"loginpage")
}

func logout(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"logoutpage")
}