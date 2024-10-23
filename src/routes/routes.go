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

	product := g.Group("/product")

	product.Use(middleware.JWTAuthMiddleware())
	{
		product.GET("/", controllers.GetAllProducts)
		product.POST("/", controllers.PostProduct)
		product.GET(byId, controllers.GetProductById)
		product.PATCH(byId, controllers.UpdateProductById)
		product.DELETE(byId, controllers.DeleteProduct)
	}

	category := g.Group("/category")

	category.Use(middleware.JWTAuthMiddleware())
	{
		category.GET("/", controllers.GetAllcategories)
		// category.POST("/", controllers.PostCategory)
		// category.GET(byId, controllers.GetCategoryById)
		// category.PATCH(byId, controllers.UpdateCategoryById)
		// category.DELETE(byId, controllers.DeleteCategory)
	}

	sale := g.Group("/sale")

	sale.Use(middleware.JWTAuthMiddleware())
	{
		sale.POST("/", controllers.PostSale)
		// sale.GET("/", controllers.GetAllSales)
		// sale.GET(byId, controllers.GetSaleById)
		// sale.DELETE(byId, controllers.DeleteSale)
	}
}
