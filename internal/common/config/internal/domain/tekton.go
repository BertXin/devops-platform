package domain

type tekton struct {
	SystemName string `toml:"system_name"`
}

func (cfg *tekton) GetSystemName() string {
	return cfg.SystemName
}
