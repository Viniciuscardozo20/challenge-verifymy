# Challenge-VerifyMy

**Candidate:** Vinicius Cardozo

**LinkedIn:** [Vinicius Cardozo LinkedIn](https://www.linkedin.com/in/vinicius-cardozo-669a15136/)

## Requirements

* [Docker-compose](https://docs.docker.com/compose/install/)

## Setup

After cloning the repository, to run the project, execute the following command:

```bash
make up
```

By default, the API runs on port `8082`, and the database on port `27017`. However, these variables can be changed in the `.env` file.

## User Routes

### Create

Method: `POST`

    /api/v1/
    
#### Example

In the terminal, execute:

    curl -XPOST -H "Content-type: application/json" 
    -d '{
      "id": "1",
      "name": "Vinicius",
      "age": 27,
      "email": "vinicius@email.com",
      "password": "1234",
      "address": "street 01"
    }' 'http://localhost:8082/api/v1/''

Or, in any other interface of your choice, such as Postman or Insomnia.

**Body** (application/json):

    {
	    "id": "1",
      "name": "Vinicius",
      "age": 27,
      "email": "vinicius@email.com",
      "password": "1234",
      "address": "street 01"
    }

Expected Output

_HTTP Status_: `201 Created`
```json
{   
    "status": "success",
    "data": {
        "user": {
            "id": "1",
            "name": "Vinicius",
            "age": 27,
            "email": "vinicius@email.com",
            "password": "1234",
            "address": "street 01"
        }
    }
}
```

### Read

Method: `GET`

    /api/v1/:id

**Parameters**: 

* ID: User ID

#### Example

In the terminal, execute:

    curl -XGET -H 'http://localhost:8082/api/v1/:id'

Or, in any other interface of your choice, such as Postman or Insomnia.

Expected Output

_HTTP Status_: `200 OK`
```json
{   
    "status": "success",
    "data": {
        "user": {
            "id": "1",
            "name": "Vinicius",
            "age": 27,
            "email": "vinicius@email.com",
            "password": "1234",
            "address": "street 01"
        }
    }
}
```

### ReadAll

Method: `GET`

    /api/v1/

#### Example

In the terminal, execute:

    curl -XGET -H 'http://localhost:8082/api/v1/'

Or, in any other interface of your choice, such as Postman or Insomnia.

Expected Output

_HTTP Status_: `200 OK`
```json
{   
    "status": "success",
    "users": [
        {
          "id": "1",
          "name": "Vinicius",
          "age": 27,
          "email": "vinicius@email.com",
          "password": "1234",
          "address": "street 01"
        }
    ]
}
```

### Update

Method: `PUT`

    /api/v1/:id

**Parameters**: 

* ID: User ID 

#### Example

In the terminal, execute:

    curl -XPUT -H "Content-type: application/json" 
    -d '{
      "id": "1",
      "name": "Vinicius",
      "age": 27,
      "email": "vinicius@email.com",
      "password": "1234",
      "address": "street 01"
    }' 'http://localhost:8082/api/v1/:id''

Or, in any other interface of your choice, such as Postman or Insomnia.

**Body** (application/json):

    {
	    "id": "1",
      "name": "Vinicius",
      "age": 27,
      "email": "vinicius@email.com",
      "password": "1234",
      "address": "street 01"
    }

Expected Output

_HTTP Status_: `201 Created`
```json
{   
    "status": "success",
    "data": {
        "user": {
            "id": "1",
            "name": "Vinicius",
            "age": 27,
            "email": "vinicius@email.com",
            "password": "1234",
            "address": "street 01"
        }
    }
}
```

### Delete

Method: `DELETE`

    /v1/:id

**Parameters**: 

* ID: User ID 

#### Example

In the terminal, execute:

    curl -XDELETE -H 'http://localhost:8082/v1/:id''

Or, in any other interface of your choice, such as Postman or Insomnia.

Expected Output

_HTTP Status_: `204 OK`

## Run Tests
### Requirements

* Golang 1:20

To run the tests:

  ```bash
make test
```

The expected output will be:

    
    go test -race challenge-verifymy/app/api challenge-verifymy/app/handlers challenge-verifymy/cmd challenge-verifymy/config challenge-verifymy/core/models challenge-verifymy/core/ports challenge-verifymy/core/ports/testutil challenge-verifymy/core/services challenge-verifymy/customerror challenge-verifymy/infrastructure/mongodb -coverprofile=coverage.out
    ?       challenge-verifymy/app/api      [no test files]
    ?       challenge-verifymy/app/handlers [no test files]
    ?       challenge-verifymy/cmd  [no test files]
    ?       challenge-verifymy/config       [no test files]
    ?       challenge-verifymy/core/ports   [no test files]
    ?       challenge-verifymy/core/ports/testutil  [no test files]
    ?       challenge-verifymy/customerror  [no test files]
    ?       challenge-verifymy/infrastructure/mongodb       [no test files]
    ok      challenge-verifymy/core/models  0.032s  coverage: 90.0% of statements
    ok      challenge-verifymy/core/services        0.055s  coverage: 27.8% of statements