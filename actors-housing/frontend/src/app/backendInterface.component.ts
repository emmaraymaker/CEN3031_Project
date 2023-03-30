import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Book } from './book';

@Injectable({
  providedIn: 'root'
})
export class backendService {
  private userApiUrl = 'http://localhost:8000/user';
  private listingApiUrl = 'http://localhost:8000/listing';

  constructor(private http: HttpClient) { }
  
  // need to add/import user class
  // this should work once the matching backend functions are complete
  regUser(usr: User): Observable<User> {
    return this.http.post<User>(this.userApiUrl, usr);
  }
  // example of how to use it here:
  /*
    registerUser(newUser: User): void {
    if (!User) { return; }
    this.backendService.regUser(newUser)
      .subscribe(() => {
        this.http.post(this.userApiUrl + '/reg', newUser).subscribe(() => {
          // Handle the response from the backend if needed
        });
      });
  }
    */
    

  /*
  getUser(id: string): Observable<Book> {
    const url = `${this.apiUrl}/${id}`;
    return this.http.get<Book>(url);
  }

  setUser(book: Book): Observable<Book[]> {
    return this.http.post<Book[]>(this.apiUrl, book);
  }

  searchUser(book: Book): Observable<Book[]> {
    const url = `${this.apiUrl}/${book.id}`;
    return this.http.put<Book[]>(url, book);
  }

  deleteBook(id: string): Observable<Book[]> {
    const url = `${this.apiUrl}/${id}`;
    return this.http.delete<Book[]>(url);
  }
  */
}
