package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	ID               string `gorm:"primaryKey"`
	Name             string
	ChatID           string `gorm:"uniqueIndex"`
	LinkCode         string
	LinkCodeExpireAt int64
}

func GetOrCreateUser(ctx context.Context, id, name string) (*User, error) {
	db := GetDB()
	var user User
	user, err := gorm.G[User](db).Where("id = ?", id).First(ctx)
	//err := db.First(&user, "id = ?", id).Error
	if err == nil {
		return &user, nil
	}
	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	user = User{ID: id, Name: name}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
