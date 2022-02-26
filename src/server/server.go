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
	a := bookCatalog.Authors{}
	err := json.NewDecoder(req.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	catalog.CreateAuthor(a)
	w.Write([]byte("createAuthorHandler: Done"))
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
	page := req.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	p, _ := strconv.Atoi(page)

	id := req.URL.Query().Get("id")
	firstName := req.URL.Query().Get("first_name")
	lastName := req.URL.Query().Get("last_name")

	if id != "" {
		renderJSON(w, catalog.SearchAuthors("author_id="+id, p))
		return
	}

	if firstName != "" {
		renderJSON(w, catalog.SearchAuthors("first_name="+firstName, p))
		return
	}

	if lastName != "" {
		renderJSON(w, catalog.SearchAuthors("last_name="+lastName, p))
		return
	}

	w.Write([]byte("Invalid search condition"))
}

func changeAuthorHandler(w http.ResponseWriter, req *http.Request) {
	a := bookCatalog.Authors{}
	err := json.NewDecoder(req.Body).Decode(&a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	catalog.ChangeAuthor(a)
	w.Write([]byte("changeAuthorHandler: Done"))
}

func deleteAuthorHandler(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	err := catalog.DeleteAuthor(id)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("deleteAuthorHandler: Done"))
}

func createBookHandler(w http.ResponseWriter, req *http.Request) {
	b := bookCatalog.Books{}
	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	catalog.CreateBook(b)
	w.Write([]byte("createBookHandler: Done"))
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
	page := req.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	p, _ := strconv.Atoi(page)

	id := req.URL.Query().Get("id")
	title := req.URL.Query().Get("title")
	releaseYear := req.URL.Query().Get("release_year")

	if id != "" {
		renderJSON(w, catalog.SearchBooks("book_id="+id, p))
		return
	}

	if title != "" {
		renderJSON(w, catalog.SearchBooks("title="+title, p))
		return
	}

	if releaseYear != "" {
		renderJSON(w, catalog.SearchBooks("release_year="+releaseYear, p))
		return
	}

	w.Write([]byte("Invalid search condition"))
}

func getBookByAuthorHandler(w http.ResponseWriter, req *http.Request) {
	page := req.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	p, _ := strconv.Atoi(page)

	author := mux.Vars(req)["author"]

	renderJSON(w, catalog.SearchBooksByAuthor(author, p))
}

func changeBookHandler(w http.ResponseWriter, req *http.Request) {
	b := bookCatalog.Books{}
	err := json.NewDecoder(req.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	catalog.ChangeBook(b)
	w.Write([]byte("changeBookHandler: Done"))
}

func deleteBookHandler(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(req)["id"])
	err := catalog.DeleteBook(id)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte("deleteBookHandler: Done"))
}

func test(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("111"))
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
	router.HandleFunc("/authors/search/", getAuthorHandler).Methods("GET")
	router.HandleFunc("/authors/", changeAuthorHandler).Methods("PUT")
	router.HandleFunc("/authors/{id:[0-9]+}/", deleteAuthorHandler).Methods("DELETE")

	router.HandleFunc("/books/", createBookHandler).Methods("POST")
	router.HandleFunc("/books/", getAllBooksHandler).Methods("GET")
	router.HandleFunc("/books/search/", getBookHandler).Methods("GET")
	router.HandleFunc("/books/", changeBookHandler).Methods("PUT")
	router.HandleFunc("/books/{id:[0-9]+}/", deleteBookHandler).Methods("DELETE")

	router.HandleFunc("/books/search/{author}/", getBookByAuthorHandler).Methods("GET")

	router.HandleFunc("/test/", test).Methods("GET")

	catalog.OpenCatalog(ctx, user, password, dbName)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
