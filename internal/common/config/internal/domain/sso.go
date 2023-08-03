package domain

type sso struct {
	ClientId     string `toml:"client_id"`
	ClientSecret string `toml:"client_secret"`
	AuthUrl      string `toml:"auth_url"`
	Realm        string `toml:"realm"`
}

func (s *sso) GetClientId() string {
	return s.ClientId
}

func (s *sso) GetClientSecret() string {
	return s.ClientSecret
}

func (s *sso) GetAuthUrl() string {
	return s.AuthUrl
}

func (s *sso) GetRealm() string {
	return s.Realm
}
