package app

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// ProjectConsent is a simple struct that enables the json marshalling and unmarshalling
// as described by the API
type ProjectConsent struct {
	ProjectApplicationID int  `json:"project_application_id"`
	GeneticConsent       bool `json:"genetic_consent"`
	ClinicalConsent      bool `json:"clinical_consent"`
}

// DefaultConsent is a simple struct that enables the json marshalling and unmarshalling
// as described by the API
type DefaultConsent struct {
	GeneticConsent  bool `json:"genetic_consent"`
	ClinicalConsent bool `json:"clinical_consent"`
}

// Consents is a simple struct that enables the json marshalling and unmarshalling
// as described by the API
type Consents struct {
	DefaultConsent  DefaultConsent   `json:"default_consent"`
	ProjectConsents []ProjectConsent `json:"project_consents"`
}

// ParticipantsRequestHeader is used by gin to
// bind header information into.
type ParticipantsRequestHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

// ConsentsGet is the API Handler for responding to GET /consents requests
// from participants. Extracts the Bearer token and JWT from the
// Authorization header, decodes and verifies the JWT, and returns
// a canned consent payload. TODO: This endpoint should
// lookup the admin from the KeyFile and request the consents from
// the consents service
func ConsentsGet(ctx *gin.Context) {
	// Extract Bearer token
	jwt, err := extractBearerToken(ctx)
	if err != nil {
		renderError(
			http.StatusUnauthorized,
			err,
			ctx,
		)
		return
	}

	// Decode Bearer Token
	_, err = extractAdminID(jwt)
	if err != nil {
		err = fmt.Errorf("Unable to successfully decode JWT: %s", err)
		renderError(
			http.StatusUnauthorized,
			err,
			ctx,
		)
		return
	}

	// Build canned consents
	projectConsents := []ProjectConsent{
		{
			ProjectApplicationID: 1,
			GeneticConsent:       false,
			ClinicalConsent:      false,
		},
		{
			ProjectApplicationID: 2,
			GeneticConsent:       true,
			ClinicalConsent:      true,
		},
	}

	ctx.JSON(
		http.StatusOK,
		Consents{
			DefaultConsent: DefaultConsent{
				GeneticConsent:  true,
				ClinicalConsent: true,
			},
			ProjectConsents: projectConsents,
		},
	)
}

func renderError(statusCode int, err error, ctx *gin.Context) {
	ctx.JSON(statusCode, gin.H{
		"error": gin.H{
			"status":  statusCode,
			"message": err.Error(),
		},
	})
}

func extractBearerToken(ctx *gin.Context) (*string, error) {
	h := ParticipantsRequestHeader{}

	err := ctx.ShouldBindHeader(&h)
	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile(`^(?P<Bearer>Bearer)\s{1}(?P<JWT>.+)$`)
	match := r.MatchString(h.AuthorizationHeader)
	if match != true {
		return nil, errors.New("Authorization header must contain a Bearer token")
	}

	matches := r.FindSubmatch([]byte(h.AuthorizationHeader))

	token := string(matches[2])
	return &token, nil
}
