import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { User } from '../user';
import { FormBuilder } from '@angular/forms';
 
@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit{
  constructor(private httpClient: HttpClient) {} 
  ngOnInit(): void {}
  user: User = {
    first_name: '',
    last_name: '',
    email: '',
    id: 0,
    password: ''
  };

 onSubmit(): void{ // need to pass information to backend 
  first_name : this.user.first_name;
  last_name : this.user.last_name;
  email : this.user.email;
  id : this.user.id;
  password : this.user.password;
 }

}

