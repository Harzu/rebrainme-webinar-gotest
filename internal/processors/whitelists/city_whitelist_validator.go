package whitelists

import (
	"context"
	"errors"

	"rebrainme/gotest/internal/repositories/clients"
)

type cityWhitelistValidator struct {
	grantedCities []string
	clientRepo    clients.Repository
}

func newCityWhitelistValidator(clientRepo clients.Repository, grantedCities []string) Validator {
	return &cityWhitelistValidator{
		grantedCities: grantedCities,
		clientRepo:    clientRepo,
	}
}

func (v *cityWhitelistValidator) Validate(ctx context.Context, email string) error {
	client, err := v.clientRepo.FindClientByEmail(ctx, email)
	if err != nil {
		return err
	}

	for _, city := range v.grantedCities {
		if client.City == city {
			return nil
		}
	}

	return errors.New("client city is not granted")
}
