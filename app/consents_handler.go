package app

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type ProjectConsent struct {
	ProjectApplicationID int  `json:"project_application_id"`
	GeneticConsent       bool `json:"genetic_consent"`
	ClinicalConsent      bool `json:"clinical_consent"`
}

type DefaultConsent struct {
	GeneticConsent  bool `json:"genetic_consent"`
	ClinicalConsent bool `json:"clinical_consent"`
}

type Consents struct {
	DefaultConsent  DefaultConsent   `json:"default_consent"`
	ProjectConsents []ProjectConsent `json:"project_consents"`
}

type ParticipantsRequestHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

func ConsentsGet(ctx *gin.Context) {
	jwt, err := ExtractBearerToken(ctx)
	if err != nil {
		renderError(
			http.StatusUnauthorized,
			err,
			ctx,
		)
		return
	}

	// Decode Bearer Token
	_, err = ExtractAdminID(jwt)
	if err != nil {
		err = fmt.Errorf("Unable to properly decode JWT: %s", err)
		renderError(
			http.StatusUnauthorized,
			err,
			ctx,
		)
		return
	}

	// Return Consents
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

func ExtractBearerToken(ctx *gin.Context) (*string, error) {
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
