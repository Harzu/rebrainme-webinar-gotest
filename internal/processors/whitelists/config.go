package whitelists

type Config struct {
	CitiesValidator struct {
		GrantedCities []string `envconfig:"default=[]"`
	}
}
