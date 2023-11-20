package background

import (
	"context"
	"database/sql"
	"encoding/json"
	"oj/api"
	"oj/db"
	"oj/element/gradient"
)

func decodedGradient(b []byte) (*gradient.Gradient, error) {
	var gg gradient.Gradient

	err := json.Unmarshal(b, &gg)
	if err != nil {
		return nil, err
	}
	return &gg, nil
}

func ForUser(userID int64) (*gradient.Gradient, error) {
	queries := api.New(db.DB)

	gradientRow, err := queries.UserGradient(context.TODO(), userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &gradient.Default, nil
		}
		return nil, err
	}

	dg, err := decodedGradient(gradientRow.Gradient)
	if err != nil {
		return nil, err
	}

	return dg, nil
}
