Detail work you've completed in Sprint 4

In Sprint 4 we were able to make progress on the uncompleted issues from Sprint 3. We were able to fully integrate the backend and frontend for our login, registration, user, and listings components. Once a user fills out the registration form and clicks submit their information is sent to be stored in the backend. The same was accomplished for the listings functionality. A user is now about to add their listings, keep track of their listings, and view all currently available listings.

List frontend unit and Cypress tests


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
