package code

import (
	"log/slog"

	"github.com/makehlv/kk/clients"
	"github.com/makehlv/kk/config"
	"github.com/makehlv/kk/repositories"
)

type CodeFlowManageService struct {
	logger       *slog.Logger
	clients      *clients.Clients
	config       *config.Config
	repositories *repositories.Repositories
}

func NewCodeFlowManageService(
	clients *clients.Clients, logger *slog.Logger, config *config.Config, repositories *repositories.Repositories) *CodeFlowManageService {
	return &CodeFlowManageService{clients: clients, logger: logger, config: config, repositories: repositories}
}
