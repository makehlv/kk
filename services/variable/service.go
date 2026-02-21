package variable

import (
	"log/slog"

	"github.com/makehlv/kk/clients"
	"github.com/makehlv/kk/config"
	"github.com/makehlv/kk/repositories"
)

type VariableService struct {
	logger       *slog.Logger
	clients      *clients.Clients
	config       *config.Config
	repositories *repositories.Repositories
}

func NewVariableService(
	clients *clients.Clients, logger *slog.Logger, config *config.Config, repositories *repositories.Repositories) *VariableService {
	return &VariableService{clients: clients, logger: logger, config: config, repositories: repositories}
}
