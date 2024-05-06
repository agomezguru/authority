package facades

import (
	"log"

	authority "github.com/agomezguru/authority"
	"github.com/agomezguru/authority/contracts"
)

func Authority() contracts.Authority {
	instance, err := authority.App.Make(authority.Binding)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return instance.(contracts.Authority)
}
