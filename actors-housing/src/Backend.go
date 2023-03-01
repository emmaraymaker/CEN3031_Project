package Backend

import (
	"fmt"
)

type Member struct {
	UID   string
	name  string
	phone string
	email string
}

type MemberList struct {
	ml []Member
}

type ListingSet struct {
	ls []Listing
}

type Listing struct {
	location string
	price    int
	// Extra filters, gender/sex, smoker, etc...
}

func (re MemberList) addMember(u string, n string, p string, e string) {
	nm := Member{
		UID:   u,
		name:  n,
		phone: p,
		email: e,
	}
	re.ml = append(re.ml, nm)
}

func (re ListingSet) addListing(l string, p int) {
	nl := Listing{
		location: l,
		price:    p,
	}
	re.ls = append(re.ls, nl)
}

func main() {
	// Test set for pre-front-end integration

	var memL MemberList
	var lisL ListingSet

	var uid string
	var name string
	var phone string
	var email string
	fmt.Println("Hello, please enter your Union Id:")
	fmt.Scanln(&uid)
	fmt.Println("Please enter your name:")
	fmt.Scanln(&name)
	fmt.Println("Please enter your Phone number:")
	fmt.Scanln(&phone)
	fmt.Println("Please enter your Email address:")
	fmt.Scanln(&email)
	memL.addMember(uid, name, phone, email)

	for {
		fmt.Println("Please select a menu option:\n1: Rent\n2:List")

		var option string
		fmt.Scanln(&option)

		var location string
		var price int
		if option == "1" {
			fmt.Println("displaying listings:")
			fmt.Println(lisL.ls[0].location)
		} else if option == "2" {

			fmt.Println("Please enter the Listing's Location:")
			fmt.Scanln(&location)
			fmt.Println("Please enter the Listing's Price:")
			fmt.Scanln(&price)
			lisL.addListing(location, price)
		} else {
			break
		}
	}
}
