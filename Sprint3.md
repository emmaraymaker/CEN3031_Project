Detail work you've completed in Sprint 3

In Sprint 3 we were able to make progress on the uncompleted issues from Sprint 2. We improved our implementation of the backend and frontend regarding the login and registration pages. We discovered new issues such as allowing the user the functionality of adding a listing. We successfully created a listing component that takes in and stores the user’s listing information.


List frontend unit tests:

-Check is AppComponent mounts

-Check is AppComponent page is titled "Actor's Housing"

-Check if Login page mounts

-Check if Login page's email and password forms are initially blank

-Check if Registration page mounts

-Check if Registration page's first name, last name, id, email, password, and confirm password forms are initially blank


List backend unit tests

List backend unit tests:

-Check if users and listings are added correctly

-Check if users and listings can be retrieved from the data structures correctly

-Check if users and listings can be removed

-Check if users and listings can be instaniated


Show updated documentation for your backend API 

The Backend API has been updated to be more in-line with its function.

The Functions are as follows:

Users:

  RegisterUser, takes in one parameter, User and appends it to the list of members

  GetUser, takes in one parameter, UID and returns the User who has the ID

  UpdateUser, takes in multiple parameters, UID, and User, UID specifying which user to change and User, containing the old user with the new modifications 

Listings:

  SetListing, takes in one parameter, Listing and appends it to the list of Listings

  GetListing, takes in one parameter, the listing ID and returns the Listing that has the ID

  UpdateListing, takes in multiple parameters, listing id and, Listing ID specifying which listing to change and listing, containing the old listing with the new modifications 
