import { Component } from '@angular/core';

@Component({
  selector: 'recipe-root',
  template:
  `
  <nav class='navbar navbar-expand navbar-light bg-light'>
      <ul class='nav nav-pills'>
        <li><a class='nav-link' routerLinkActive='active' [routerLink]="['/recipeList']">Home</a></li>
      </ul>
  </nav>
  <div class='container'>
    <router-outlet></router-outlet>
  </div>
  `
})
export class AppComponent {
  pageTitle: string = 'Cooking App';
}
