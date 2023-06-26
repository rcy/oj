package gradients

import (
	"database/sql"
	"encoding/json"
	"oj/db"
	"oj/element/gradient"
	"time"
)

type GradientRow struct {
	ID        int64
	CreatedAt time.Time
	UserID    int64
	Gradient  []byte
}

func (g GradientRow) DecodedGradient() (*gradient.Gradient, error) {
	var gg gradient.Gradient

	err := json.Unmarshal(g.Gradient, &gg)
	if err != nil {
		return nil, err
	}
	return &gg, nil
}

func UserGradient(userID int64) (*GradientRow, error) {
	var row GradientRow
	err := db.DB.Get(&row, "select gradient from gradients where user_id = ? order by created_at desc limit 1", userID)
	if err != nil {
		return nil, err
	}
	return &row, nil
}

func UserBackground(userID int64) (*gradient.Gradient, error) {
	var gradientRow *GradientRow

	gradientRow, err := UserGradient(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &gradient.Default, nil
		}
		return nil, err
	}

	dg, err := gradientRow.DecodedGradient()
	if err != nil {
		return nil, err
	}

	return dg, nil
}
