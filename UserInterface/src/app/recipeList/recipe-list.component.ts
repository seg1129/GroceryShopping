import { Component } from '@angular/core';
import { IRecipe } from './recipe';
import { RecipeService } from './recipe.service';

@Component({
  selector: 'recipe-list',
  templateUrl: './recipe-list.component.html'
})
export class RecipeListComponent {
  errorMessage = '';

  pageTitle: string = 'Recipe List';
  recipes: any[] = [];

  filteredRecipes: IRecipe[] = [];
  // recipes: IRecipe[] = [];

  constructor(private recipeService: RecipeService) {

  }


ngOnInit(): void {
  this.recipeService.getRecipes().subscribe(
    recipes => {
      this.recipes = recipes;
      this.filteredRecipes = this.recipes;
    },
    error => this.errorMessage = <any>error
  );

}
}
