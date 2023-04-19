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
  //{FirstName: "Kyle", LastName: "Lina", Email: "yay@gmail.com", UnionNumber: 100, Password: "anything"},
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
	{State: "New York", City: "Brooklyn", Street: "452", Unit: "12", Zipcode: "10010", Price: "1400"},
	{State: "New York", City: "Chelsea", Street: "899", Unit: "C", Zipcode: "10014", Price: "4000"},
}

func main() {
  r := gin.Default()

  backend := r.Group("/backend")
  {
    backend.GET("/user",getCurrentUser)
    backend.GET("/", getListings)
    backend.POST("/", addToListing)
	backend.GET("/listings",getAllListings)
  }

  r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// get Listings function
func getListings(c *gin.Context){
  c.IndentedJSON(http.StatusOK, listings)
}

func getAllListings(c *gin.Context){
	c.IndentedJSON(http.StatusOK, allListings)
  }

func getCurrentUser(c *gin.Context){
  c.IndentedJSON(http.StatusOK, currentMember)
}

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
		c.Status(http.StatusNoContent)
}