# An Example For A Goravel Extend Package

## Directory Structure

This is a directory standard, but you can change it if you like.

| Directory        | Action           |
| -----------      | --------------   |
| commands         | Store the command files   |
| config            | Store the config files   |
| contracts        | Store the contract files   |
| facades          | Store the facade files   |
| root             | Store the service provider and package source code   |

## Install

1. Add package

```
go get -u github.com/agomezguru/authority
```

2. Register service provider

```
// config/app.go
import authority "github.com/agomezguru/authority"

"providers": []foundation.ServiceProvider{
    ...
    &authority.ServiceProvider{},
}
```

3. Publish Configuration

```
go run . artisan vendor:publish --package=github.com/agomezguru/authority
```

4. Testing

```
// main.go
import authorityfacades "github.com/agomezguru/authority/facades"

fmt.Println(authorityfacades.Authority().World())
```

The console will print `Welcome To Goravel Package`.
