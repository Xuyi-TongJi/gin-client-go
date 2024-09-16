package config

import (
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/klog"
	"os"
	"sync"
)

var (
	once   sync.Once
	KeyMap map[KeyName]string
)

type Config struct {
	Server Server
}

type Server struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func initConfig() {
	once.Do(func() {
		var config Config
		yamlFile, err := os.ReadFile("./.gin-client-go.yaml")
		if err != nil {
			klog.Fatal(err)
			return
		}
		err = yaml.Unmarshal(yamlFile, &config)
		if err != nil {
			klog.Fatal(err)
		}
		KeyMap = make(map[KeyName]string)
		KeyMap[ServerName] = config.Server.Name
		KeyMap[ServerHost] = config.Server.Host
		KeyMap[ServerPort] = config.Server.Port
	})
}

func GetString(keyName KeyName) string {
	return KeyMap[keyName]
}
