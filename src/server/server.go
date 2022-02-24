package server

import (
	"context"
	"encoding/json"
	"log"
	"main/src/bookCatalog"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	port     string = ":8080"
	user     string = "root"
	password string = "password"
	dbName   string = "BOOK_CATALOG"
)

var catalog bookCatalog.BookCatalog

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func createAuthorHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("createAuthorHandler"))
}

func getAllAuthorsHandler(w http.ResponseWriter, req *http.Request) {
	page := req.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	p, _ := strconv.Atoi(page)
	renderJSON(w, catalog.GetAllAuthors(p))
}

func getAuthorHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("getAuthorHandler"))

}

func changeAuthorHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("changeAuthorHandler"))
}

func deleteAllAuthorsHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("deleteAllAuthorsHandler"))
}

func deleteAuthorHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("deleteAuthorHandler"))
}

func createBookHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("createBookHandler"))
}

func getAllBooksHandler(w http.ResponseWriter, req *http.Request) {
	page := req.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	p, _ := strconv.Atoi(page)
	renderJSON(w, catalog.GetAllBooks(p))
}

func getBookHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("getBookHandler"))
}

func changeBookHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("changeBookHandler"))
}

func deleteAllBooksHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("deleteAllBooksHandler"))
}

func deleteBookHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("deleteBookHandler"))
}

func StartServer(ctx context.Context) {
	go func() {
		select {
		case <-ctx.Done():
			catalog.CloseCatalog()
			return
		default:
		}
	}()

	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/authors/", createAuthorHandler).Methods("POST")
	router.HandleFunc("/authors/", getAllAuthorsHandler).Methods("GET")
	router.HandleFunc("/authors/{id:[0-9]+}/", getAuthorHandler).Methods("GET")
	router.HandleFunc("/authors/{id:[0-9]+}/", changeAuthorHandler).Methods("PUT")
	router.HandleFunc("/authors/", deleteAllAuthorsHandler).Methods("DELETE")
	router.HandleFunc("/authors/{id:[0-9]+}/", deleteAuthorHandler).Methods("DELETE")

	router.HandleFunc("/books/", createBookHandler).Methods("POST")
	router.HandleFunc("/books/", getAllBooksHandler).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}/", getBookHandler).Methods("GET")
	router.HandleFunc("/books/{id:[0-9]+}/", changeBookHandler).Methods("PUT")
	router.HandleFunc("/books/", deleteAllBooksHandler).Methods("DELETE")
	router.HandleFunc("/books/{id:[0-9]+}/", deleteBookHandler).Methods("DELETE")

	catalog.OpenCatalog(ctx, user, password, dbName)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
