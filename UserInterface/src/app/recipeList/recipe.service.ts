import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { catchError, tap } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';

import { IRecipe } from './recipe';

@Injectable({
  providedIn: 'root'
})
export class RecipeService {
  private productUrl = 'http://localhost:8080/allRecipes';

  constructor(private http: HttpClient){}

  getRecipes(): Observable<IRecipe[]>{
    return this.http.get<IRecipe[]>(this.productUrl).pipe(
      tap(data => console.log('All: ' + JSON.stringify(data))),
      catchError(this.handleError)
    );
  }
  private handleError(err: HttpErrorResponse){
    console.error(err);
    return throwError(err);
  }
}
