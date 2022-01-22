package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	TelegramID int64 `gorm:"index:,unique"`
	ReplyMsgID int
}

type users struct{}

func (u *users) Create(usr *UserModel) error {
	return db.Create(usr).Error
}

func (u *users) GetAllTelegramID() []int64 {
	var allUsers []UserModel
	if err := db.Find(&allUsers).Error; err != nil {
		fmt.Println(err)
	}

	var allTelegramID = make([]int64, len(allUsers))
	for i := range allTelegramID {
		allTelegramID[i] = allUsers[i].TelegramID
	}

	return allTelegramID
}

func (u *users) Count() (n int64) {
	if err := db.Model(&UserModel{}).Count(&n).Error; err != nil {
		fmt.Println(err)
	}
	return n
}

func (u *users) GetReplyMsgID(userTelegramID int64) (ID int) {
	var usr UserModel
	if err := db.Model(&UserModel{}).Where("telegram_id = ?", userTelegramID).First(&usr).Error; err != nil {
		fmt.Println(err)
	}
	return usr.ReplyMsgID
}

func (u *users) SetReplyMsgID(userTelegramID int64, messageID int) {
	if err := db.Model(&UserModel{}).Where("telegram_id = ?", userTelegramID).Update("reply_msg_id", messageID).Error; err != nil {
		fmt.Println(err)
	}
}
