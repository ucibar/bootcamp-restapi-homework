# Book Catalogue, Order RESTful API

## Planned Endpoints
- `GET /authors` get details of all authors
- `POST /authors` add new author
- `GET /authors/:id` get details of an author
- `PUT /authors/:id` update author
- `DELETE /authors/:id` delete author
- `DELETE /authors/:id` delete author
- `GET /authors/:id/books` get books of author


- `GET /books` get details of all books. 
- `GET /books?authors=1,2,3` get books with by authors
- `GET /books?price=":price"` get books with price equal to <price>
- `GET /books?price="< :price"` get books with price lower than <price>
- `GET /books?price="> :price"` get books with price higher than <price>
- `POST /books` add new book
- `GET /books/:id` get details of a book
- `PUT /books/:id` update book
- `DELETE /books/:id` delete book
- `GET /book/:id/author` get author of book


- `GET /orders` get all orders
- `POST /orders` new order