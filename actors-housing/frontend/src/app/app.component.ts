import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Actors Housing';


  ngOnInit(){
    setTimeout(
      () => {
        this.title = "Hey, its been updated"
        console.log('helo?')
      },
      2000
    )
  }
}
