# todo-api

A simple API for task management (to-do list).

## Local development

```
docker-compose up
```

## Project architecture

	.
	├── config                  # Structs for initializing global configuration
	├── http                    # Layer that exposes application to HTTP protocol
	├── k8s                     # Deployment source files
	├── migrations              # SQL files for Postgres migrations
	├── models                  # Communicates with persistence layer using repository pattern
	├── password                # Crypto utilities
	├── services                # Thin layer that composes application use cases
	├── storage                 # Interface for persist objects in different cloud providers
	├── main.go                 # Constructs container for starting project with its HTTP server

## Unauthenticated endpoints

### Create a user

POST http://localhost:8000/v1/users
```json
{
	"firstname": "André",
	"lastname": "Bazaglia",
	"email": "andre@example.com",
	"password": "123"
}
```

### Login as a user

POST http://localhost:8000/v1/login
```json
{
	"email": "andre@example.com",
	"password": "123"
}
```

## Authenticated endpoints

### Create a task

POST http://localhost:8000/v1/tasks
```json
{
	"name": "My first task",
	"description": "I have to do this and that on my first task",
	"location": "Amsterdam - NL",
	"date": "2019-10-10",
	"labels": ["urgent", "personal"],
	"comments": ["Close this task as soon as possible", "Try it out!"],
	"is_favorite": true
}
```

### List a task

GET http://localhost:8000/v1/tasks