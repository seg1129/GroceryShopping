import { Component } from '@angular/core';

@Component({
  selector: 'recipe-root',
  template:
    `<div>
        <h1>{{pageTitle}}</h1>
        <recipe-list></recipe-list>
     </div>`
})
export class AppComponent {
  pageTitle: string = 'Recipe List';
}
