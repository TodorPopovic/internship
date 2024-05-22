package service

import (
	"github.com/google/wire"
)

var Providers = wire.NewSet(
	NewUserService, wire.Bind(new(IUserService), new(*UserService)),
)
