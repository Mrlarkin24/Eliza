/*
*	Name: Peter Larkin (G00332346)
*	Data Represenation and Querying Project
*	Eliza Program
 */

package main

import (
	"fmt"
	"net/http"
	// import eliza package
	"./eliza"
)

func main() {

	//serve the files from the /Interface folder
	dir := http.Dir("./Interface")
	fileServer := http.FileServer(dir)
	//handle requests to /
	http.Handle("/", fileServer)
	//handle request to /chat
	http.HandleFunc("/chat", chatHandler)
	//listen on tcp and serve requests on port :8080
	http.ListenAndServe(":8080", nil)
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	// this is code that runs when a request is made to the /Chat resource.
	userInput := r.URL.Query().Get("user-input")
	reply := eliza.AskEliza(userInput)
	fmt.Fprintf(w, reply)
} //chatHandler
