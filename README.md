# InstaLike

<p>
Go-based REST API webservice for an Instagram-like app, thus InstaLike :-)
</p>

## 📝 Table of Contents

- [About](#about)
- [Quick start](#quickstart)
- [Features](#features)
- [Configuration](#config)
- [API](#api)

## About <a name = "about"></a>

This is a simple web service provides a simple way managing user posts.
Is is Go web service written in Go that uses Postgres for data storage.


## Quick start <a name = "quickstart"></a>

## Docker world

Run fully functioning service with the database in docker containers using docker compose (assuming docker and docker compose are installed):
```
docker-compose up --build
```
In your browser go to http://localhost:3000/swagger/index.html


### Developer world
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
todo: apply db/init.sql

Start web service:
```
  make start
```

Build:
```
  make build
```
This will build the code and put it in ./bin directory along with required supporting files (configuration and RSA keys for JWT)

## Features <a name = "features"></a>

- REST standard (followed as much as possible)
- New users registration and logging them in uses a simple username/password pattern.
- Accessing posts and other resources are protected by JWT.  RSA public key scheme used in token creation and verification.
- For simplicity and given time constraints post’s multimedia attachments (image and video content) are stored in the database.  TODO: This needs to be changed to store the binary data in the file system. 

## Approach and design

### Tech
- go …duh!
- go fiber – very fast API framework build on top Fasthttp
- go playground validator – user input validation
- go playgroud jwt – for token creation and validation
- brcypt – for password hashing
- pgx– a postgres driver
- sqlx – is a library which provides a set of extensions on go's standard database/sql library.

### Code


## Configuration <a name = "config"></a>

This web service uses [node-config-ts](https://www.npmjs.com/package/node-config-ts) for configuration.
The config file is 
```
./config/default.json

{
  "port": 4040,
  "mongoUri": "mongodb://localhost:27017/acronyms",
  "seedFile": "./data/acronym.json"
}
```
You may change the values in that file. Also you can have settings applied based on a runtime environment.
Example:
set environment variable 
```
  NODE_ENV=production
```
then the service will use setting from *./config/env/production.json* config file.
```
  ./config/env/production.json
```
You can also use command line arguments to override settings.
```
  --port 5000
```

For full understanding of configuration features please refer to [node-config-ts](https://www.npmjs.com/package/node-config-ts)  documentation


## API <a name="api"></a>

### Get acronyms by *fuzzy* searching definitions
```
  GET /acronym?from=:fromlimit=:limit&search=:search
```
Query arguments:
- *:search* is a mandatory string what to search for (case insensitive)
- *:from* is a mandatory non-negative number for paging: how many results to skip
- *:limit* is a mandatory non-negative number for paging: maximum number of acronyms to return

Result is a JSON array of acronyms with definitions
```
  JSON:
  [
    {
      acronym: string
      definition: string
    },
    ...
  ]
```
Paging note: after getting a page of results if more data available the response header will contain an entry "next" with a path to the next page of results.
```
  GET /acronym?from=10limit=5&search=freack
  
  Response header (when more data available):
    next: /acronym?from=15limit=5&search=freack
```
- On success returns: HTTP Status 200 (OK)
- In case of missing or invalid query parameters returns: HTTP Status 400 (BAD_REQUEST)

Testing
```
curl --location --request GET 'http://localhost:4040/acronym?from=50&limit=10&search=one'
```


### Get acronym's definitions
```
  GET /acronym/:acronym
```
Query arguments:
- *:acronym* a mandatory string of an acronym (case insensitive)

Result is a JSON object of an acronym and its definition
```
  JSON:
    {
      acronym: string
      definition: string
    }
```
*Since the client alreany has the acronym, an agument could be made for retuning just the acronym's definition.  For consitency I decided returning the "whole" object across all GET verbs.*
- On success returns: HTTP Status 200 (OK)
- When acronym does not exist returns: HTTP Status 404 (NOT_FOUND)

Testing
```
curl --location --request GET 'http://localhost:4040/acronym/test99'
```


### Create new acronym with definition
```
  POST /acronym
  Header:
    Content-Type: application/json
  Body:
  {
    acronym: string
    definition: string
  }
```
- On success returns: HTTP Status 201 (CREATED)
- If acronym already exist returns: HTTP Status 409 (CONFLICT)
- If either acronym or definition is missing returns: HTTP Status 400 (BAD_REQUEST)

Testing
```
curl --location --request POST 'http://localhost:4040/acronym' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "acronym": "TEST99",
    "definition": "Test ninety nine"
}'
```


### Update acronym's definition</i>
```
  PUT /acronym/:acronym
  Header:
    Authorization: XXXXX
    Content-Type: application/json
  Body: {
    definition: string
  }
```
- On success returns: HTTP Status 204 (NO_CONTENT)
- On failure returns: HTTP Status 400 (BAD_REQUEST)
- On missing authorization header returns: HTTP Status 400 (UNAUTHORIZED)

Note: *This API uses an authorization header to ensure acronyms are protected.  Currently this implementation just checks for the presense of Authorization header. It does not validate the token.*

Testing
```
curl --location --request PUT 'http://localhost:4040/acronym/test99' \
  --header 'Content-Type: application/json' \
  --header 'Authorization: auth-token' \
  --data-raw '{
    "definition": "Test ninety nine. Take 2"
}'
```


### Delete an acronym
```
  GET /acronym/:acronym
  Header:
    Authorization: XXXXX
```
Note: *This API uses an authorization header to ensure acronyms are protected.  Currently this implementation just checks for the presense of Authorization header. It does not validate the token.*

- On success returns: HTTP Status 204 (NO_CONTENT)
- On failure returns: HTTP Status 400 (BAD_REQUEST)
- On missing authorization header returns: HTTP Status 400 (UNAUTHORIZED)

Testing
```
curl --location --request DELETE 'http://localhost:4040/acronym/test99' \
  --header 'Authorization: auth-token'
```

## �TODO <a name = "todo"></a>
- logging
- unit testing
