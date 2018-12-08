#API Documentation

## GET all recipes
### Request
`GET /allRecipes HTTP/1.1
Host: localhost:8080`

### Example Response
`[
    {
        "Id": 1,
        "Name": "whole chicken",
        "Ingredient1": "whole chicken",
        "Ingredient2": "",
        "Ingredient3": "",
        "Ingredient4": "",
        "Ingredient5": "",
        "Ingredient6": "",
        "Ingredient7": "",
        "Ingredient8": "",
        "Ingredient9": "",
        "Ingredient10": "",
        "Instructions": "",
        "Estimated_time": "",
        "Source": ""
    },
    {
        "Id": 2,
        "Name": "Tacos",
        "Ingredient1": "1lb ground beef",
        "Ingredient2": "",
        "Ingredient3": "",
        "Ingredient4": "",
        "Ingredient5": "",
        "Ingredient6": "",
        "Ingredient7": "",
        "Ingredient8": "",
        "Ingredient9": "",
        "Ingredient10": "",
        "Instructions": "",
        "Estimated_time": "",
        "Source": ""
    },
    {
        "Id": 3,
        "Name": "Chicken Parmesan",
        "Ingredient1": "Chicken breasts",
        "Ingredient2": "",
        "Ingredient3": "",
        "Ingredient4": "",
        "Ingredient5": "",
        "Ingredient6": "",
        "Ingredient7": "",
        "Ingredient8": "",
        "Ingredient9": "",
        "Ingredient10": "",
        "Instructions": "",
        "Estimated_time": "",
        "Source": ""
    }
]`

## POST retrieve all ingredients for given recipes
### Request
`POST /shoppingList HTTP/1.1
Host: localhost:8080
[{"id" : 1},
{"id" : 3}]`

### Example Response
`[
    "whole chicken",
    "fingerling potatoes",
    "baby carrots",
    "apple juice",
    "Chicken breasts",
    "Parmesan cheese",
    "mozzarella cheese",
    "jar of pasta sauce"
]`
