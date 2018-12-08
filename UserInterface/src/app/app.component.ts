import { Component } from '@angular/core';

@Component({
  selector: 'recipe-root',
  template:
    `<div>
        <h1>{{pageTitle}}</h1>
        <router-outlet></router-outlet>
     </div>`
})
export class AppComponent {
  pageTitle: string = 'Cooking App';
}

// <recipe-list></recipe-list>
// <app-ingredients></app-ingredients>
