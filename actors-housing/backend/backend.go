package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

type MemberList struct {
	Ml []Member
}

type Member struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	UnionNumber int `json:"unionNumber"`
	Password string `json:"password"`
}

var currentMember = []Member{
  
}

type ListingSet struct {
	Ls []Listing
}

type Listing struct {
	State	string `json:"state"`
	City	string	`json:"city"`
	Street	string `json:"street"`
	Unit	string `json:"unit"`
	Zipcode	string `json:"zipcode`
	Price	string `json:"price"`
 }

var listings = []Listing{

}

var allListings = []Listing{
	{State: "New York", City: "Brooklyn", Street: "452 London St", Unit: "12", Zipcode: "10010", Price: "1400"},
	{State: "New York", City: "Chelsea", Street: "105 Inverness St", Unit: "C15", Zipcode: "10014", Price: "4000"},
	{State: "New York", City: "Chinatown", Street: "9080 E. Pawnee St", Unit: "5", Zipcode: "10111", Price: "1700"},
	{State: "New York", City: "Greenwich Village", Street: "899 South Blvd", Unit: "25", Zipcode: "10589", Price: "1780"},
	{State: "New York", City: "Harlem", Street: "Avalon Bay", Unit: "2", Zipcode: "10560", Price: "3500"},
	{State: "New York", City: "Long Island City", Street: "301 W 21st", Unit: "9", Zipcode: "10014", Price: "2100"},
	{State: "New York", City: "Manhattan", Street: "101 W 15th St", Unit: "B4", Zipcode: "12015", Price: "2250"},
	{State: "New York", City: "Queens", Street: "461 Dean", Unit: "566", Zipcode: "11111", Price: "1600"},
	{State: "New York", City: "SoHo", Street: "Montague St", Unit: "7855", Zipcode: "10013", Price: "1450"},
	{State: "New York", City: "Upper East Side", Street: "Avalon West", Unit: "58", Zipcode: "10104", Price: "3750"},
}

func main() {
  r := gin.Default()

  backend := r.Group("/backend")
  {
    backend.GET("/user",getCurrentUser)
	backend.POST("/user",addCurrentMember)
    backend.GET("/", getListings)
    backend.POST("/", addToListing)
	backend.GET("/listings",getAllListings)
  }

  r.Run("0.0.0.0:8080")
}

// get current usersrs function
func getListings(c *gin.Context){
  c.IndentedJSON(http.StatusOK, listings)
}

/// get all listings 
func getAllListings(c *gin.Context){
	c.IndentedJSON(http.StatusOK, allListings)
  }

  // get current user
func getCurrentUser(c *gin.Context){
  c.IndentedJSON(http.StatusOK, currentMember)
}

// add the current user on login/ registration
func addCurrentMember(c *gin.Context) {
  memberContents := Member{}
  c.Bind(&memberContents)

  thisMember := Member{
    FirstName:memberContents.FirstName,
    LastName:	memberContents.LastName,
    Email:	memberContents.Email,
    UnionNumber:	memberContents.UnionNumber,
    Password: memberContents.Password,
  }
  
  currentMember = append(currentMember, thisMember)
  c.Status(http.StatusNoContent)
}

func addToListing(c *gin.Context) {
		listingContents := Listing{}
		c.Bind(&listingContents)

    newListing := Listing{
      State:listingContents.State,
      City:	listingContents.City,
      Street:	listingContents.Street,
      Unit:	listingContents.Unit,
      Zipcode: listingContents.Zipcode,
      Price: listingContents.Price,
    }

    listings = append(listings, newListing)
	allListings = append(allListings,newListing)
		c.Status(http.StatusNoContent)
}