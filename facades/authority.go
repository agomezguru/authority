package facades

import (
	"log"

	authority "github.com/agomezguru/authority"
)

func Authority() error {
	instance, err := authority.App.Make(authority.Binding)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return instance.(error)
}
