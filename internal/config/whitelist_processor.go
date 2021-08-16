package config

type WhitelistProcessor struct {
	CitiesValidator struct {
		GrantedCities []string `envconfig:"GRANTED_CITIES,default=[]"`
	}
}
