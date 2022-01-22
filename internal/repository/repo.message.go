package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type MessageModel struct {
	gorm.Model
	MessageID   int `gorm:"index:,unique"`
	FromID      int64
	RecipientID int64
}

type messages struct{}

func (m *messages) Add(messageID int, fromID int64, recipID int64) error {
	return db.Create(&MessageModel{
		MessageID:   messageID,
		FromID:      fromID,
		RecipientID: recipID,
	}).Error
}

func (m *messages) Get(messageID int) *MessageModel {
	var msg MessageModel
	if err := db.Where("message_id = ?", messageID).First(&msg).Error; err != nil {
		fmt.Println(err)
	}
	return &msg
}

func (m *messages) Count() (n int64) {
	if err := db.Model(&MessageModel{}).Count(&n).Error; err != nil {
		fmt.Println(err)
	}
	return n
}
