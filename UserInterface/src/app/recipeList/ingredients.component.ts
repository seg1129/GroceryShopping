import { Component, OnInit } from '@angular/core';
import { IngredientService } from './ingredient.service';

@Component({
  selector: 'app-ingredients',
  templateUrl: './ingredients.component.html',
  styleUrls: ['./ingredients.component.css']
})
export class IngredientsComponent implements OnInit {
  ingredients: any[] = [];
  errorMessage = '';
  pageTitle: string = 'Shopping List';

  constructor(private ingredientService: IngredientService) { }

  ngOnInit(): void {
    this.ingredientService.getIngredients().subscribe(
      ingredients => {
        this.ingredients = ingredients;
      },
      error => this.errorMessage = <any>error
  );

}
  }
