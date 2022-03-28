package handler

import service2 "github.com/AzamatAbdranbayev/deployer/internal/service"

type StaticHandlerOption func(h *StaticHandler)

func WithDeployStatic(service service2.StaticService) StaticHandlerOption {
	return func(h *StaticHandler) {
		h.service = service
	}
}
