package repository

import (
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	TelegramID int64 `gorm:"index:,unique"`
}

type users struct{}

func (u *users) Create(usr *UserModel) error {
	return db.Create(usr).Error
}

func (u *users) GetAllTelegramID() []int64 {
	var allUsers []UserModel
	db.Find(&allUsers)

	var allTelegramID = make([]int64, len(allUsers))
	for i := range allTelegramID {
		allTelegramID[i] = allUsers[i].TelegramID
	}

	return allTelegramID
}
