package config

import (
	"os"

	"github.com/golang/glog"
	"gopkg.in/yaml.v3"
)

type Config struct {
	SrcRegistry         string `yaml:"srcRegistry"`
	SrcRegisrtyUsername string `yaml:"srcRegisrtyUsername"`
	SrcRegisrtyPassword string `yaml:"srcRegisrtyPassword"`

	DestRegistry           string `yaml:"destRegistry"`
	DestRegistryKubeconfig string `yaml:"destRegistryKubeconfig"`

	Images map[string][]string `yaml:"images"`
}

// format as
/*
srcRegistry: docker.io
srcRegisrtyUsername: lingdie
srcRegisrtyPassword: password

destRegistry: hub.sealos.cn
destRegistryKubeconfig: kubeconfig.yaml

images:
	labring/redis: [ "3.1.4" ] # repo name should be same in both src and dest.
*/

func NewConfig(filePath string) Config {
	var config Config
	source, err := os.ReadFile(filePath)
	if err != nil {
		glog.Errorln(err)
		return config
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		glog.Errorln(err)
		return config
	}
	return config
}
