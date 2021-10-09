# Features

1. Password encryption using MD5 and AES
2. Thread safe server
3. Pagination with 2 results per page
4. Clean, well documented, resuable, and scalable code

# Endpoints

## **Create a User**

Returns the id of the newly created User.

- **URL**

  /users

- **Method:**

  `POST`

- **URL Params**

  None

- **Data Params**
  ```go
  {
  	ID int
  	Name string
  	Email string
  	Password string
  }
  ```
- **Success Response:**
  **Content:** `<latest id> int`

- **Error Response:**

  - **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "404 not found" }`
    **Create an User**

## **Find a User by ID**

Returns the details of a user in json format

- **URL**

  /users/:id

- **Method:**

  `GET`

- **URL Params**

  **Required:**

  `id=[integer]`

- **Data Params**

  None

- **Success Response:**
  **Content:**
  ```go
  {
   "ID": 1,
  	"Name": "Apoorva Srivastava",
  	"Email": "apoorvasrivastava.14@gmail.com",
  	"Password":"53+��3�蟼�ݎ5,�\n�~�7����޷\f�"
  }
  ```
- **Error Response:**

  - **Code:** 400 Bad Request <br />
    **Content:** `{ error : "Bad Request. Id missing" }`
    OR
  - **Code:** 400 Bad Request <br />
    **Content:** `{ error : "Bad Request. Id invalid" }`

## **Create a Post**

Returns the id of the newly created Post.

- **URL**

  /posts

- **Method:**

  `POST`

- **URL Params**

  None

- **Data Params**
  ```go
  {
  	ID int
  	Caption string
  	ImageURL string
  	CreatedAt timestamp
  	Password string
  }
  ```
- **Success Response:**
  **Content:** `<latest id> int`

- **Error Response:**

  - **Code:** 404 NOT FOUND <br />
    **Content:** `{ error : "404 not found" }`
    **Create an User**

## **Find a Post by ID**

Returns the details of a post in json format

- **URL**

  /posts/:id

- **Method:**

  `GET`

- **URL Params**

  **Required:**

  `id=[integer]`

- **Data Params**

  None

- **Success Response:**
  **Content:**
  ```go
  {
   "ID": 1,
  	"Caption": "Test Caption",
  	"ImageURL": "www.google.com",
  	"CreatedAt": "2021-10-08T23:58:26.188Z",
  	"UserId": 1 (id of the user who created the post)
  }
  ```
- **Error Response:**

  - **Code:** 400 Bad Request <br />
    **Content:** `{ error : "Bad Request. Id missing" }`
    OR
  - **Code:** 400 Bad Request <br />
    **Content:** `{ error : "Bad Request. Id invalid" }`

## **Find all Posts by a User by User ID**

Returns the details of all posts by a user in json format

- **URL**

  /posts/users/:id?page=:page_no

- **Method:**

  `GET`

- **URL Params**

  **Required:**

  `id=[integer]`

- **URL Queries**

  `?page=[integer]` Returns `2` results per page

- **Data Params**

  None

- **Success Response:**
  **Content:**
  ```go
  [
   {
  	    "_id": 3,
  	    "caption": "",
  	    "created_at": "0001-01-01T05:53:28+05:53",
  	    "image_url": "",
  	    "user_id": 8
   },
   {
  	    "_id": 4,
  	    "caption": "",
  	    "created_at": "0001-01-01T05:53:28+05:53",
  	    "image_url": "",
  	    "user_id": 8
   }
  ]
  ```
- **Error Response:**

  - **Code:** 400 Bad Request
  - **Content:** `{ error : "Bad Request. Id invalid" }`
