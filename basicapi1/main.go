package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct{
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles[] Article

func allAticles(w http.ResponseWriter,r *http.Request){
	articles:=Articles{
		Article{Title:"Test Title",Desc:"Test Description",Content:"Test content"},
		Article{Title:"Test Title",Desc:"Test Description",Content:"Test content"},
	}
		fmt.Println("End point hit: All Article end")
		json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Homepage EndPOint Hit")
}
func ravi(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"Ravi Prakash")
}

func handleRequests(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/articles",allAticles)
	http.HandleFunc("/ravi",ravi)
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func main(){
	handleRequests()
}



