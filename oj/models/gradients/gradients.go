package gradients

import (
	"database/sql"
	"encoding/json"
	"oj/db"
	"oj/element/gradient"
)

func UserBackground(userID int64) (gradient.Gradient, error) {
	var encodedGradient []byte
	var backgroundGradient gradient.Gradient

	err := db.DB.Get(&encodedGradient, "select gradient from gradients order by created_at desc limit 1")
	if err != nil {
		if err == sql.ErrNoRows {
			return gradient.Default, nil
		}
		return gradient.Gradient{}, err
	}
	err = json.Unmarshal(encodedGradient, &backgroundGradient)
	if err != nil {
		return gradient.Gradient{}, err
	}
	return backgroundGradient, err
}
