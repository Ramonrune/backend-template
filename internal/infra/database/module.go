package database

import (
	"github.com/ramonrune/assina-backend/internal/infra/database/user"
	"go.uber.org/fx"
)

var Module = fx.Options(
	user.Module,
)
