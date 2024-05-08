package authority

import (
	"fmt"

	"github.com/goravel/framework/contracts/config"
	"github.com/harranali/authority"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var Driver = &Authority{}

type Authority struct {
	//config   config.Config
	instance *authority.Authority
}

func NewAuthority(config config.Config) (*Authority, error) {
	conn := config.GetString("database.default")
	if conn == "" {
		return nil, nil
	}

	host := config.GetString(fmt.Sprintf("database.connections.%s.host", conn))
	port := config.GetString(fmt.Sprintf("database.connections.%s.port", conn))
	database := config.GetString(fmt.Sprintf("database.connections.%s.database", conn))
	username := config.GetString(fmt.Sprintf("database.connections.%s.username", conn))
	password := config.GetString(fmt.Sprintf("database.connections.%s.password", conn))
	if host == "" || port == "" || database == "" ||
		username == "" || password == "" {
		return nil, nil
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)", username, password, host, port)
	dsn += fmt.Sprintf("/%s?charset=utf8mb4&parseTime=True&loc=Local", database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db, err := facades.Orm().Connection(conn).DB()
	if err != nil {
		return nil, errors.WithMessage(err, "init database configuration error")
	}
	aliveDB, _ := db.DB()
	pingErr := aliveDB.Ping()
	if pingErr != nil {
		return nil, errors.WithMessage(err, "init connection error")
	}

	auth := authority.New(authority.Options{
		TablesPrefix: config.GetString("authority.TablesPrefix"),
		DB:           db,
	})
	return &Authority{
		instance: auth,
	}, nil
}

// Add a new role to the database it accepts the Role struct as a parameter it
// returns an error in case of any it returns an error if the role is
// already exists
func (a *Authority) CreateRole(r authority.Role) error {
	err := a.instance.CreateRole(r)
	return err
}

// Add a new permission to the database it accepts the Permission struct
// as a parameter it returns an error in case of any it returns an error if
// the permission is already exists
func (a *Authority) CreatePermission(p authority.Permission) error {
	err := a.instance.CreatePermission(p)
	return err
}

// Assigns a group of permissions to a given role it accepts the role
// slug as the first parameter, the second parameter is a slice of permission
// slugs (strings) to be assigned to the role. It returns an error in case
// of any. It returns an error in case the role does not exists. It returns an
// error in case any of the permissions does not exists. It returns an error
// in case any of the permissions is already assigned
func (a *Authority) AssignPermissionsToRole(
	roleSlug string, permSlugs []string) error {

	err := a.instance.AssignPermissionsToRole(roleSlug, permSlugs)
	return err
}

// Assigns a role to a given user it accepts the user id as the first parameter
// the second parameter the role slug. It returns an error in case of any.
// It returns an error in case the role does not exists. It returns an error
// in case the role is already assigned
func (a *Authority) AssignRoleToUser(
	userID interface{}, roleSlug string) error {

	err := a.instance.AssignRoleToUser(userID, roleSlug)
	return err
}

// Checks if a role is assigned to a user it accepts the user id as the first
// parameter the second parameter the role slug.  It returns two parameters
// the first parameter of the return is a boolean represents whether the role
// is assigned or not. The second is an error in case of any.
// In case the role does not exists, an error is returned
func (a *Authority) CheckUserRole(
	userID interface{}, roleSlug string) (bool, error) {

	ok, err := a.instance.CheckUserRole(userID, roleSlug)
	return ok, err
}

// Checks if a permission is assigned to a user it accepts in the user id as
// the first parameter, the second parameter the role slug. It returns two
// parameters: the first parameter of the return is a boolean represents
// whether the role is assigned or not, the second is an error in case of any
// In case the role does not exists, an error is returned
func (a *Authority) CheckUserPermission(
	userID interface{}, roleSlug string) (bool, error) {

	ok, err := a.instance.CheckUserPermission(userID, roleSlug)
	return ok, err
}

// Checks if a permission is assigned to a role, it accepts in the role slug
// as the first parameter. The second parameter, the permission slug.
// It returns two parameters: the first parameter of the return is a boolean
// represents whether the permission is assigned or not; the second is an error
// in case of any. In case the role does not exists, an error is returned.
// in case the permission does not exists, an error is returned
func (a *Authority) CheckRolePermission(
	roleSlug, permSlug string) (bool, error) {

	ok, err := a.instance.CheckRolePermission(roleSlug, permSlug)
	return ok, err
}
