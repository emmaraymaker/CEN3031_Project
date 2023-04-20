import { Component, OnInit } from '@angular/core';
import { User } from '../user';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';

interface userProfile{
  firstName: string
  lastName: string
  email: string
  unionID: string
  password: string
}

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit{
  firstName = ''
  lastName = ''
  email = ''
  unionID = ''
  password = ''
  User: userProfile [] = [] // array to hold user information

  constructor(private httpClient: HttpClient) {}

  async ngOnInit() {await this.currentUser()}

  async currentUser() {
    this.User = await firstValueFrom(this.httpClient
    .get<userProfile[]>('/backend/user'))
  }
}
