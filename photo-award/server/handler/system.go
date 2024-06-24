package handler

import (
	"github.com/go-openapi/runtime/middleware"
	server "github.com/modockey/pi-home/photo-award"
	"github.com/modockey/pi-home/photo-award/gen/models"
	"github.com/modockey/pi-home/photo-award/gen/restapi/operations/system"
)

func GetSystemVersionHandler(params system.GetSystemVersionParams) middleware.Responder {
	body := models.Version{Version: server.Version}
	return system.NewGetSystemVersionOK().WithPayload(&body)
}
