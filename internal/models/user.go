package models

import "fmt"

type User struct {
	Id         int64
	Name       string
	TelegramId int64
	IsDeleted  bool
}

func (u *User) String() string {
	return fmt.Sprintf("id: %d, name: %s, telegram_id: %d, IsDeleted: %t", u.Id, u.Name, u.TelegramId, u.IsDeleted)
}
