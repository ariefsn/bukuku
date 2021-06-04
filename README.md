# BUKU KU

> Simple book store API

## Requirements

  1. Go
  2. Docker

### Install

  1. Clone

      ```bash
        git clone https://github.com/ariefsn/gom.git
      ```

  2. Docker Compose

      ```bash
        docker-compose up
      ```

### Endpoints

  1. Auth

      | Method      | Bearer    | Endpoint  | Payload   |
      |-------------|-----------|-----------|-----------|
      | POST        | No        | [/auth/register](http://localhost:3001/auth/register) | [User Model](#models) |
      | POST        | No        | [/auth/token](http://localhost:3001/auth/token) | [User Model](#models) |
      | GET         | Yes       | [/auth/me](http://localhost:3001/auth/me) | -         |
      | PUT         | Yes       | [/auth/me](http://localhost:3001/auth/me) | [User Model](#models) |
      | POST        | Yes       | [/auth/user](http://localhost:3001/auth/user) | [User Model](#models) |
      | GET         | Yes       | [/auth/user](http://localhost:3001/auth/user) | -         |
      | GET         | Yes       | [/auth/user/:id](http://localhost:3001/auth/user/:id) | -         |
      | PUT         | Yes       | [/auth/user/:id](http://localhost:3001/auth/user/:id) | [User Model](#models) |
      | DELETE      | Yes       | [/auth/user/:id](http://localhost:3001/auth/user/:id) | - |

  2. Book

      | Method      | Bearer    | Endpoint  | Payload   |
      |-------------|-----------|-----------|-----------|
      | POST        | Yes       | [/book](http://localhost:3001/book) | [Book Model](#models) |
      | GET         | Yes       | [/book](http://localhost:3001/book) | -         |
      | GET         | Yes       | [/book/:id](http://localhost:3001/book/:id) | -         |
      | PUT         | Yes       | [/book/:id](http://localhost:3001/book/:id) | [Book Model](#models) |
      | DELETE      | Yes       | [/book/:id](http://localhost:3001/book/:id) | - |

### Models

- User

    ```json
      {
        "address": "Earth",
        "birth": "2050-01-10T00:00:00+07:00",
        "email": "john.doe@gmail.com",
        "firstName": "John",
        "lastName": "Doe",
        "isAdmin": false,
        "password": "Password.123",
      }
    ```

- Book

    ```json
      {
        "title": "Detective Conan",
        "description": "Excepteur ex ea ea non.",
        "author": "Gosho Aoyama",
        "publisher": "Elex Media Computindo",
        "publicationYear": 2012
      }
    ```
