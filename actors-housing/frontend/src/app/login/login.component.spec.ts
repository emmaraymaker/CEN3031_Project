import { LoginComponent } from './login.component';

describe('LoginComponent', () => {
  let component: LoginComponent;

  beforeEach(() => {
    component = new LoginComponent();
  });

  it('should set email on submit', () => {
    const email = 'test@example.com';
    component.email = email;

    component.onSubmit();

    expect(component.email).toEqual(email);
  });
});
