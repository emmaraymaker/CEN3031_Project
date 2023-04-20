import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ListingsComponent } from './listings.component';
import {LoginComponent} from "../login/login.component";

describe('ListingComponent mounts', () => {
  it('mounts', () => {
    cy.mount(ListingsComponent)
  })
})

it('State should start off as blank', () => {
  cy.mount(ListingsComponent)
  cy.get('[formControlName="state"]').should('have.text', '')
})

it('City should start off as blank', () => {
  cy.mount(ListingsComponent)
  cy.get('[formControlName="city"]').should('have.text', '')
})

it('Street should start off as blank', () => {
  cy.mount(ListingsComponent)
  cy.get('[formControlName="street"]').should('have.text', '')
})

it('Unit should start off as blank', () => {
  cy.mount(ListingsComponent)
  cy.get('[formControlName="unit"]').should('have.text', '')
})

it('Zipcode should start off as blank', () => {
  cy.mount(ListingsComponent)
  cy.get('[formControlName="zipcode"]').should('have.text', '')
})


it('Description should start off as blank', () => {
  cy.mount(ListingsComponent)
  cy.get('[formControlName="description"]').should('have.text', '')
})



