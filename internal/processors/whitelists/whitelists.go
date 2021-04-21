package whitelists

import (
	"context"

	"rebrainme/gotest/internal/repositories"
)

type Validator interface {
	Validate(ctx context.Context, email string) error
}

type whitelistsProcessor struct {
	cityWhitelistValidator Validator
}

func New(cfg *Config, repoContainer repositories.Container) Validator {
	return &whitelistsProcessor{
		cityWhitelistValidator: newCityWhitelistValidator(repoContainer.Clients, cfg.CitiesValidator.GrantedCities),
	}
}

func (p *whitelistsProcessor) Validate(ctx context.Context, email string) error {
	for _, validator := range p.validators() {
		if err := validator.Validate(ctx, email); err != nil {
			return err
		}
	}

	return nil
}

func (p *whitelistsProcessor) validators() []Validator {
	return []Validator{
		p.cityWhitelistValidator,
	}
}
