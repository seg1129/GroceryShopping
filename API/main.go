package main

import (
  "fmt"
  "log"
  "net/http"
  "database/sql"
  "encoding/json"
  _"github.com/lib/pq"
)

type Recipe struct {
  Id int
  Name string
  Ingredient1 string
  Ingredient2 string
  Ingredient3 string
  Ingredient4 string
  Ingredient5 string
  Ingredient6 string
  Ingredient7 string
  Ingredient8 string
  Ingredient9 string
  Ingredient10 string
  Instructions string
  Estimated_time string
  Source string

}
func establishDatabaseConnection() *sql.DB{
  // check configuration of database
  db, err := sql.Open("postgres","user=postgres dbname=grocerylist sslmode=disable")
  if err != nil {
    log.Fatal("Error: The data source arguments are not valid")
  }
  // check connection to databaseerr = db.Ping()
  if err != nil {
    log.Fatal("Error: Could not establish a connection with the database")
  }
  return db
}
func allRecipes(w http.ResponseWriter, r *http.Request){

  db := establishDatabaseConnection()

  rows, err := db.Query(`SELECT id, name, ingredient1 FROM recipe_simple`)
  if err != nil {
      panic(err)
  }

  defer rows.Close()
  var recipes []Recipe
  for rows.Next() {
      recipe := Recipe{}
      err = rows.Scan(&recipe.Id, &recipe.Name, &recipe.Ingredient1)
      if err != nil {
          panic(err)
      }
      recipes = append(recipes, recipe)
  }
  err = rows.Err()
  if err != nil {
      panic(err)
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(&recipes)
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "testing homepage")
}

func handleRequests(){
    http.HandleFunc("/", homePage)
    http.HandleFunc("/allRecipes", allRecipes)
    log.Fatal(http.ListenAndServe(":8080", nil))


}

func main(){
  handleRequests()
}
