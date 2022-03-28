package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"

	configs2 "github.com/AzamatAbdranbayev/deployer/internal/config"
	"github.com/AzamatAbdranbayev/deployer/internal/errors_code"
	"github.com/AzamatAbdranbayev/deployer/internal/model"
	"github.com/AzamatAbdranbayev/deployer/internal/response"
)

type StaticService interface {
	DeployStatic(ctx context.Context, reqBody model.PayloadModel) *response.Response
	SendTelegramNotification(ctx context.Context, message string)
}

type StaticServiceImpl struct {
	cfg *configs2.Config
	cli *http.Client
}

func NewStaticService(cfg *configs2.Config, cli *http.Client) StaticService {
	return &StaticServiceImpl{cfg: cfg, cli: cli}
}

func (d StaticServiceImpl) DeployStatic(ctx context.Context, reqBody model.PayloadModel) *response.Response {
	resp := response.InitResp()

	//
	pos := strings.LastIndex(reqBody.Repository.FullName, "/")

	if pos == -1 {
		resp.SetError(errors_code.GetErr(101))
		return resp
	}
	adjustedPos := pos + len("/")
	if adjustedPos >= len(reqBody.Repository.FullName) {
		// do
	}
	repoName := reqBody.Repository.FullName[adjustedPos:len(reqBody.Repository.FullName)]
	//

	cmdName := "git"
	cmdArgs := []string{"pull"}
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Dir = fmt.Sprintf("%s/%s", "/var/www", repoName)

	out, err := cmd.Output()
	if err != nil {
		resp.SetError(errors_code.GetErr(102))
		resp.SetError(0, string(err.Error()))
		d.SendTelegramNotification(ctx, err.Error())
		return resp
	}
	d.SendTelegramNotification(ctx, fmt.Sprintf("repo:%s - status:success", repoName))
	resp.SetValue(string(out))

	return resp
}

func (d StaticServiceImpl) SendTelegramNotification(ctx context.Context, message string) {

	url := fmt.Sprintf("%s='%s'", "url telegram bot", strings.Replace(message, "/", "-", 0))
	req, errReq := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if errReq != nil {
		// do
	}
	respTeleg, respTelegReq := d.cli.Do(req)
	if respTelegReq != nil {
		// do
	}
	bodyResp, errBody := io.ReadAll(respTeleg.Body)
	if errBody != nil {
		// do
	}
	if respTeleg.StatusCode >= http.StatusMultipleChoices {
		er1 := errors.New(string(bodyResp))
		// do
	}
	defer respTeleg.Body.Close()
}
