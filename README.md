# Go REST API Boilerplate

A basic project setup for Go REST APIs using Gin router, Postgres & Redis. This app currently handles sessions/authentication using Redis & gin middleware.

## Endpoints 

```
POST /v1/auth/signup - Creates a new user and creates a new session
POST /v1/auth/login - Logs in a user and creates a new session 
POST /v1/auth/logout - Logs out user and deletes their existing session
** Requires a user to be authenticated **
GET /v1/users/self - Gets the user's info using their existing session id (curl -v http://localhost:9000/v1/users/self --cookie "sid=my_session_id")
GET /v1/users/:id - Gets the user's info using the provided query param ID
```

## Setup steps

1. Create a new `.env` file and copy the placeholder values from `.env.example` to this new file. You will need to manually create your new database when starting your new project and add real values to your `.env` in order for the app to work as expected.
2. In the `go.mod` file update the module name, you will also need to update this in any other files that used the default module name
3. Start your Docker containers using `make up`
4. Start your app using `make run` &mdash; et voilà!