package routes

import (
	"capstone/controller"
	"capstone/middleware"
	"capstone/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	e.POST("/register", controller.RegisterUserController)
	e.POST("/login", controller.LoginUserController)

	adm := e.Group("/admin", middleware.IsLoggedIn)
	adm.GET("/users", controller.GetUsersController)
	adm.PUT("/user/:id", controller.UpdateUserController)
	adm.POST("/plant", controller.CreatePlantController)
	adm.GET("/plants", controller.GetAllPlantController)
	adm.GET("/plant/:id", controller.GetPlantByIdController)
	adm.PUT("/plant/:id", controller.UpdatePlantController)
	adm.DELETE("/plant/:id", controller.DeletePlantController)
	adm.GET("/bookmark", controller.GetAllBookmarkPlantController)
	adm.DELETE("/bookmark/:id", controller.DeleteBookmarkController)

	pl := e.Group("/plant")
	pl.POST("/bookmark", controller.AddBookmarkPlantController, middleware.IsLoggedIn)
	pl.GET("", controller.GetAllPlantController, middleware.IsLoggedIn)
	pl.GET("/search", controller.SearchPlantController)
	pl.GET("/scan", controller.ScanPlantController, middleware.IsLoggedIn) //

	us := e.Group("/user", middleware.IsLoggedIn)
	us.GET("/bookmark", controller.GetBookmarkPlantListByUserIdController)

	upload := e.Group("/upload")
	upload.POST("/image", controller.UploadFileController)
}
