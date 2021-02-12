package app

import (
	"encoding/json"
	"errors"
	"os"

	jose "github.com/dvsekhvalnov/jose2go"
)

// extractAdminID accepts a JWT and returns
// the adminID inside. It uses the jose2go
// library to unpack and verify the JWT. It then
// parses the json payload and verifies that
// the user's role is research-participant and
// has an ID present.
func extractAdminID(jwt *string) (*string, error) {
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

	if payload.Role != "research-participant" {
		return nil, errors.New("subject role must be 'research-participant'")
	}

	if payload.AdminID == "" {
		return nil, errors.New("admin id is missing")
	}

	return &payload.AdminID, nil
}

type Payload struct {
	AdminID string `json:"admin_id"`
	Role    string `json:"role"`
}
