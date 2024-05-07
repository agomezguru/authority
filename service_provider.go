package authority

import (
	"github.com/goravel/framework/contracts/foundation"
)

const Binding = "hello"

var App foundation.Application

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	App = app

	app.Bind(Binding, func(app foundation.Application) (any, error) {
		return NewAuthority(app.MakeConfig())
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	app.Publishes("github.com/agomezguru/authority", map[string]string{
		"config/authority.go": app.ConfigPath("authority.go"),
	})
}
