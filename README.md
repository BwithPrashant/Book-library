# Book-library
This is a golang project which exposes Rest API's for book management in a library

# Rest Endpoints

- Add book
  endpoint : /books
  method : POST

- Modify book
  endpoint : /books/{id}
  method : PUT


- Delete book
  endpoint : /books/{id}
  method : DELETE
  
 
- Get book by ID
  endpoint : /books/{id}
  method : PUT
  
  
- GetAll book
  endpoint : /books
  method : PUT

- Get books by filter
  endpoint : /books?filter1=param1,param2&filter2=param3
  method: GET
  
 # DB
   - DB name : Postgres
 # DBSchema
 
 ```javascript
postgres=# \d book_library.*
           Table "book_library.books"
 Column  | Type | Collation | Nullable | Default 
---------+------+-----------+----------+---------
 id      | uuid |           | not null | 
 isbn    | text |           |          | 
 title   | text |           |          | 
 author  | text |           |          | 
 country | text |           |          | 
Indexes:
    "books_pkey" PRIMARY KEY, btree (id)
    "books_isbn_key" UNIQUE CONSTRAINT, btree (isbn)

Index "book_library.books_isbn_key"
 Column | Type | Definition 
--------+------+------------
 isbn   | text | isbn
unique, btree, for table "book_library.books"

Index "book_library.books_pkey"
 Column | Type | Definition 
--------+------+------------
 id     | uuid | id
primary key, btree, for table "book_library.books"
````
