package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"demo-user/models"
	"demo-user/modules/database"
)

// UserCreate ...
func UserCreate(doc models.UserBSON) (models.UserBSON, error) {
	var (
		collection = database.UserCol()
		ctx        = context.Background()
	)

	// Insert one
	_, err := collection.InsertOne(ctx, doc)
	return doc, err

}

// UserFindByID ...
func UserFindByID(id primitive.ObjectID) (models.UserBSON, error) {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
		result  models.UserBSON
		filter  = bson.M{"_id": id}
	)

	// Find
	err := userCol.FindOne(ctx, filter).Decode(&result)

	return result, err
}

// UserUpdateByID ...
func UserUpdateByID(filter bson.M, updateData bson.M) error {
	var (
		userCol = database.UserCol()
		ctx     = context.Background()
	)

	// Update
	_, err := userCol.UpdateOne(ctx, filter, updateData)

	return err
}
