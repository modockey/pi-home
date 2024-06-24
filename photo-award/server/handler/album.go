package handler

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/modockey/pi-home/photo-award/gen/restapi/operations/album"
)

func GetAlbumsHandler(params album.GetAlbumsParams, user interface{}) middleware.Responder {
	log.Printf("%s\n", user)
	return middleware.NotImplemented("todo")
}

func PostAlbumsHanlder(params album.PostAlbumsParams, user interface{}) middleware.Responder {
	return middleware.NotImplemented("operation album.PostAlbumsHanlder has not yet been implemented")
}

func GetAlbumsByIDHanlder(params album.GetAlbumsByID, user interface{}) middleware.Responder {
	return middleware.NotImplemented("operation album.GetAlbumsByIDHanlder has not yet been implemented")
}

func PostAlbumsByIdHanlder(params album.PostPhoto, user interface{}) middleware.Responder {
	return middleware.NotImplemented("operation album.PostAlbumsByIdHanlder has not yet been implemented")
}
