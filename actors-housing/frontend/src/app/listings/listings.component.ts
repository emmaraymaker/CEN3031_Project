import { Component } from '@angular/core';

interface userListings{
  state: string
  city: string
  street: string
  unit: string
  description: string
  zipcode: string
}

@Component({
  selector: 'app-listings',
  templateUrl: './listings.component.html',
  styleUrls: ['./listings.component.css']
})
export class ListingsComponent {
  state = ''
  city = ''
  street = ''
  unit = ''
  zipcode = ''
  description = ''
  Listings: userListings [] = [] // array to hold current users listings

  // function to add listing
  addListing(){
    this.Listings.push({
      state: this.state,
      city: this.city,
      street: this.street,
      unit: this.unit,
      zipcode: this.zipcode,
      description: this.description
    })
    // clear out input
    this.state = ''
    this.city = ''
    this.street = ''
    this.unit = ''
    this.zipcode = ''
    this.description = ''
  }
}
