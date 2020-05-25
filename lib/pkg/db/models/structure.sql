CREATE SCHEMA IF NOT EXISTS book_library;
CREATE TABLE IF NOT EXISTS book_library.books (id uuid PRIMARY KEY, isbn INT, title TEXT, author TEXT, country TEXT);