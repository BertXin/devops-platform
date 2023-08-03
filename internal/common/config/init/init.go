package init

import (
	"devops-platform/internal/common/config/internal/domain"
	"devops-platform/pkg/beans"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {

	configPath := os.Getenv("DEPLOY_SYSTEM_CONFIG_PATH")
	if configPath == "" {
		configPath = "config/dev.toml"
	}
	var conf domain.Config

	if _, err := toml.DecodeFile(configPath, &conf); err != nil {
		dir, _ := os.Getwd()
		logrus.Panicf("路径[%s]下配置文件[%s]加载错误:%s\n", dir, configPath, err.Error())
		return
	}

	beans.Register(domain.BeanDatabase, &conf.Database)
	beans.Register(domain.BeanLog, &conf.Log)
	beans.Register(domain.BeanApp, &conf.App)
	beans.Register(domain.BeanTekton, &conf.Tekton)
	beans.Register(domain.BeanSsoLogin, &conf.Sso)
}
