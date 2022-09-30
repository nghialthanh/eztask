package services

import (
	"context"
	"eztask/database"
	"eztask/models"
)

var db = database.Init()

// CreateUser is the api used to tget a single user
func GetUserByEmail(ctx context.Context, email string) (models.User, error) {

	var err error

	user := models.User{}

	err = db.Debug().WithContext(ctx).Model(models.User{}).Where("email = ?", email).Take(&user).Error

	return user, err
}
