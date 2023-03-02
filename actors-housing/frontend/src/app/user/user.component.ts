import { Component } from '@angular/core';
import { User } from '../user';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent {
  user: User = {
    first_name:'',
    last_name: '',
    email: '',
    id: 0,
    password: ''
  };

  constructor(private http: HttpClient){}
}
