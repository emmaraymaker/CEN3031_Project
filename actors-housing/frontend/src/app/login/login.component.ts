import { Component, Input, OnInit } from '@angular/core';
import { User } from '../user';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';

interface loginCredentials{
  email: string
  password: string
}

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit{
email : string = '';
password : string = '';
User: loginCredentials [] = []

constructor(private httpClient: HttpClient) {}

async ngOnInit() {await this.currentUser()}
 
async currentUser() {
   this.User = await firstValueFrom(this.httpClient
   .get<loginCredentials[]>('/backend/currentUser'))
 }

async onSubmit(){ 
  await firstValueFrom(this.httpClient.post('/backend/currentUser',{
    email: this.email,
  }))
    await this.currentUser()
  }
}
