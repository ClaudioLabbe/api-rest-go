package routes

import (
	"api-rest-go/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(g *gin.Engine) {
	album := g.Group("/album")
	const albumById string = "/:id"

	album.GET("/", controllers.GetAllAlbum)
	album.POST("/", controllers.PostAlbum)
	album.GET(albumById, controllers.GetAlbumById)
	album.PATCH(albumById, controllers.UpdateAlbumById)
	album.DELETE(albumById, controllers.DeleteAlbum)
}
