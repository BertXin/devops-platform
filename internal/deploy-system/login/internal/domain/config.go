package domain

type Config interface {
	GetClientId() string
	GetClientSecret() string
	GetAuthUrl() string
	GetRealm() string
}
