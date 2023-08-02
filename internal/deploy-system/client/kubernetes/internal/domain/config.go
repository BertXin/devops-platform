package domain

import (
	"encoding/base64"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

type Config interface {
	GetSystemName() string
}

type KubernetesConfig interface {
	GetAddress() string
	GetCertificateAuthority() string
	GetUserClientCertificate() string
	GetUserClientKey() string
}

func NewConfig(kubeConfig KubernetesConfig) (*rest.Config, error) {
	c := &config{
		clusterServer:               kubeConfig.GetAddress(),
		clusterCertificateAuthority: kubeConfig.GetCertificateAuthority(),
		userClientCertificate:       kubeConfig.GetUserClientCertificate(),
		userClientKey:               kubeConfig.GetUserClientKey(),
	}
	return c.GetKubeConfig()
}

type config struct {
	clusterServer               string
	clusterCertificateAuthority string
	userClientCertificate       string
	userClientKey               string
}

func (c *config) kubeConfigGetter() (*clientcmdapi.Config, error) {
	var config = clientcmdapi.NewConfig()
	config.Kind = "Config"
	config.APIVersion = "V1"
	var cluster = clientcmdapi.NewCluster()
	cluster.Server = c.clusterServer
	if c.clusterCertificateAuthority == "" {
		cluster.InsecureSkipTLSVerify = true
	} else {
		cluster.CertificateAuthorityData, _ = base64.StdEncoding.DecodeString(c.clusterCertificateAuthority)
	}

	config.Clusters["kubernetes"] = cluster

	var context = clientcmdapi.NewContext()
	context.Cluster = "kubernetes"
	context.AuthInfo = "default"

	config.Contexts["default@kubernetes"] = context

	config.CurrentContext = "default@kubernetes"

	var user = clientcmdapi.NewAuthInfo()

	user.ClientCertificateData, _ = base64.StdEncoding.DecodeString(c.userClientCertificate)

	user.ClientKeyData, _ = base64.StdEncoding.DecodeString(c.userClientKey)

	config.AuthInfos["default"] = user

	return config, nil
}

func (c *config) GetKubeConfig() (*rest.Config, error) {
	return clientcmd.BuildConfigFromKubeconfigGetter("", c.kubeConfigGetter)
}
