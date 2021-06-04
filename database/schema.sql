CREATE DATABASE IF NOT EXISTS book_store;

USE book_store;

CREATE TABLE IF NOT EXISTS users (
  id int NOT NULL AUTO_INCREMENT,
  firstName VARCHAR(100) NOT NULL,
  lastName VARCHAR(100),
  email VARCHAR(100) NOT NULL,
  password VARCHAR(100),
  birth DATETIME,
  address VARCHAR(200),
  isAdmin BOOLEAN,
  createdAt DATETIME,
  updatedAt DATETIME,
  PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS books (
  id int NOT NULL AUTO_INCREMENT,
  title VARCHAR(100) NOT NULL,
  description VARCHAR(200),
  author VARCHAR(100) NOT NULL,
  publisher VARCHAR(100),
  publicationYear int,
  createdAt DATETIME,
  updatedAt DATETIME,
  PRIMARY KEY(id)
);
