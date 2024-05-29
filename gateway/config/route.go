package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Route struct {
	Upstream []string // 路由对应的上游服务
	Jwt      bool     // 是否需要JWT验证
	Count    uint64   // 计数器，用来实现轮询：count % len(Upstream)
}

func routeInit() map[string]*Route {
	var routeMap = make(map[string]*Route)

	var ymlRoute = viper.Get("route")
	routeArray := ymlRoute.([]interface{})
	for _, routeInterface := range routeArray {
		if route, ok := routeInterface.(map[string]interface{}); ok {
			var r Route // 创建一个Route对象

			// 获取path，作为map的key
			routePath := route["path"].(string)

			// 获取upstream数组
			if routeUpstream, ok := route["upstream"].([]interface{}); ok {
				var upstream []string
				for _, upstreamInterface := range routeUpstream {
					upstream = append(upstream, upstreamInterface.(string))
				}
				r.Upstream = upstream
			} else {
				fmt.Printf("the format of %s'upstream is invalid\n", routePath)
			}

			// 获取jwt（可选字段）
			routeJwt := route["jwt"]
			if routeJwt == nil { // 如果jwt字段不存在，直接跳过
				r.Jwt = false
			} else {
				if jwt, ok := routeJwt.(bool); ok {
					r.Jwt = jwt
				} else {
					fmt.Printf("the jwt of %s must be bollean\n", routePath)
				}
			}

			// 将Route对象存入map
			routeMap[routePath] = &r
		} else {
			println("route must be an array of map[string]interface{}")
		}
	}

	return routeMap
}
