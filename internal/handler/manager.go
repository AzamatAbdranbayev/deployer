package handler

import (
	"github.com/AzamatAbdranbayev/deployer/internal/service"
)

type Manager struct {
	Static StaticHandler
}

func NewManager(service *service.Manager) *Manager {
	st := NewStaticHandler(WithDeployStatic(service.Static))
	return &Manager{
		Static: *st,
	}
}
