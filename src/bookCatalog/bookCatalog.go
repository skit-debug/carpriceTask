package bookCatalog

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const pageSize int = 5

type BookCatalog struct {
	db       *sql.DB
	user     string
	password string
	dbName   string
}

type Authors struct {
	AuthorID  int    `json:"author_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Born      string `json:"born"`
	Died      string `json:"died"`
	BooksNo   int    `json:"books_no"`
}

type Books struct {
	BookID         int    `json:"book_id"`
	Title          string `json:"title"`
	ReleaseYear    int    `json:"release_year"`
	Abstract       string `json:"abstract"`
	AuthorID       int    `json:"author_id"`
	AuthorLastName string `json:"author_last_name"`
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

func (bc *BookCatalog) GetAllAuthors(page int) []Authors {
	var len int
	err := bc.db.QueryRow("SELECT COUNT(author_id) FROM authors").Scan(&len)
	if err != nil {
		log.Fatal(err)
	}
	maxPages := len / pageSize
	if len%pageSize != 0 {
		maxPages = maxPages + 1
	}
	if page > maxPages {
		return []Authors{}
	}
	rows, err := bc.db.Query("SELECT * FROM authors LIMIT ? OFFSET ?", pageSize, pageSize*(page-1))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var authorsArray []Authors

	for rows.Next() {
		row := Authors{}
		err = rows.Scan(&row.AuthorID, &row.FirstName, &row.LastName, &row.Born, &row.Died)
		if err != nil {
			row.Died = ""
		}
		var booksNo int
		err = bc.db.QueryRow("SELECT COUNT(book_id) FROM books WHERE author_id=?", row.AuthorID).Scan(&booksNo)
		if err != nil {
			log.Fatal(err)
		}
		row.BooksNo = booksNo

		authorsArray = append(authorsArray, row)
	}

	return authorsArray
}

func (bc *BookCatalog) SearchAuthors(condition string, page int) []Authors {
	var len int
	err := bc.db.QueryRow("SELECT COUNT(author_id) FROM authors WHERE " + condition).Scan(&len)
	if err != nil {
		log.Fatal(err)
	}
	maxPages := len / pageSize
	if len%pageSize != 0 {
		maxPages = maxPages + 1
	}
	if page > maxPages {
		return []Authors{}
	}

	rows, err := bc.db.Query("SELECT * FROM authors WHERE "+condition+" LIMIT ? OFFSET ?", pageSize, pageSize*(page-1))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var authorsArray []Authors

	for rows.Next() {
		row := Authors{}
		err = rows.Scan(&row.AuthorID, &row.FirstName, &row.LastName, &row.Born, &row.Died)
		if err != nil {
			row.Died = ""
		}
		var booksNo int
		err = bc.db.QueryRow("SELECT COUNT(book_id) FROM books WHERE author_id=?", row.AuthorID).Scan(&booksNo)
		if err != nil {
			log.Fatal(err)
		}
		row.BooksNo = booksNo

		authorsArray = append(authorsArray, row)
	}

	return authorsArray
}

func (bc *BookCatalog) CreateAuthor(a Authors) {
	_, err := bc.db.Exec("INSERT INTO authors VALUES (0, ?, ?, ?, ?)", a.FirstName, a.LastName, a.Born, a.Died)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BookCatalog) ChangeAuthor(a Authors) {
	_, err := bc.db.Exec("UPDATE authors SET first_name=?, last_name=?, born=?, died=? WHERE author_id=?", a.FirstName, a.LastName, a.Born, a.Died, a.AuthorID)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BookCatalog) DeleteAuthor(id int) error {
	_, err := bc.db.Exec("DELETE FROM authors WHERE author_id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (bc *BookCatalog) GetAllBooks(page int) []Books {
	var len int
	err := bc.db.QueryRow("SELECT COUNT(book_id) FROM books").Scan(&len)
	if err != nil {
		log.Fatal(err)
	}
	maxPages := len / pageSize
	if len%pageSize != 0 {
		maxPages = maxPages + 1
	}
	if page > maxPages {
		return []Books{}
	}

	rows, err := bc.db.Query("SELECT * FROM books LIMIT ? OFFSET ?", pageSize, pageSize*(page-1))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var booksArray []Books

	for rows.Next() {
		row := Books{}
		err = rows.Scan(&row.BookID, &row.Title, &row.ReleaseYear, &row.Abstract, &row.AuthorID)
		if err != nil {
			log.Fatal(err)
		}
		var AuthorLastName string
		err = bc.db.QueryRow("SELECT last_name FROM authors WHERE author_id=?", row.AuthorID).Scan(&AuthorLastName)
		if err != nil {
			log.Fatal(err)
		}
		row.AuthorLastName = AuthorLastName

		booksArray = append(booksArray, row)
	}

	return booksArray
}

func (bc *BookCatalog) CreateBook(b Books) {
	_, err := bc.db.Exec("INSERT INTO books VALUES (0, ?, ?, ?, ?)", b.Title, b.ReleaseYear, b.Abstract, b.AuthorID)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BookCatalog) ChangeBook(b Books) {
	_, err := bc.db.Exec("UPDATE books SET title=?, release_year=?, abstract=?, author_id=? WHERE book_id=?", b.Title, b.ReleaseYear, b.Abstract, b.AuthorID, b.BookID)
	if err != nil {
		log.Fatal(err)
	}
}

func (bc *BookCatalog) DeleteBook(id int) error {
	_, err := bc.db.Exec("DELETE FROM books WHERE book_id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func (bc *BookCatalog) SearchBooks(condition string, page int) []Books {
	var len int
	err := bc.db.QueryRow("SELECT COUNT(book_id) FROM books WHERE " + condition).Scan(&len)
	if err != nil {
		log.Fatal(err)
	}
	maxPages := len / pageSize
	if len%pageSize != 0 {
		maxPages = maxPages + 1
	}
	if page > maxPages {
		return []Books{}
	}

	rows, err := bc.db.Query("SELECT * FROM books WHERE "+condition+" LIMIT ? OFFSET ?", pageSize, pageSize*(page-1))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var booksArray []Books

	for rows.Next() {
		row := Books{}
		err = rows.Scan(&row.BookID, &row.Title, &row.ReleaseYear, &row.Abstract, &row.AuthorID)
		if err != nil {
			log.Fatal(err)
		}
		var AuthorLastName string
		err = bc.db.QueryRow("SELECT last_name FROM authors WHERE author_id=?", row.AuthorID).Scan(&AuthorLastName)
		if err != nil {
			log.Fatal(err)
		}
		row.AuthorLastName = AuthorLastName

		booksArray = append(booksArray, row)
	}

	return booksArray
}

func (bc *BookCatalog) SearchBooksByAuthor(author string, page int) []Books {
	rows, err := bc.db.Query("SELECT * FROM books WHERE author_id IN (SELECT author_id from authors WHERE last_name=\"" + author + "\")")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var booksArray []Books

	i := -1

	for rows.Next() {
		i = i + 1
		if !((i >= (page-1)*pageSize) && (i < page*pageSize)) {
			continue
		}

		row := Books{}
		err = rows.Scan(&row.BookID, &row.Title, &row.ReleaseYear, &row.Abstract, &row.AuthorID)
		if err != nil {
			log.Fatal(err)
		}
		var AuthorLastName string
		err = bc.db.QueryRow("SELECT last_name FROM authors WHERE author_id=?", row.AuthorID).Scan(&AuthorLastName)
		if err != nil {
			log.Fatal(err)
		}
		row.AuthorLastName = AuthorLastName

		booksArray = append(booksArray, row)
	}

	return booksArray
}
