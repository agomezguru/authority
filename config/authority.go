package config

import (
	"github.com/goravel/framework/facades"
)

func init() {
	config := facades.Config()
	config.Add("authority", map[string]any{
		// Authority creates 4 database tables to store the roles, permissions and
		// their relationships, to initialize authority you need to pass two
		// parameters, the first one is the prefix of the table names, the second
		// one is an instance of gorm, here is how you can initiate the package
		//
		// Default database table prefix.
		"TablesPrefix": "authority_",

		// Using this prefix, the package uses gorm auto migration, so there is no
		// need to migrate or manage the database tables since it's taken care of
		// for you automatically
	})
}
