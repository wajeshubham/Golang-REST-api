# Simple GO Lang REST API

> Simple RESTful API to create, read, update and delete books. No database

## Required package

```bash
# Install mux router
go get -u github.com/gorilla/mux
```

## Run the project

```bash
go build
./restapi
# After running go build, new .exe file will be created in your project directory
# to run the server run ./<.exe file name>
```

## Endpoints

### Get All Books

```bash
GET /api/books
```

### Get Single Book

```bash
GET /api/book/{id}
```

### Delete Book

```bash
DELETE /api/books/delete/{id}
```

### Create Book

```bash
POST /api/books/create

# Request sample
# {
#   "id":"124134"
#   "serialno":"4545454",
#   "title":"Book two",
#   "author":{"firstname":"John",  "lastname":"Doe"}
# }
```

### Update Book

```bash
PUT api/books/update/{id}

# Request sample
# {
#   "id":"124134"
#   "serialno":"4545454",
#   "title":"Updated Book two",
#   "author":{"firstname":"Harry",  "lastname":"Potter"}
# }

```
