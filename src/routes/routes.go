package routes

import (
	"api-rest-go/src/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(g *gin.Engine) {
	album := g.Group("/album")
	const byId string = "/:id"

	album.GET("/", controllers.GetAllAlbum)
	album.POST("/", controllers.PostAlbum)
	album.GET(byId, controllers.GetAlbumById)
	album.PATCH(byId, controllers.UpdateAlbumById)
	album.DELETE(byId, controllers.DeleteAlbum)

	rol := g.Group("/rol")

	rol.GET("/", controllers.GetRols)
	rol.POST("/", controllers.PostRol)
	rol.GET(byId, controllers.GetRolById)
	rol.DELETE(byId, controllers.DeleteRol)

	user := g.Group("/user")

	user.GET("/", controllers.GetUsers)
	user.POST("/", controllers.PostUser)
	user.GET(byId, controllers.GetUserById)
	user.DELETE(byId, controllers.DeleteUser)
}
