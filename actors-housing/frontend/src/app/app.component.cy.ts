import { AppComponent } from './app.component';
import {TestBed} from "@angular/core/testing";

describe('StepperComponent', () => {
  it('mounts', () => {
    cy.mount(AppComponent)
  }),
    it(`should have as title 'actors-housing-project'`, () => {
      const fixture = TestBed.createComponent(AppComponent);
      const app = fixture.componentInstance;
      expect(app.title).equal('Actors Housing');
    })
})
