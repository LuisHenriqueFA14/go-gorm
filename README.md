<h1 align="center">go-gorm</h1>
<p align="center">Complete API made in golang.</p>

## Project

I made this project, because i love golang, and i wanted to try to do things that i know in NodeJS, but in golang.

Then i had the idea to make a full API, with database, authentication, and 6 different routes.

## Stack

This project uses:
- `go-gorm`: Golang ORM
- `jwt`: Json Web Token (authentication)
- `bcrypt`: Password encryption
- `uuid`: Professional IDs

## Database

User model has these fields:

- `ID`:           uuid    - unique
- `Name`:         string  - unique
- `Description`:  string
- `Email`:        string  - unique
- `Password`:     string

## API

POST Routes:
- `/users/register`: Register a new user
- `/users/login`: Authenticate a user (return JWT Token)
- `/users/update`: Update user info (require authentication)
- `/users/delete`: Delete user (require authentication)

GET Routes:
- `/users/list?limit={limit}`: List all users (default limit: 10)
- `/user/{name}`: List info about {user}

## How to use

After download, run this command to build the code:

```
chmod +x ./scripts/build.sh && ./scripts/build.sh
```

Then create you `.env` file. (use the model file `~/.env_model`)

Create the database folder:
```
mkdir database
```

Finally run the program:
```
./build/go-gorm
```
