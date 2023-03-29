package main

import (
	backend "test/Structs"
	"testing"
	_ "testing"
)

func TestAddmL(t *testing.T) {
	var memL backend.MemberList
	memL.AddMember("1", "Test, t", "888-888-8888", "test@test.com", "password", false, false)
	got := len(memL.Ml)
	want := 1

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestSearchmL(t *testing.T) {
	var memL backend.MemberList
	memL.AddMember("1", "Test, t", "888-888-8888", "test@test.com", "password", false, false)
	var mem, _ = memL.SearchMember("1")
	var name = mem.Name
	got := name
	want := "Test, t"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestRemovemL(t *testing.T) {
	var memL backend.MemberList
	memL.AddMember("1", "Test, t", "888-888-8888", "test@test.com", "password", false, false)
	memL.RemoveMember("1")
	got := len(memL.Ml)
	want := 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestAddlS(t *testing.T) {
	var ls backend.ListingSet
	ls.AddListing(1, "Test, t", 8888, []byte{}, false, false)
	got := len(ls.Ls)
	want := 1

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestSearchlS(t *testing.T) {
	var ls backend.ListingSet
	ls.AddListing(1, "Test, t", 8888, []byte{}, false, false)
	var l, _ = ls.SearchListing(1)
	var loc = l.Location
	got := loc
	want := "Test, t"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestRemovelS(t *testing.T) {
	var ls backend.ListingSet
	ls.AddListing(1, "Test, t", 8888, []byte{}, false, false)
	ls.RemoveListing(1)
	got := len(ls.Ls)
	want := 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
