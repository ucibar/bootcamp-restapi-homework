# Book Catalogue, Order RESTful API
## Install
```
$ go build ./cmd/http/
$ ./http
```

## Implemented Endpoints
- `GET /authors` get details of all authors
- `POST /authors` add new author:

```
{
    "name": "Author Name"
}
```

- `GET /authors/:id` get details of an author
- `PUT /authors/:id` update author:

```
{
    "name": "Author Name"
}
```

- `DELETE /authors/:id` delete author
- `DELETE /authors/:id` delete author
- `GET /authors/:id/books` get books of author


- `GET /books` get details of all books. 
- `GET /books?authors=1,2,3` get books with by authors
- `GET /books?price=":price"` get books with price equal to <price>
- `GET /books?price="< :price"` get books with price lower than <price>
- `GET /books?price="> :price"` get books with price higher than <price>
- `POST /books` add new book:

```
{
    "name": "Book Name",
    "price": 12.5,
    "author_id": 1
}
```

- `GET /books/:id` get details of a book
- `PUT /books/:id` update book:

```
{
    "name": "Book Name",
    "price": 12.5,
    "author_id": 2
}
```

- `DELETE /books/:id` delete book
- `GET /book/:id/author` get author of book


- `GET /orders` get all orders
- `POST /orders` new order:
```
{
    "items": [
        {
            "book_id": 1,
            "quantity": 1
        },
        {
            "book_id": 2,
            "quantity": 4
        }
    ]
}
```