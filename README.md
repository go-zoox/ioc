# Container - Simple Dependency Injection Container

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/container)](https://pkg.go.dev/github.com/go-zoox/container)
[![Build Status](https://github.com/go-zoox/container/actions/workflows/ci.yml/badge.svg?branch=master)](https://github.com/go-zoox/container/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/container)](https://goreportcard.com/report/github.com/go-zoox/container)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/container/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/container?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/container.svg)](https://github.com/go-zoox/container/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/container.svg?label=Release)](https://github.com/go-zoox/container/tags)

## Installation
To install the package, run:
```bash
go get github.com/go-zoox/container
```

## Getting Started

```go
import (
  "testing"
  "github.com/go-zoox/container"
)

func main(t *testing.T) {
	container.Register("config", Config.New())
	container.Register("logger",	Logger.New())
	// container.Register("database", Database.New())
	container.Register("service.user", UserService.New())
	// container.Register("service.post", PostService.New())
}
```

```go
// config/config.go
package config

import (
	cfg "github.com/go-zoox/config"
)

type Config struct {
	Database struct {
		Host string
		Port int
	}
	Logger struct {
		Level string
	}
}

func New() *Config {
	var c Config
	if err := cfg.Load(&c); err != nil {
		panic(err)
	}
	return &c
}
```

```go
// logger/logger.go
package logger

import (
	log "github.com/go-zoox/logger"
)

func New() *log.Logger {
	return log.New()
}
```

```go
// service/user.go
package service

type UserService struct {
	Config *Config
	Logger *log.Logger
}


func New() *UserService {
	return &UserService{
		Config: container.MustGet("config").(*Config),
		Logger: container.MustGet("logger").(*log.Logger),
	}
}

func (u *UserService) GetUser(id int) (*User, error) {
	u.Logger.Info("GetUser")

	return &User{
		ID: id,
		Name: "John Doe",
	}, nil
}
```

## Inspired by
* [vardius/gocontainer](https://github.com/vardius/gocontainer) - Simple Dependency Injection Container
* [golobby/container](https://github.com/golobby/container) - A lightweight yet powerful IoC dependency injection container for the Go programming language
* [goava/di](https://github.com/goava/di) - 🛠 A full-featured dependency injection container for go programming language
* [goioc/di](https://github.com/goioc/di) - Simple and yet powerful Dependency Injection for Go


## License
GoZoox is released under the [MIT License](./LICENSE).
