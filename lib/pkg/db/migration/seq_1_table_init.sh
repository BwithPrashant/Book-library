
#!/bin/sh

POSTGRES_HOST="localhost"
POSTGRES_PORT=5432
POSTGRES_USER="postgres"
POSTGRES_PASSWORD="password"
POSTGRES_DBNAME="postgres"

psql "host=$POSTGRES_HOST port=$POSTGRES_PORT dbname=$POSTGRES_DBNAME user=$POSTGRES_USER password=$POSTGRES_PASSWORD" <<EOF

CREATE SCHEMA IF NOT EXISTS book_library;

CREATE TABLE IF NOT EXISTS book_library.books (id uuid PRIMARY KEY, isbn INT, title TEXT, author TEXT, country TEXT);

EOF

