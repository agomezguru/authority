package authority

import (
	"fmt"

	"github.com/goravel/framework/contracts/config"
)

type Authority struct {
	config config.Config
}

func NewAuthority(config config.Config) *Authority {
	return &Authority{config: config}
}

func (s *Authority) World() string {
	return fmt.Sprintf("Welcome To Goravel %s", s.config.GetString("hello.name"))
}
