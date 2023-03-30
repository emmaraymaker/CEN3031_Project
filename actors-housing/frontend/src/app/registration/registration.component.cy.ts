import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RegistrationComponent } from '../registration/registration.component';

describe('LoginComponent mounts', () => {
  it('mounts', () => {
    cy.mount(RegistrationComponent)
  })
})

it('First name should start off as blank', () => {
  cy.mount(RegistrationComponent)
  cy.get('[formControlName="first_name"]').should('have.text', '')
})

it('Last name should start off as blank', () => {
  cy.mount(RegistrationComponent)
  cy.get('[formControlName="last_name"]').should('have.text', '')
})

it('Email should start off as blank', () => {
  cy.mount(RegistrationComponent)
  cy.get('[formControlName="email"]').should('have.text', '')
})

it('Union ID should start off as blank', () => {
  cy.mount(RegistrationComponent)
  cy.get('[formControlName="unionid"]').should('have.text', '')
})

it('Password should start off as blank', () => {
  cy.mount(RegistrationComponent)
  cy.get('[formControlName="password"]').should('have.text', '')
})

it('Confirm password should start off as blank', () => {
  cy.mount(RegistrationComponent)
  cy.get('[formControlName="confirmpassword"]').should('have.text', '')
})
