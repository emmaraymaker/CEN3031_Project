import { HttpClient } from '@angular/common/http';
import { UserComponent } from './user.component';

describe('UserComponent', () => {
  let component: UserComponent;
  let httpClient: HttpClient;

  beforeEach(() => {
    httpClient = new HttpClient(null);
    component = new UserComponent(httpClient);
  });

  it('should have a user', () => {
    expect(component.user).toBeDefined();
  });

  it('should have an empty user', () => {
    expect(component.user.first_name).toEqual('');
    expect(component.user.last_name).toEqual('');
    expect(component.user.email).toEqual('');
    expect(component.user.id).toEqual(0);
    expect(component.user.password).toEqual('');
  });
});
