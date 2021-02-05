package app

import (
	"encoding/json"
	"os"

	jose "github.com/dvsekhvalnov/jose2go"
)

func ExtractAdminID(jwt *string) (*string, error) {
	sharedKey := os.Getenv("SHARED_KEY")
	p, _, err := jose.Decode(*jwt, []byte(sharedKey))

	if err != nil {
		return nil, err
	}

	payload := Payload{}

	err = json.Unmarshal([]byte(p), &payload)
	if err != nil {
		return nil, err
	}

	return &payload.Sub, nil
}

type Payload struct {
	Sub string `json:"sub"`
}
