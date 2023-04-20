package backen_test

import (
	backen "CEN3031_Project-main/actors-housing/backend/Structs"
	"testing"
)

func TestAddMember(t *testing.T) {
	ml := backen.MemberList{}

	// Add a new member
	ml.AddMember("uid123", "John Doe", "123-456-7890", "johndoe@example.com", "password", true, false)

	// Verify that the member was added
	if len(ml.Ml) != 1 {
		t.Errorf("expected 1 member, but got %d", len(ml.Ml))
	}
}

func TestSearchMember(t *testing.T) {
	ml := backen.MemberList{}

	// Add a new member
	ml.AddMember("uid123", "John Doe", "123-456-7890", "johndoe@example.com", "password", true, false)

	// Search for the added member
	member, err := ml.SearchMember("uid123")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Verify that the correct member was found
	if member.Name != "John Doe" {
		t.Errorf("expected name %q, but got %q", "John Doe", member.Name)
	}
}

func TestRemoveMember(t *testing.T) {
	ml := backen.MemberList{}

	// Add a new member
	ml.AddMember("uid123", "John Doe", "123-456-7890", "johndoe@example.com", "password", true, false)

	// Remove the added member
	ml.RemoveMember("uid123")

	// Verify that the member was removed
	if len(ml.Ml) != 0 {
		t.Errorf("expected 0 members, but got %d", len(ml.Ml))
	}
}
