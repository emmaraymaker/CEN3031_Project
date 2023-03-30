import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LoginComponent } from '../login/login.component';

describe('LoginComponent mounts', () => {
  it('mounts', () => {
    cy.mount(LoginComponent)
  })
})

it('Email should start off as blank', () => {
  cy.mount(LoginComponent)
  cy.get('[formControlName="email"]').should('have.text', '')
})

it('Password should start off as blank', () => {
  cy.mount(LoginComponent)
  cy.get('[formControlName="password"]').should('have.text', '')
})
