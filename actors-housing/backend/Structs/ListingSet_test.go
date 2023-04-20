package backen_test

import (
	backen "CEN3031_Project-main/actors-housing/backend/Structs"
	"testing"
)

func TestAddListing(t *testing.T) {
	re := backen.ListingSet{}

	// add a new listing
	re.AddListing(1, "USA, California, San Francisco, 123 Main St", 2000, []byte{'M'}, true, false)

	// verify that the listing was added
	if len(re.Ls) != 1 {
		t.Errorf("expected 1 listing, but got %d", len(re.Ls))
	}
}

func TestSearchListing(t *testing.T) {
	re := backen.ListingSet{}

	// add a new listing
	re.AddListing(1, "USA, California, San Francisco, 123 Main St", 2000, []byte{'M'}, true, false)

	// search for the added listing
	listing, err := re.SearchListing(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// verify that the correct listing was found
	if listing.Location != "USA, California, San Francisco, 123 Main St" {
		t.Errorf("expected location %q, but got %q", "USA, California, San Francisco, 123 Main St", listing.Location)
	}
}

func TestRemoveListing(t *testing.T) {
	re := backen.ListingSet{}

	// add a new listing
	re.AddListing(1, "USA, California, San Francisco, 123 Main St", 2000, []byte{'M'}, true, false)

	// remove the added listing
	re.RemoveListing(1)

	// verify that the listing was removed
	if len(re.Ls) != 0 {
		t.Errorf("expected 0 listings, but got %d", len(re.Ls))
	}
}
