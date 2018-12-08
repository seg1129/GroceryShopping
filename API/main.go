package main

import (
  "fmt"
  "log"
  "net/http"
  "database/sql"
  "encoding/json"
  // "bytes"
  // "io/ioutil"
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

type RecipeID struct {
  Id int
}
type Ingredients struct{
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
}
var(
  ingredient1 string
  ingredient2 string
  ingredient3 string
  ingredient4 string
  ingredient5 string
  ingredient6 string
  ingredient7 string
  ingredient8 string
  ingredient9 string
  ingredient10 string
)

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
func cleanArray(arr []string) []string{
  // var length int
  length := len(arr)
  index := 0
  for index < length {
		if arr[index] == "" {
			arr = append(arr[:index], arr[index+1:]...)
      length --
		} else {
      index ++
    }
	}
  return arr
}
// retreived method from https://flaviocopes.com/golang-enable-cors/
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
// retreived method from https://flaviocopes.com/golang-enable-cors/
func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
  (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
  (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func allRecipes(w http.ResponseWriter, r *http.Request){
  enableCors(&w)

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

func shoppingList(w http.ResponseWriter, r *http.Request){
  var allIngredients []string

  setupResponse(&w, r)
  db := establishDatabaseConnection()

  var recipeIds []RecipeID
  err := json.NewDecoder(r.Body).Decode(&recipeIds)
    if err != nil {
        fmt.Fprintf(w, "please enter recipe IDs")
    }

    // for each item in RecipeID - query database for ingredients
    for _, elem := range recipeIds {
      //This is a target for sql injection
      rows, err := db.Query(`SELECT ingredient1, ingredient2, ingredient3,
                            ingredient4, ingredient5, ingredient6, ingredient7,
                            ingredient8, ingredient9, ingredient10
                            FROM recipe_simple
                            WHERE id=$1`, elem.Id)

      defer rows.Close()
      for rows.Next() {

        err = rows.Scan(&ingredient1, &ingredient2,
                        &ingredient3, &ingredient4,
                        &ingredient5, &ingredient6,
                        &ingredient7, &ingredient8,
                        &ingredient9, &ingredient10)
        log.Println(ingredient1, ingredient2)
        if err != nil {
            panic(err)

        }
        allIngredients = append(allIngredients, ingredient1, ingredient2, ingredient3, ingredient4, ingredient5, ingredient6, ingredient7, ingredient8, ingredient9, ingredient10)
    }
  }
  allIngredients = cleanArray(allIngredients)
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(&allIngredients)

}


func homePage(w http.ResponseWriter, r *http.Request){
  enableCors(&w)
  fmt.Fprintf(w, "testing homepage")
}

func handleRequests(){
    http.HandleFunc("/", homePage)
    http.HandleFunc("/allRecipes", allRecipes)
    http.HandleFunc("/shoppingList", shoppingList)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main(){
  handleRequests()
}
