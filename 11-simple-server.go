package main

import (
	"fmt"
	"net/http"
)

func homePageHandle(w http.ResponseWriter, r *http.Request){
	// Set the response header to indicate content type
	w.Header().Set("Content-Type", "text/plain")

	// response message
	fmt.Fprintln(w, "Welcome to the homepage")
}

func handleReurnStaticUserDetails(res http.ResponseWriter, req *http.Request){
	// create map
	user := make(map[string]string)

	// write to map
	user["name"]="alahira jeffrey"
	user["age"] = "27"
	user["mood"] = "normal"

	// Set the response header to indicate content type
	res.Header().Set("Content-Type", "application/json")

	// response message
	fmt.Fprintln(res, user)
}



func eleventhMain() {
    // Register the handler function to respond to all requests on the root path "/"
    http.HandleFunc("/", homePageHandle)
	http.HandleFunc("/return-static-user-details", handleReurnStaticUserDetails)

    // Start the HTTP server on port 8080
    err := http.ListenAndServe(":3000", nil)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    }
}