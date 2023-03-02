import { Component, Input } from '@angular/core';
import { User } from '../user';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
email : string = '';
password : string = '';

onSubmit(): void{
  // need to make api call to backend to verify login credentials
  this.email = this.email;
}

}
