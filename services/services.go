package services

import (
	"log/slog"

	"github.com/makehlv/kk/clients"
	"github.com/makehlv/kk/config"
	"github.com/makehlv/kk/repositories"
	"github.com/makehlv/kk/services/code"
	"github.com/makehlv/kk/services/swagger"
	"github.com/makehlv/kk/services/variable"
)

type Services struct {
	CodeFlowManage *code.CodeFlowManageService
	Swagger        *swagger.SwaggerService
	Variable       *variable.VariableService
}

func NewServices(clients *clients.Clients, logger *slog.Logger, config *config.Config, repos *repositories.Repositories) *Services {
	return &Services{
		CodeFlowManage: code.NewCodeFlowManageService(clients, logger, config, repos),
		Swagger:        swagger.NewSwaggerService(clients, logger, config, repos),
		Variable:       variable.NewVariableService(clients, logger, config, repos),
	}
}
