# Grocery List App
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

### Video
[Application description YouTube Video](https://youtu.be/hO-WqYBmL8Y)
### Requirements
  * Go is installed and GOPATH is set up
  * NPM is installed and set up
  * you have node js version 8.9 or higher (this is needed to support angular V6)
  * Postgres is installed and set up locally using port 5432
  * make sure ports 8080 and 4200 are not currently being used by another application

### Install Application
  * clone this repo into the GOPATH file path `git clone https://github.com/seg1129/GroceryShopping.git`


### Setup database
  * Create database: `createdb -T template0 grocerylist`
  * `psql grocerylist < groceryList.db`

### Setup API
  * `go get github.com/lib/pq`
  * go to the file path of main.go in API and use the following command: `go run main.go`

## Setup UserInterface
### Set up Locally
  * go into UserInterface and use the following command to install all required packages `npm install`
  * Start application `npm start`

### Set up using Docker
The following steps were taken from [here](https://mherman.org/blog/dockerizing-an-angular-app/)

  * go to the userInterface folder and run the following command to build the docker image. `docker build -t Grocery_List_app_ui .`
  * Run the docker container.
  `docker run -it \
  -v ${PWD}:/usr/src/app \
  -v /usr/src/app/node_modules \
  -p 4200:4200 \
  --rm \
  Grocery_List_app_ui`
