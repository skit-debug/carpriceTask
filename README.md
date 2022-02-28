# carpriceTestTask

To run the app use:
> docker-compose up -d --build    

Be shure that ports 8080(app) and 3306(mysql) are free.

Description is in *swagger.yaml*.

Some browser commands to get data:

- localhost:8080/authors/
- localhost:8080/books/?page=3
- localhost:8080/authors/search/?last_name="Pelevin"
- localhost:8080/books/search/?release_year=1999
- localhost:8080/books/search/Rowling?page=2

Some curl commands to add/update/delete data:

- curl -d "{\"first_name\": \"newFirstName\", \"last_name\": \"newLastName\", \"born\": \"1900-01-01\", \"died\": \"2000-01-01\"}" -H "Content-Type: application/json" -X POST localhost:8080/authors/
- curl -d "{\"title\": \"newBook\", \"release_year\": 2022, \"abstract\": \"some abstract\", \"author_id\": 5}" -H "Content-Type: application/json" -X POST localhost:8080/books/

- curl -d "{\"author_id\": 5, \"first_name\": \"updatedFirstName\", \"last_name\": \"updatedLastName\", \"born\": \"2000-01-01\", \"died\": \"3000-01-01\"}" -H "Content-Type: application/json" -X PUT localhost:8080/authors/
- curl -d "{\"book_id\": 13,\"title\": \"updatedBook\", \"release_year\": 3022, \"abstract\": \"some updated abstract\", \"author_id\": 4}" -H "Content-Type: application/json" -X PUT localhost:8080/books/

- curl -X DELETE localhost:8080/books/13/
- curl -X DELETE localhost:8080/authors/5/

On Windows you must use escape character near " in your curl requests.
