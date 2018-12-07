import { Component, OnInit } from '@angular/core';
import { IRecipe } from './recipe';
import { RecipeService } from './recipe.service';

@Component({
  selector: 'recipe-list',
  templateUrl: './recipe-list.component.html'
})
export class RecipeListComponent implements OnInit {
  errorMessage = '';
  pageTitle: string = 'Recipe List';
  recipes: any[] = [];
  filteredRecipes: IRecipe[] = [];

  _recipeFilter = ' ';
  get recipeFilter(): string {
    return this._recipeFilter;
    }

  set recipeFilter(value: string) {
    this._recipeFilter = value;
    this.filteredRecipes= this.recipeFilter ? this.filter(this.recipeFilter) : this.recipes
  }

  constructor(private recipeService: RecipeService) {
    this.filteredRecipes = this.recipes;
    this.recipeFilter = '';
  }

  filter(filterBy: string): IRecipe[] {
    filterBy = filterBy.toLocaleLowerCase();
    return this.recipes.filter((recipe: IRecipe ) =>
        recipe.Name.toLocaleLowerCase().indexOf(filterBy)!==-1)
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
