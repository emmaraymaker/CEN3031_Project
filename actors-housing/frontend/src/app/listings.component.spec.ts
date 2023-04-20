import { ListingsComponent } from './listings.component';

describe('ListingsComponent', () => {
  let component: ListingsComponent;

  beforeEach(() => {
    component = new ListingsComponent();
  });

  it('should add a listing to the Listings array', () => {
    const initialLength = component.Listings.length;
    component.state = 'California';
    component.city = 'Los Angeles';
    component.street = '123 Main St';
    component.unit = 'Apt 1';
    component.zipcode = '90001';
    component.description = 'A cozy apartment';

    component.addListing();

    expect(component.Listings.length).toEqual(initialLength + 1);
    expect(component.Listings[initialLength]).toEqual({
      state: 'California',
      city: 'Los Angeles',
      street: '123 Main St',
      unit: 'Apt 1',
      zipcode: '90001',
      description: 'A cozy apartment'
    });
  });
});
