package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-note/gateway/config"
	"golang-note/gateway/middleware"
	"net/http/httputil"
	"net/url"
)

var (
	upstreamMap map[string]string
	routeMap    map[string]*config.Route
	server      *gin.Engine
)

func main() {
	// 读取配置文件，初始化上游服务和路由
	upstreamMap, routeMap = config.ViperInit()
	// 初始化gin框架
	server = gin.Default()
	// 根据配置文件中的路由信息，初始化路由
	routeInit()

	fmt.Println("\u001B[32mGateway is running on :8000\u001B[0m")
	server.Run(":8000")
}

func routeInit() {
	for path, info := range routeMap {
		if info.Jwt {
			jwtProxy(path, info)
		} else {
			defaultProxy(path, info)
		}
	}
}

// 反向代理：默认路由
func defaultProxy(path string, info *config.Route) {
	server.Any(path+"/*any", func(c *gin.Context) {
		// 获取上游服务的URL，并解析
		// 使用了轮询算法，将请求分发到不同的上游服务
		target, err := url.Parse(upstreamMap[info.Upstream[info.Count%uint64(len(info.Upstream))]])
		if err != nil {
			fmt.Println(err)
			return
		}
		info.Count++

		// 配置反向代理
		target.Scheme = "http"
		proxy := httputil.NewSingleHostReverseProxy(target)
		req := c.Request
		req.Host = target.Host
		req.URL.Host = target.Host
		req.URL.Scheme = target.Scheme

		// 转发请求
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}

// 反向代理：启用JWT认证的路由
func jwtProxy(path string, info *config.Route) {
	server.Any(path+"/*any", middleware.JwtHandler(), func(c *gin.Context) {
		target, err := url.Parse(upstreamMap[info.Upstream[info.Count%uint64(len(info.Upstream))]])
		if err != nil {
			fmt.Println(err)
			return
		}
		info.Count++
		target.Scheme = "http"
		proxy := httputil.NewSingleHostReverseProxy(target)
		req := c.Request
		req.Host = target.Host
		req.URL.Host = target.Host
		req.URL.Scheme = target.Scheme
		proxy.ServeHTTP(c.Writer, c.Request)
	})
}

// 重定向：默认路由
func defaultRoute(path string, info *config.Route) {
	server.Any(path+"/*any", func(c *gin.Context) {
		// 获取路由对应的上游服务
		upstream := info.Upstream[info.Count%uint64(len(info.Upstream))]
		info.Count++

		// 获取上游服务的URL
		upstreamURL := upstreamMap[upstream]

		// 转发请求
		c.Redirect(302, upstreamURL+c.Request.RequestURI)
	})
}

// 重定向：启用JWT认证的路由
func jwtRoute(path string, info *config.Route) {
	server.Any(path+"/*any", middleware.JwtHandler(), func(c *gin.Context) {
		// 获取路由对应的上游服务
		upstream := info.Upstream[info.Count%uint64(len(info.Upstream))]
		info.Count++

		// 获取上游服务的URL
		upstreamURL := upstreamMap[upstream]

		// 转发请求
		c.Redirect(302, upstreamURL+c.Request.RequestURI)
	})
}
