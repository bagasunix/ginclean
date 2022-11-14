![Clean Template](docs/logo.svg)

# Gin Clean template

## Technologies - Libraries

- ✔️ **[`gin-gonic/gin`](https://github.com/gin-gonic/gin)** -Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster.
- ✔️ **[`go-gorm/gorm`](https://github.com/go-gorm/gorm)** - The fantastic ORM library for Go, aims to be developer friendly
- ✔️ **[`uber-go/zap`](https://github.com/uber-go/zap)** - Blazing fast, structured, leveled logging in Go.
- ✔️ **[`joho/godotenv`](https://github.com/joho/godotenv)**
- ✔️ **[`go-ozzo/ozzo-validation`](https://github.com/go-ozzo/ozzo-validation)** - Ozzo-validation is a Go package that provides configurable and extensible data validation capabilities.
- ✔️ **[`gofrs/uuid`](https://github.com/gofrs/uuid)** - Package uuid provides a pure Go implementation of Universally Unique Identifiers (UUID) variant as defined in RFC-4122. This package supports both the creation and parsing of UUIDs in different formats.

# API Group List

## API Role `/v1/role`

| Route                 | HTTP   | Description                        |
| --------------------- | ------ | ---------------------------------- |
| /                     | POST   | Route used to create role          |
| /                     | GET    | Route used to list role            |
| /:id                  | GET    | Route used to view role            |
| /:id                  | PATCH  | Route used to update role          |
| /update-multiple-role | PATCH  | Route used to update multiple role |
| /:id                  | DELETE | Route used to delete role          |

## API Role `/v1/user`

| Route | HTTP   | Description                   |
| ----- | ------ | ----------------------------- |
| /     | POST   | Route used to create account  |
| /     | GET    | Route used to list account    |
| /:id  | GET    | Route used to view account    |
| /:id  | DELETE | Route used to delete account  |
| /:id  | PUT    | Route used to disable account |

## API Role `/v1/auth`

| Route          | HTTP | Description                      |
| -------------- | ---- | -------------------------------- |
| /refresh-token | POST | Route used to refresh token      |
| /login         | POST | Route used to login account/user |
