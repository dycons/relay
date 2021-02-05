package app

import (
	"net/http"

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

func ConsentsGet(ctx *gin.Context) {
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
