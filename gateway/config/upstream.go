package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func upstreamInit() map[string]string {
	var upstreamMap = make(map[string]string)
	upstreamInterface := viper.Get("upstream")
	if upstream, ok := upstreamInterface.([]interface{}); ok {
		for _, serviceInterface := range upstream {
			if service, ok := serviceInterface.(map[string]interface{}); ok {
				serviceName := service["name"].(string)
				serviceURL := service["url"].(string)
				upstreamMap[serviceName] = serviceURL
			} else {
				fmt.Println("Cannot convert to map[string]interface{}")
				return nil
			}
		}
	} else {
		fmt.Println("Cannot convert to []interface{}")
	}
	return upstreamMap
}
