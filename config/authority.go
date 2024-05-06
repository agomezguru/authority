package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("authority", map[string]any{
		"name": "Package",
	})
}
