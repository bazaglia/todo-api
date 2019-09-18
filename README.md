# todo-api

A simple API for creating task management (to-do list).

## Local development

```
docker-compose up
```

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