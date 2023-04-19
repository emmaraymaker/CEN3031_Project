import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';

interface allListings{
  state: string
  city: string
  street: string
  unit: string
  zipcode: string
  price: string
}

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  state = ''
  city = ''
  street = ''
  unit = ''
  zipcode = ''
  price = ''
  Listings: allListings [] = [] // array to hold current all listings

  constructor(private httpClient: HttpClient) {}

  async ngOnInit() {await this.allListings()}

  async allListings() {
    this.Listings = await firstValueFrom(this.httpClient
    .get<allListings[]>('/backend/listings'))
  }
}
