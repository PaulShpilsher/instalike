# InstaLike

<p>
Go-based REST API webservice for an Instagram-like app, thus InstaLike :-)
</p>

## 📝 Table of Contents

- [About](#about)
- [Quick start](#quickstart)
- [Features](#features)
- [Design](#design)
- [Development](#development)
- [Configuration](#config)
- [API](#api)

## About <a name = "about"></a>

This is a simple web service provides a simple way managing user posts.
Is is Go web service written in Go that uses Postgres for data storage.


## Quick start <a name = "quickstart"></a>

Run fully functioning service with the database in docker containers using docker compose (assuming docker and docker compose are installed):
```
git clone https://github.com/PaulShpilsher/instalike.git
cd instalike
docker-compose up --build
```
In your browser go to http://localhost:3000/swagger/index.html


## Features <a name = "features"></a>

- REST standard (followed as much as possible)
- CORS
- New users registration and logging them in uses a simple username/password pattern.
- Accessing posts and other resources are protected by JWT.  RSA public key scheme used in token creation and verification.
- For simplicity and given time constraints post’s multimedia attachments (image and video content) are stored in the database.  TODO: This needs to be changed to store the binary data in the file system.
- Smallest possible docker image using fromscratch base image.


## Design

### Tech
- go …duh!
- go fiber - very fast API framework build on top Fasthttp
- go playground validator - structure validation
- go playgroud jwt - for token creation and validation
- brcypt - for password hashing
- swag - swagger
- pgx - a postgres driver
- sqlx - is a library which provides a set of extensions on go's standard database/sql library.

### Code
Project’s structure if fairly simple.  Here is run down of top level subdirectories:
- db: contains database initialization sql script files
- docs: contains swagger generated documentation files (to create them run command “make swag”
- keys: contains RSA public and private keys for token authentication (to create then run command “make cert”)
- pkg: contains all of the application logic

The application uses Domain Driven Design (DDD). ‘pkg/domain’ contains code for 3 domain bounded contexts. They are:
- users - contains everything related to user related use cases
- posts - contains  everything related to user posts related use cases.
- media - contains multimedia related code, such as downloading post images and videos

Each of the bounded contexts uses layered separation of responsibilities design pattern.  It uses dependency injection for layer interaction.  The layers are:
- repository - provides \database / store access
- service - provides domain’s business logic
- endpoint - provides API functionality, i.e. API handlers
- router - defines API routes

Also domain may contain definitions for domain’s data model, API DTO, and layer interface declarations


### Database

The model:

![diagram](https://github.com/PaulShpilsher/instalike/assets/20777554/cca01ccf-7c0c-41e7-a5b3-f21193c1d594)


An architectural decision was made not to delete any data.  Delete entity business use case is approached by having a deleted BOOLEAN column on “posts” and “post_comments” tables.  This allows to retain for historical and auditing purposes all the data.

Since the data returned to the user is filtered WHERE deleted IS FALSE we use filtered indexes to efficiently get the relevant data.

Another performance improvement is use of the VIEWS.  This allows us to save database access calls when for example we need to return post data with filled in author’s email (or in future username).  Without this we would have to query posts and users tables, and in code match user information to each of the posts read from the database.  

Important notes:  
- posts likes are stored in an array column containing ids of users who liked the post.  This is done in order to prevent user liking more than once provide the “unlike” functionality (not implement). And counting likes is provided for our pleasure by Postgres’s cardinality method (see posts_view)
- in future storing multimedia files needs to be removed from post_attachmens table.  This table will store only attachment’s metadata information and a reference (probably file id) to the file storage.


## Development <a name = "development"></a>
Assuming Go, docker, git and make installed, here are simple instructions to get you up and running in no time

Get the code from github
```
  git clone https://github.com/PaulShpilsher/instalike.git
  cd instalike
```

Start Postgres server instance in a docker container (*Docker must be installed*)
```
  make postgres
```

TODO: instructions on how to apply db/init.sql.  But for now just execute db/init.sql script in the database using you favorite Postgres client. (look .env file for database connection information)

Generate RSA keys (they will be stored in ./keys directory):
```
  make cert
```

Generate Swagger documents ( they will be stored in ./doc/instalike directory):
```
  make swag
```

Start web service:
```
  make start
```

Build:
```
  make build
```
This will build the code and put it in ./bin directory along with required supporting files (configuration and RSA keys for JWT)


## Configuration <a name = "config"></a>

This web service uses .env file for configuration.

example:
```
# ./.env

# Server configuration
HOST="0.0.0.0"
PORT=3000
DOMAIN="localhost"

CORS_ALLOWED_ORIGINS="*"

TOKEN_EXPIRATION_MINUTES=720
TOKEN_PRIVATE_KEY_FILE="keys/rsa"
TOKEN_PUBLIC_KEY_FILE="keys/rsa.pub"


# Database configuration
DB_URL="postgresql://pusr:pusr_secret@localhost:5432/instalike-data?sslmode=disable"
DB_MAX_CONNECTIONS=100
DB_MAX_IDLE_CONNECTIONS=10
DB_MAX_LIFETIME_CONNECTIONS=2
```

## API <a name="api"></a>

The APIs definitions are available with swagger.  Just start the server using docker compose and navigate to at http://localhost:3000/swagger/index.html

Summary of APIs:

Users:
  - POST /api/users/register
  - POST /api/users/login

 Posts:
  - GET /api/posts - get all posts
  - GET /api/posts/:postId - get a post by id
  - POST /api/posts - creates a post
  - PUT /api/posts/:postId - updates a post by id (only authors are allowed to update)
  - DELETE /api/posts/:postId - deletes a post by id (only authors are allowed to delete)
  - POST /api/posts/:postId/attachment - attaches a multimedia file to post (only authors are allowed to update)
  - POST /api/posts/:postId/like - likes a post (likes are limited to 1 per user)

Post comments:
  - POST/api/posts/{postId}/comments - creates a new post comment
  - GET /api/posts/{postId}/comments - gets a list of post's comments by id

Multimedia:
  - GET /media/attachments/{attachmentId} - downloads a attached multimedia file by id.  


## TODO <a name = "todo"></a>
- Tests
- File storage
