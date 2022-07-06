package models

import "altman/global"

type Admin struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt int    `json:"created_at"`
}

func (Admin) TableName() string {
	return global.AmConfig.MySQL.Prefix + "admin"
}
