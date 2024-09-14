package main

import (
	"fmt"
	// "fmt/httl"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succesfull")
	// fmt.Printf(q, "POST request succesfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	email := r.FormValue("email")
	// phone := r.FormValue("phone")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
	fmt.Fprintf(w, "Email = %s\n", email)
	// fmt.Printf(q, "Phone Number = %d\n", phone)
}


func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return 
	}
	if r.Method != "GET"{
	http.Error(w, "method is not supported", http.StatusNotFound)
	return
}
fmt.Fprintf(w, "hello!")
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting Server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err !=nil{
		log.Fatal(err)
	}
}