Detail work you've completed in Sprint 2

For sprint 2 we were able to make progress on the uncompleted issues from sprint 1. We added on to our implementations of the login and registration pages for both the backend and the frontend. We were not able to fully integrate the frontend and the backend and plan to continue working on the issues. The frontend presents users with the option to login or register. Once the integration is successfully completed the user information entered in the registration form will be stored in the backend, and used to verify user credentials upon logging in.


List unit tests and Cypress test for frontend


List unit tests for backend

Currently in the works, one general test is the current progress made

Documentation for the backend API 

MemberList

  Functions:
  
    searchMember(UID) (member, error)
    Searches through the list of members and returns the first one with the proper user id.
    
    addMember(UID, Name, Phone, Email, Lister status, Tenant status)
    Adds a member to the list.
    Currently the implementation does not have proper checks on the input, later iterations should check the arguments for validity.

    removeMember(UID)
    Removes a member from the list, returns nothing.

    Later implementations might return whether the member was successfully removed or not.


ListingSet

  Functions:
  
    searchListing(LID) (member, error)
    Searches through the set of listings and returns the first one with the proper listing id.

    addListing(LID, location, Price, genderPreference, smokerFriendly, petFriendly)
    Adds a Listing to the set.
    Currently the implementation does not have proper checks on the input, later iterations should check the arguments for validity.

    removeListing(LID)
    Removes a List from the set, returns nothing.

    Later implementations might return whether the Listing was successfully removed or not.
