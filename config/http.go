package config

import (
	fiberfacades "github.com/goravel/fiber/facades"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("http", map[string]any{
		// HTTP Driver
		"default": "fiber",
		// HTTP Drivers
		"drivers": map[string]any{
			"fiber": map[string]any{
				// prefork mode, see https://docs.gofiber.io/api/fiber/#config
				"prefork": false,
				// Optional, default is 4096 KB
				"body_limit": 1024 * 1024 * 4,
				"route": func() (route.Route, error) {
					return fiberfacades.Route("fiber"), nil
				},
			},
		},
		// HTTP URL
		"url": "http://localhost",
		// HTTP Host
		"host": "",
		// HTTP Port
		"port": config.Env("APP_PORT", "8888"),
		// HTTP Entrance
		"entrance": config.Env("APP_ENTRANCE", "/"),
		// HTTPS Configuration
		"tls": map[string]any{
			// HTTPS Host
			"host": "",
			// HTTPS Port
			"port": config.Env("APP_PORT", "8888"),
			// SSL Certificate
			"ssl": map[string]any{
				// ca.pem
				"cert": config.Env("APP_SSL_CERT", "/www/panel/storage/ssl.crt"),
				// ca.key
				"key": config.Env("APP_SSL_KEY", "/www/panel/storage/ssl.key"),
			},
		},
	})
}
