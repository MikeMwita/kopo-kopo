version: '3'

services:
  sqlite:
    image: nouchka/sqlite3
    volumes:
      - ./data:/data
      - ./schema.sql:/docker-entrypoint-initdb.d/01_schema.sql  # Mount the schema.sql file as an initialization script
    ports:
      - 8080:80
    command: sqlite3 /data/database.db