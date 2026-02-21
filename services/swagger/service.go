package swagger

import (
	"log/slog"

	"github.com/makehlv/kk/clients"
	"github.com/makehlv/kk/config"
	"github.com/makehlv/kk/repositories"
)

type SwaggerService struct {
	logger  *slog.Logger
	clients *clients.Clients
	config  *config.Config

	repositories *repositories.Repositories
}

func NewSwaggerService(clients *clients.Clients, logger *slog.Logger, config *config.Config, repositories *repositories.Repositories) *SwaggerService {
	return &SwaggerService{clients: clients, logger: logger, config: config, repositories: repositories}
}
