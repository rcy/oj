package background

import (
	"context"
	"database/sql"
	"errors"
	"oj/api"
	"oj/db"
	"oj/element/gradient"
)

func ForUser(ctx context.Context, userID int64) (*gradient.Gradient, error) {
	queries := api.New(db.DB)

	gradientRow, err := queries.UserGradient(ctx, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return &gradient.Default, nil
	}
	if err != nil {
		return nil, err
	}

	return &gradientRow.Gradient, nil
}
