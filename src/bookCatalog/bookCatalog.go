package bookCatalog

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BookCatalog struct {
	db       *sql.DB
	user     string
	password string
	dbName   string
}

func (bc *BookCatalog) OpenCatalog(ctx context.Context, user string, password string, dbName string) {
	bc.user = user
	bc.password = password
	bc.dbName = dbName
	var err error
	bc.db, err = sql.Open("mysql", user+":"+password+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
	if err := bc.db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}
}

func (bc *BookCatalog) CloseCatalog() {
	bc.db.Close()
}

func (bc *BookCatalog) Get() {
	row := bc.db.QueryRow("SELECT last_name FROM authors WHERE author_id=1")
	var a string
	row.Scan(&a)
	/*
		if err != nil {
		    if err == sql.ErrNoRows {
		        fmt.Println("Zero rows found")
		    } else {
		        panic(err)
		    }
		}*/
}

type Authors struct {
	AuthorID  int    `json:"author_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	//Born      time.Date `json:"born"`
	//Died      time.Date `json:"died"`
}

type Books struct {
	BookID      int    `json:"book_id"`
	Title       string `json:"title"`
	ReleaseYear int    `json:"release_year"`
	Abstract    string `json:"abstract"`
	AuthorID    int    `json:"author_id"`
}
