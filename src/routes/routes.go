package routes

import (
	"api-rest-go/src/controllers"
	"api-rest-go/src/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(g *gin.Engine) {

	// Ruta pública para login
	g.POST("/login", controllers.Login)

	album := g.Group("/album")
	const byId string = "/:id"

	album.Use(middleware.JWTAuthMiddleware())
	{
		album.GET("/", controllers.GetAllAlbum)
		album.POST("/", controllers.PostAlbum)
		album.GET(byId, controllers.GetAlbumById)
		album.PATCH(byId, controllers.UpdateAlbumById)
		album.DELETE(byId, controllers.DeleteAlbum)
	}

	rol := g.Group("/rol")

	rol.Use(middleware.JWTAuthMiddleware())
	{
		rol.GET("/", controllers.GetRols)
		rol.POST("/", controllers.PostRol)
		rol.GET(byId, controllers.GetRolById)
		rol.DELETE(byId, controllers.DeleteRol)
	}

	user := g.Group("/user")

	user.Use(middleware.JWTAuthMiddleware())
	{
		user.GET("/", controllers.GetUsers)
		user.POST("/", controllers.PostUser)
		user.GET(byId, controllers.GetUserById)
		user.DELETE(byId, controllers.DeleteUser)
	}

}
