package domain

type TokenResponse struct {
	AccessToken  string
	RefreshToken string
	ExpiresIn    int64
}

type UserInfo struct {
	ID       string
	Name     string
	Username string
	Email    string
}

type LoginUser struct {
	UserInfo
	TokenResponse
}
