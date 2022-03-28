package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AzamatAbdranbayev/deployer/internal/errors_code"
	"github.com/AzamatAbdranbayev/deployer/internal/model"
	"github.com/AzamatAbdranbayev/deployer/internal/response"
	"github.com/AzamatAbdranbayev/deployer/internal/service"
	"github.com/labstack/echo/v4"
)

type StaticHandler struct {
	service service.StaticService
}

func NewStaticHandler(opts ...StaticHandlerOption) *StaticHandler {
	ph := &StaticHandler{}

	for _, v := range opts {
		v(ph)
	}
	return ph
}

func (h *StaticHandler) DeployStatic(c echo.Context) error {
	resp := response.InitResp()
	var reqPayload model.PayloadModel
	err := json.NewDecoder(c.Request().Body).Decode(&reqPayload)

	if err != nil {
		resp.SetError(errors_code.GetErr(103))
		h.service.SendTelegramNotification(c.Request().Context(), err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, &resp)
	}
	data := h.service.DeployStatic(c.Request().Context(), reqPayload)
	if !data.Status {
		return echo.NewHTTPError(http.StatusBadRequest, data)
	}
	return c.JSON(http.StatusOK, data)
}
