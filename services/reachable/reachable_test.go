package reachable

import (
	"context"
	"oj/api"
	"oj/db"
	"oj/services/family"
	"testing"

	"golang.org/x/exp/slices"
)

func TestReachableKids(t *testing.T) {
	queries := api.New(db.DB)
	ctx := context.TODO()

	alice, _ := queries.CreateParent(ctx, api.CreateParentParams{Username: "alice"})
	alicejr, _ := family.CreateKid(ctx, alice.ID, "alice jr")

	bob, _ := queries.CreateParent(ctx, api.CreateParentParams{Username: "bob"})
	bobjr, _ := family.CreateKid(ctx, bob.ID, "bob jr")

	carol, _ := queries.CreateParent(ctx, api.CreateParentParams{Username: "carol"})
	caroljr, _ := family.CreateKid(ctx, carol.ID, "carol jr")

	// connect parents alice with bob
	queries.CreateFriend(ctx, api.CreateFriendParams{AID: alice.ID, BID: bob.ID, BRole: "friend"})
	queries.CreateFriend(ctx, api.CreateFriendParams{AID: bob.ID, BID: alice.ID, BRole: "friend"})

	// aj can reach only bj
	connections, err := ReachableKids(ctx, alicejr.ID)
	if err != nil {
		t.Fatalf("error %s", err)
	}
	if !slices.ContainsFunc(connections, func(c api.GetConnectionRow) bool { return c.Username == "bob jr" }) {
		t.Errorf("expected bob jr to be reachable from alice jr")
	}
	if len(connections) != 1 {
		t.Errorf("expected connections to be exactly 1")
	}

	// bj can reach only aj
	connections, err = ReachableKids(ctx, bobjr.ID)
	if err != nil {
		t.Fatalf("error %s", err)
	}
	if !slices.ContainsFunc(connections, func(c api.GetConnectionRow) bool { return c.Username == "alice jr" }) {
		t.Errorf("expected alice jr to be reachable from bob jr")
	}
	if len(connections) != 1 {
		t.Errorf("expected connections to be exactly 1")
	}

	// cj can reach no one
	connections, err = ReachableKids(ctx, caroljr.ID)
	if err != nil {
		t.Fatalf("error %s", err)
	}
	if len(connections) != 0 {
		t.Errorf("expected connections to be exactly 0")
	}
}
