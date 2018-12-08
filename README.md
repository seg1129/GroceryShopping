# Grocery List App
Not ready to be graded - Setup is not working properly

For CS 575 fall quarter

## Features
  * Ability to see all recipes available in the web browser
  * Ability to filter all recipes in the web browser
  * Ability to generate shopping list with all ingredients in first three recipes available in web browser

## Local setup
There are Three components to this application: UserInterface (Angular), API (Golang), Database
Everything is running on default local ports and in current state, these are not parameritized so here is the port information for each component:
  * UserInterface: localhost:4200
  * API: localhost:8080
  * Postgres port: 5432

### Requirements
  * Go is installed and GOPATH is set up
  * NPM is installed and set up
  * Postgres is installed and set up locally using port 5432

### Setup database
  * Create database: `createdb -T template0 grocerylist`
  * `psql grocerylist < groceryList.db`

### Setup API
  * `go get github.com/lib/pq`
  * go to the file path of main.go in API and use the following command: `go run main.go`

## Setup UserInterface
  * go into UserInterface and use the following command `npm start`
