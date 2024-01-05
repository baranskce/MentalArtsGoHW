package main

import (
	"fmt"
	"mentalarts_go/database"
	"mentalarts_go/docs"
	"mentalarts_go/handlers"
	middleware "mentalarts_go/middlewares"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// gin-swagger middleware
// swagger embed files

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	r := gin.New()

	db, err := database.InitDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mw := middleware.NewMentalArtsMiddleware()
	r.Use(mw.Logger())

	docs.SwaggerInfo.BasePath = "/"

	h := handlers.NewHandler(db)

	v1 := r.Group("/v1")
	{
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // Swagger page router

		album := v1.Group("/album")
		{
			album.POST("", h.CreateAlbum)
			album.GET("", h.GetAlbums)
			album.GET(":id", h.GetAlbumByID)
			album.PUT(":id", h.UpdateAlbum)
			album.DELETE(":id", h.DeleteAlbumByID)
			album.DELETE("", h.DeleteAllAlbums)
		}
		musician := v1.Group("/musician")
		{
			musician.POST("", h.CreateMusician)
			musician.GET("", h.GetMusicians)
			musician.GET(":id", h.GetMusicianById)
			musician.PUT(":id", h.UpdateMusician)
			musician.DELETE(":id", h.DeleteMusicianById)
			musician.DELETE("", h.DeleteAllMusicians)
		}
	}
	err = r.Run(":8080") // For example, run on port 8080
	if err != nil {
		fmt.Println("Failed to start server:", err)
		os.Exit(1)
	}
}
