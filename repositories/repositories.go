package repositories

import (
	"log/slog"

	"github.com/makehlv/kk/config"
	"github.com/makehlv/kk/repositories/swagger"
	"github.com/makehlv/kk/repositories/variable"
)

type Repositories struct {
	Swagger  *swagger.SwaggerRepository
	Variable *variable.VariableRepository
}

func NewRepositories(logger *slog.Logger, config *config.Config) *Repositories {
	return &Repositories{
		Swagger:  swagger.NewSwaggerRepository(logger, config),
		Variable: variable.NewVariableRepository(logger, config),
	}
}
