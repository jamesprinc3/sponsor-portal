package main

import (
	"log"

	"github.com/docsocsf/sponsor-portal/auth"
	"github.com/docsocsf/sponsor-portal/config"
	"github.com/docsocsf/sponsor-portal/sponsor"
)

func makeSponsorService(staticFiles string) *sponsor.Service {
	authEnvConfig, err := config.GetAuth()
	if err != nil {
		log.Fatal(err, "Make sponsor service")
	}

	authConfig := &auth.Config{
		CookieSecret: []byte(authEnvConfig.CookieSecret),

		BaseURL:      authEnvConfig.BaseURL,
		Issuer:       authEnvConfig.Issuer,
		ClientID:     authEnvConfig.ClientID,
		ClientSecret: authEnvConfig.ClientSecret,

		JwtSecret: []byte(authEnvConfig.JwtSecret),
		JwtIssuer: authEnvConfig.JwtIssuer,
	}

	service, err := sponsor.New(authConfig, staticFiles)
	if err != nil {
		log.Fatal(err)
	}

	db, err := config.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	err = service.SetupDatabase(db)
	if err != nil {
		log.Fatal(err)
	}

	return service
}
