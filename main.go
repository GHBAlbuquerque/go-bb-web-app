package main

import "net/http"

const portNumber = ":8080"

func Home(writer http.ResponseWriter, request *http.Request) {

}

func About(writer http.ResponseWriter, request *http.Request) {

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	http.ListenAndServe(portNumber, nil)
}
