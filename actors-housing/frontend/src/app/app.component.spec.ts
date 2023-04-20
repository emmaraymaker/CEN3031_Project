import { AppComponent } from './app.component';

describe('AppComponent', () => {
  let component: AppComponent;

  beforeEach(() => {
    component = new AppComponent();
  });

  it('should have a title', () => {
    expect(component.title).toBeDefined();
  });

  it('should have the correct title', () => {
    expect(component.title).toEqual('Actors Housing');
  });
});
