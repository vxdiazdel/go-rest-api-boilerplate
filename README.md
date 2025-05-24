# Go REST API Boilerplate

A basic project setup for Go REST APIs using Gin router, Postgres & Redis.

## Setup steps

1. Create a new `.env` file and copy the placeholder values from `.env.example` to this new file. You will need to manually create your new database when starting your new project in order for the app to work as expected.
2. In the `go.mod` file update the module name, you will also need to update this in any other files that used the default module name
3. Start your Docker containers using `make up`
4. Start your app using `make run` &mdash; et voilà!