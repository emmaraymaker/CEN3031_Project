package Backend

type Member struct {
	UID      string
	name     string
	phone    string
	email    string
	isLister bool
	isTenant bool
}

type MemberList struct {
	ml []Member
}

type ListingSet struct {
	ls []Listing
}

type Listing struct {
	// Info

	// Format: (Country) State, City, Street, Unit
	// Notes: Comma delimited, currently only USA
	location string

	// Notes:
	// Currently, only supports full US dollar amounts
	// Considering change to double/float to support change
	price int

	// Preferences:

	// current gender preferences, M, F, N
	// M - male, F - Female, N - Non-binary
	// Notes:
	// It's a list, to specify all, it'd have to include all 3 or none
	// Not certain about the current implementation
	gender []byte

	// Notes:
	smokerFriendly bool

	// Notes:
	// Future implementation might have more options, such as dogs only
	petFriendly bool

	// Extra filters, gender/sex, smoker, etc...
}

func (re MemberList) addMember(u string, n string, p string, e string, l bool, t bool) {
	nm := Member{
		UID:      u,
		name:     n,
		phone:    p,
		email:    e,
		isLister: l,
		isTenant: t,
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

	/*
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
		memL.addMember(uid, name, phone, email, isL, isT)

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
	*/
}
