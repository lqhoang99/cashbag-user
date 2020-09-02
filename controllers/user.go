package controllers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"

	"demo-user/models"
	"demo-user/services"
	"demo-user/util"
)

// UserCreate ...
func UserCreate(c echo.Context) error {
	var (
		body = c.Get("body").(models.UserCreatePayload)
	)

	//Process data
	rawData, err := services.UserCreate(body)

	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	//Success
	return util.Response200(c, bson.M{
		"_id":       rawData.ID,
		"createdAt": rawData.CreatedAt,
	}, "")

}

// UserUpdate ...
func UserUpdate(c echo.Context) error {
	var (
		body   = c.Get("body").(models.UserUpdatePayload)
		user   = c.Get("user").(models.UserBSON)
		userID = user.ID
	)

	// Process data
	rawData, err := services.UserUpdate(userID, body)

	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	//success
	return util.Response200(c, bson.M{
		"updatedAt": rawData.UpdatedAt,
	}, "")
}

// UserFindByID ....
func UserFindByID(c echo.Context) error {
	var (
		user   = c.Get("user").(models.UserBSON)
		userID = user.ID
	)

	// Process data
	rawData,err :=services.UserFindByID(userID)
	
	// if err
	if err != nil {
		return util.Response400(c, nil, err.Error())
	}

	// success 
	return util.Response200(c,rawData,"")
}