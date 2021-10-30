# toggl-cardgame

Project made for coding test from Toggl for create card deck APIs. List of APIs are as following:

* Create Deck - POST /v1/decks
* Open Deck - GET /v1/decks/:deck_id
* Draw Cards - GET /v1/decks/:deck_id/draw?count=

## Requirements
* Go 1.17
* Mysql server*

*Author was using Mysql 8.0.26 during the development of the project 


## Building Code

Inside the project there is a Makefile file which will do the following:
1. Initialize Database Objects
2. Run Unit Tests
3. Build Binary File

Inside the root directory, run `make` to build the code based on the Makefile configuration

You may face some issue when initializing the database due to authorization access to the database using root user. For such case, you may modify inside `cmd/migration/main.go` and change the credentials.

After build is successful, you can run the binary using `./cmd/main` command

Once running, you can validate by opening browser to url http://localhost:8000/

## API Testings
Api testings has been done using Postman tool. You can use the following Postman collection to run a series of test cases that will be calling the services and validating the reponses

**Postman Link**:
https://www.getpostman.com/collections/c8239993d058f50b6c3a

With the link above, you can import it to your local postman and run the collection right away. Make sure the service is already running
