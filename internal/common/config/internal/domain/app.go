package domain

type app struct {
	Url      string
	FrontUrl string
	Addr     string
	Env      string
	Gods     []string
}

func (cfg *app) GetServerAddress() string {
	return cfg.Addr
}
func (cfg *app) GetEnv() string {
	return cfg.Env
}
