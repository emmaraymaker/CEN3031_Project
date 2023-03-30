import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { User } from '../user';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';

@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit{
  requiredForm: FormGroup;
  constructor(private fb: FormBuilder) {
    this.createForm();
  }
  createForm() {
    this.requiredForm = this.fb.group({
      first_name: ['', [Validators.required,
        Validators.minLength(2),
        Validators.pattern("^[a-zA-Z ]*$")
      ]],
      last_name: ['', [Validators.required,
        Validators.minLength(2),
        Validators.pattern("^[a-zA-Z ]*$")
      ]],
      email: ['', [Validators.required,
        Validators.pattern("^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")
      ]],
      unionid: ['', [Validators.required,
        Validators.minLength(6),
        Validators.pattern("^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")
      ]],
      password: ['', [Validators.required,
        Validators.minLength(9),
        Validators.pattern("^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")
      ]],
      confirmpassword: ['', [Validators.required,
        Validators.minLength(9),
        Validators.pattern("^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$")
      ]],
    });
  }

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

