package config

import "devops-platform/internal/common/config/internal/domain"

const (
	EnvDev     = "dev"
	EnvTesting = "testing"
	EnvUat     = "uat"
	EnvProd    = "prod"
)

//注册bean
const (
	BeanDatabase = domain.BeanDatabase
	BeanLog      = domain.BeanLog
	BeanApp      = domain.BeanApp
	BeanTekton   = domain.BeanTekton
)
