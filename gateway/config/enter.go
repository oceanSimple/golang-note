package config

import "github.com/spf13/viper"

const (
	filePath = "./gateway/config.yaml"
)

func init() {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func ViperInit() (upstreamMap map[string]string, routeMap map[string]*Route) {
	return upstreamInit(), routeInit()
}
