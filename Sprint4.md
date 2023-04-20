Detail work you've completed in Sprint 4

In Sprint 4 we were able to make progress on the uncompleted issues from Sprint 3. We were able to fully integrate the backend and frontend for our login, registration, user, and listings components. Once a user fills out the registration form and clicks submit their information is sent to be stored in the backend. The same was accomplished for the listings functionality. A user is now about to add their listings, keep track of their listings, and view all currently available listings.

List frontend unit and Cypress tests:
  Cypress:
  -Check is AppComponent mounts

  -Check is AppComponent page is titled "Actor's Housing"

  -Check if Login page mounts

  -Check if Login page's email and password forms are initially blank

  -Check if Registration page mounts
  
  Karma:
    -Listings.component.spec.ts:
    - This test creates an instance of the ListingsComponent, sets some input values, calls the addListing() function, and checks if a new listing is added to the Listings array with the correct values. Note that this test does not require any Angular testing utilities since we are not testing the component's template or any Angular-specific functionality.

    -Login.component.spec.ts:
    -This test creates an instance of the LoginComponent, sets the email property, calls the onSubmit() function, and checks if the email property is still set to the same value after the function is called. Note that this test does not require any Angular testing utilities since we are not testing the component's template or any Angular-specific functionality.

    -Registration.component.spec.ts:
    -This test creates an instance of the RegistrationComponent and checks if the required form controls are created correctly in the createForm() function. Note that we need to create a FormBuilder instance and pass it to the RegistrationComponent constructor in order to create the form controls.

    -User.component.spec.ts:
    -This test creates an instance of the UserComponent and checks if the user property is defined and has empty values for its fields. Note that we need to create an instance of HttpClient and pass it to the constructor to make HTTP requests.

    -AppComponent.component.spec.ts:
    -This test creates an instance of the AppComponent and checks if the title property is defined and has the correct value. Note that this test does not require any Angular testing utilities since we are not testing the component's template or any Angular-specific functionality.


Introducing a specialized platform for NYC actors to find and list housing based on specific criteria, including dates, price, and preferences. Actors will register using their Actors Equity Union number and verified email address, making it easier for them to find short- or long-term leases that meet their needs. The project uses Golang and Angular technologies and aims to support and empower the professional community of live theatrical performance in their journeys.



-Check if Registration page's first name, last name, id, email, password, and confirm password forms are initially blank

List backend unit tests:

List backend unit tests

ListingSet tests:

TestAddListing - Tests the backend's ability to add a listing

TestSearchListing - Tests the backend's ability to search for a listing

TestRemoveListing - Tests the backend's ability to remove a listing

MemberList tests:

TestAddMember - Tests the backend's ability to add a member

TestSearchMember - Tests the backend's ability to search for a member

TestRemoveMember - Tests the backend's ability to remove a member

Show updated documentation for your backend API 

  - getListings: returns the listings of the current user
  - getAllListings: returns all of the currently available listings
  - getCurrentUser: returns the current users in the Members list
  - addCurrentUser: takes in a pointer to the frontend data, and stores the new user information in the Members list
  - addToListing: takes in a pointer to the frontend data, and stores the new listings
