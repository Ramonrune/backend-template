package user

import (
	"github.com/ramonrune/assina-backend/internal/app/domain/user"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewUserRepository,
		fx.As(new(user.Repository)),
	),
)
