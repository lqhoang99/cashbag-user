package routes

import (
	"github.com/labstack/echo/v4"

	"demo-user/controllers"
	"demo-user/validations"
)

// User ...
func User(e *echo.Echo) {
	routes := e.Group("/users")

	routes.POST("", controllers.UserCreate, validations.UserCreate)
	routes.PUT("/:id", controllers.UserUpdate, validations.UserValidateID, userCheckExistedByID, validations.UserUpdate)
	routes.GET("/:id",controllers.UserFindByID,validations.UserValidateID,userCheckExistedByID)

}
