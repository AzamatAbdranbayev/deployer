package service

import (
	"net/http"

	configs2 "github.com/AzamatAbdranbayev/deployer/internal/config"
)

type Manager struct {
	Static StaticService
}

func NewManager(cfg *configs2.Config, cli *http.Client) (*Manager, error) {
	staticS := NewStaticService(cfg, cli)
	return &Manager{
		Static: staticS,
	}, nil
}
