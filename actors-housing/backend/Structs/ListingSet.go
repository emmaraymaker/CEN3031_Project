package backen

import "errors"

type ListingSet struct {
	Ls []Listing
}

type Listing struct {
	// Info

	// ID of
	lid int

	// Format: (Country) State, City, Street, Unit
	// Notes: Comma delimited, currently only USA
	Location string

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

func (re *ListingSet) AddListing(id int, l string, p int, g []byte, sF bool, pF bool) {
	nl := Listing{
		lid:            id,
		Location:       l,
		price:          p,
		gender:         g,
		smokerFriendly: sF,
		petFriendly:    pF,
	}
	re.Ls = append(re.Ls, nl)
}

func (re ListingSet) SearchListing(lid int) (Listing, error) {
	for i := 0; i <= len(re.Ls); i++ {
		if re.Ls[i].lid == lid {
			return re.Ls[i], nil
		}
	}
	return Listing{}, errors.New("Listing not found")
}

func (re *ListingSet) RemoveListing(lid int) {
	for i := 0; i <= len(re.Ls); i++ {
		if re.Ls[i].lid == lid {
			re.Ls = append(re.Ls[:i], re.Ls[i+1:]...)
			return
		}
	}
}
