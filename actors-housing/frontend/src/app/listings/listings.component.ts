import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { firstValueFrom } from 'rxjs';

interface userListings{
  state: string
  city: string
  street: string
  unit: string
  zipcode: string
  price: string
}

@Component({
  selector: 'app-listings',
  templateUrl: './listings.component.html',
  styleUrls: ['./listings.component.css']
})
export class ListingsComponent implements OnInit{
  state = ''
  city = ''
  street = ''
  unit = ''
  zipcode = ''
  price = ''
  Listings: userListings [] = [] // array to hold current users listings

  constructor(private httpClient: HttpClient) {}

  async ngOnInit() {await this.myCurrentListings()}

  async myCurrentListings() {
    this.Listings = await firstValueFrom(this.httpClient
    .get<userListings[]>('/backend'))
  }

  // function to add listing
  async addListing(){ 
    await firstValueFrom(this.httpClient.post('/backend', {
      state: this.state,
      city: this.city,
      street: this.street,
      unit: this.unit,
      zipcode: this.zipcode,
      price: this.price
    }))
      
    await this.myCurrentListings()

      // clear out input
      this.state = ''
      this.city = ''
      this.street = ''
      this.unit = ''
      this.zipcode = ''
      this.price = ''
    }
  }