package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/modockey/pi-home/photo-award/gen/restapi/operations/album"
)

func GetAlbumsHandler(params album.GetAlbumsParams, _ interface{}) middleware.Responder {
	return middleware.NotImplemented("todo")
}

func PostAlbumsHanlder(params album.PostAlbumsParams) middleware.Responder {
	return middleware.NotImplemented("operation album.PostAlbumsHanlder has not yet been implemented")
}

func GetAlbumsByIDHanlder(params album.GetAlbumsByID) middleware.Responder {
	return middleware.NotImplemented("operation album.GetAlbumsByIDHanlder has not yet been implemented")
}

func PostAlbumsByIdHanlder(params album.PostPhoto) middleware.Responder {
	return middleware.NotImplemented("operation album.PostAlbumsByIdHanlder has not yet been implemented")
}
