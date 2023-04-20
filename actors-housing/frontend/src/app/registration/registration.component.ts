import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { User } from '../user';
import {FormBuilder, FormGroup, Validators} from '@angular/forms';
import { firstValueFrom } from 'rxjs';

interface userInformation{
  firstName: string
  lastName: string
  email: string
  password: string
  unionID: string
}


@Component({
  selector: 'app-registration',
  templateUrl: './registration.component.html',
  styleUrls: ['./registration.component.css']
})
export class RegistrationComponent implements OnInit{
  requiredForm: FormGroup;
  constructor(private fb: FormBuilder, private httpClient: HttpClient) {
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

  firstName: ''
  lastName: ''
  email: ''
  unionID: ''
  password: ''
  User: userInformation [] = []

  async ngOnInit() {await this.thisUserInformation()}

  async thisUserInformation() {
    this.User = await firstValueFrom(this.httpClient
    .get<userInformation[]>('/backend/user'))
  }

  // function to add user
  async onSubmit(){ 
    await firstValueFrom(this.httpClient.post('/backend/user', {
      firstName : this.firstName,
      lastName : this.lastName,
      email : this.email,
      uionID : this.unionID,
      password : this.password,
    }))
      
    await this.thisUserInformation()
  }

}