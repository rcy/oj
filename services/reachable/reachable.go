package reachable

import (
	"context"
	"log"
	"oj/api"
	"oj/db"
)

func ReachableKids(ctx context.Context, kidID int64) ([]api.GetConnectionRow, error) {
	queries := api.New(db.DB)

	// get possible friend connections
	// find all the kids of all the friends of this kid's parenst

	reachableKids := map[int64]any{}

	parents, err := queries.GetParents(ctx, kidID)
	if err != nil {
		return nil, err
	}
	for _, parent := range parents {
		kids, err := queries.GetKids(ctx, parent.ID)
		if err != nil {
			return nil, err
		}
		for _, kids := range kids {
			reachableKids[kids.ID] = true
		}

		friends, err := queries.GetFriends(ctx, parent.ID)
		if err != nil {
			return nil, err
		}
		for _, friend := range friends {
			kids, err := queries.GetKids(ctx, friend.ID)
			if err != nil {
				return nil, err
			}
			for _, kids := range kids {
				reachableKids[kids.ID] = true
			}
		}
	}

	// remove self
	delete(reachableKids, kidID)

	var connections []api.GetConnectionRow
	for kidID := range reachableKids {
		connection, err := queries.GetConnection(ctx, api.GetConnectionParams{AID: kidID, ID: kidID})
		if err != nil {
			return nil, err
		}
		log.Printf("%v", connection)
		connections = append(connections, connection)
	}

	return connections, nil
}
