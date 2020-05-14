package model

import (
	"go-scheduler/pkg/util"
)

type User struct {
	Id        string    `json:"id" xorm:"not null pk comment('用户ID') CHAR(36)"`
	Name      string    `json:"name" xorm:"not null comment('用户名') CHAR(50)"`
	Email     string    `json:"email" xorm:"not null comment('邮箱') CHAR(255)" `
	Password  string    `json:"-" xorm:"not null comment('密码') CHAR(255)" `
	Manager   bool      `json:"manager" xorm:"not null default 0 comment('管理员') TINYINT(1)"`
	CreatedAt util.Time `json:"created_at" xorm:"not null created comment('创建日期') DATETIME"`
	UpdatedAt util.Time `json:"updated_at" xorm:"not null updated comment('更新日期') DATETIME"`
}
