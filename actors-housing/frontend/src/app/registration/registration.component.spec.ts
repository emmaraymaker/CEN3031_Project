import { FormBuilder } from '@angular/forms';
import { RegistrationComponent } from './registration.component';

describe('RegistrationComponent', () => {
  let component: RegistrationComponent;
  let formBuilder: FormBuilder;

  beforeEach(() => {
    formBuilder = new FormBuilder();
    component = new RegistrationComponent(formBuilder);
  });

  it('should create a form with required fields', () => {
    expect(component.requiredForm.get('first_name')).toBeTruthy();
    expect(component.requiredForm.get('last_name')).toBeTruthy();
    expect(component.requiredForm.get('email')).toBeTruthy();
    expect(component.requiredForm.get('unionid')).toBeTruthy();
    expect(component.requiredForm.get('password')).toBeTruthy();
    expect(component.requiredForm.get('confirmpassword')).toBeTruthy();
  });
});
