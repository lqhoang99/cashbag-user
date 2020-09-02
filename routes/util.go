package routes

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-user/dao"
	"demo-user/util"
)

func userCheckExistedByID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			userID = c.Get("userID").(primitive.ObjectID)
		)

		user, _ := dao.UserFindByID(userID)

		// check existed
		if user.ID.IsZero() {
			return util.Response404(c, nil, "Khong tim thay user")
		}

		c.Set("user", user)
		return next(c)
	}
}
