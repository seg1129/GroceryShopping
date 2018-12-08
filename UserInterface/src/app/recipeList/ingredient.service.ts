import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { catchError, tap } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';

import { IIngredient } from './ingredient';

@Injectable({
  providedIn: 'root'
})
export class IngredientService {
  // private productUrl = 'http://localhost:8080/allRecipes';

  constructor(private http: HttpClient){}

  getIngredients(): Observable<any>{
    // return this.http.get<IIngredient[]>(this.httpClient.post("http://localhost:8080/shoppingList",
    //     {
    //       [{"id" : 1},
    //       {"id" : 2},
    //       {"id" : 3}]
    //     })
    //     .subscribe(
    //         data => {
    //             console.log("POST Request is successful ", data);
    //         },
    //         error => {
    //             console.log("Error", error);
    //         }
    //     )).pipe(
    //   tap(data => console.log('All: ' + JSON.stringify(data))),
    //   catchError(this.handleError)
    // );
    return this.http.post("http://localhost:8080/shoppingList",[
        {
          "id" : 1},
          {"id" : 2},
          {"id" : 3
        }])
  }
  private handleError(err: HttpErrorResponse){
    console.error(err);
    return throwError(err);
  }
}
