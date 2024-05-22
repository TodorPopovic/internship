package db

import "github.com/google/wire"

var Providers = wire.NewSet(
	NewMyRepo,
	wire.Bind(new(IPostgresRepo), new(*MyRepo)),
)
