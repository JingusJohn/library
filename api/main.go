package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"log"
	"net/http"
  
  "github.com/google/uuid"
  "github.com/rs/cors"
	"github.com/bitly/go-simplejson"
	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Book struct {
	ID string `json:"id,omitempty"`
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
  w.Header().Set("Content-Type", "application/json")
  // declare new book
  var newBook Book

  _ = json.NewDecoder(r.Body).Decode(&newBook)
  newBook.ID = uuid.New().String()

  books = append(books, newBook)

  json.NewEncoder(w).Encode(newBook)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id := params["id"]

  for _, book := range books {
    if book.ID == id {
      json.NewEncoder(w).Encode(book)
    }
  }
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink).Methods("GET")
	router.HandleFunc("/helloWorld", helloWorldF).Methods("GET")
  router.HandleFunc("/create-book", createBookLink).Methods("POST")
	router.HandleFunc("/get-books", getBooks).Methods("GET")
  router.HandleFunc("/get-book/{id}", getBook).Methods("GET")

  c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:5173"},
    AllowCredentials: true,
  })

  handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
