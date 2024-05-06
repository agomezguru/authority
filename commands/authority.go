package commands

import (
	"fmt"

	"github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/console/command"
)

type Authority struct{}

func NewAuthority() *Authority {
	return &Authority{}
}

// Signature The name and signature of the console command.
func (receiver *Authority) Signature() string {
	return "hello"
}

// Description The console command description.
func (receiver *Authority) Description() string {
	return "Hello"
}

// Extend The console command extend.
func (receiver *Authority) Extend() command.Extend {
	return command.Extend{}
}

// Handle Execute the console command.
func (receiver *Authority) Handle(ctx console.Context) error {
	fmt.Println("Run Authority command")

	return nil
}
