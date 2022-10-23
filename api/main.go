package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
  
  "github.com/rs/cors"
	"github.com/bitly/go-simplejson"
	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Book struct {
	// ID string `json:""`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

type allBooks []Book

var books = allBooks{}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome home!")
}

func helloWorldF(w http.ResponseWriter, r *http.Request) {
	json := simplejson.New()
	json.Set("msg", "Hello world, guy!")

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func createBookLink(w http.ResponseWriter, r *http.Request) {
  log.Println("hit da api")
  // declare new book
  var newBook Book

  // try to decode the request body into the struct. If there's an error, 
  //  respond to the client with the error message and a 400 status code
  err := json.NewDecoder(r.Body).Decode(&newBook)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }

  //
  fmt.Fprintf(w, "Book: %+v", newBook)

  books = append(books, newBook)

  w.Header().Set("Content-Type", "application/json")
	payload, err := json.Marshal(newBook)
  if err != nil {
    log.Println(err)
  }

  w.Write(payload)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func main() {
	router := mux.NewRouter()//.StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/helloWorld", helloWorldF)
  router.HandleFunc("/create-book", createBookLink).Methods("POST")
	router.HandleFunc("/get-books", getBooks)

  c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:5173"},
    AllowCredentials: true,
  })

  handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
