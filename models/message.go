package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FormId   uint   // 发送者
	TargetId uint   // 接收者
	Type     int    // 消息类型 1私聊 2群聊 3广播
	Media    int    // 消息类型 1文字 2表情包 3图片 4音频
	Content  string // 消息内容
	Pic      string
	Url      string
	Desc     string
	Amount   int
}

func (table *Message) TableName() string {
	return "message"
}
