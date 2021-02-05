package app

import (
	"net/http"
	"strings"

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
	_, err := ExtracBearerToken(ctx)
	if err != nil {
		// TODO respond with 401
		panic(err)
	}

	// Decode Bearer Token

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

func ExtracBearerToken(ctx *gin.Context) (*string, error) {
	h := ParticipantsRequestHeader{}

	err := ctx.ShouldBindHeader(&h)
	if err != nil {
		return nil, err
	}

	s := strings.Fields(h.AuthorizationHeader)
	token := s[1]
	return &token, nil
}
