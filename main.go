package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// ResponseWriter and Request to accept the client request,
// and create a http response to send
func helloHandler(w http.ResponseWriter, r *http.Request) {

	// check if the path is right
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// check the right method, for hello, it should not be a POST method
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	//print the message
	fmt.Fprintf(w, "Hello there, Welcome !!!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	// parse the incoming form and check for error
	err := r.ParseForm()
	checkError(err)

	fmt.Fprintf(w, "POST request successful\n")

	//extract the values from the submitted form
	// gives the values of the specified strings from the form
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name : %v\n", name)
	fmt.Fprintf(w, "Address : %v\n", address)
}

func resumeHandler(w http.ResponseWriter, r *http.Request) {

	// check if the path is right
	if r.URL.Path != "/resume" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// check the right method, for hello, it should not be a POST method
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	//read the resume file and convert it into bytes to send
	htmlBytes, err := ioutil.ReadFile("./static/resume.html")
	checkError(err)

	// set the content type header to html
	w.Header().Set("Content-Type", "text/html")

	//write the html content to the response writer
	fmt.Fprint(w, string(htmlBytes))
}

func main() {
	// create a file server for the files
	// look into the given directory for the files
	fileserver := http.FileServer(http.Dir("./static"))

	// will server the index.html file
	http.Handle("/", fileserver)

	//handle other functions
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/resume", resumeHandler)

	fmt.Println("Starting server at port 8080...")

	// creates a server and listens at the given port
	err := http.ListenAndServe(":8080", nil)
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
