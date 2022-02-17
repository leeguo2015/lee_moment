package model

import (
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

type (
	User struct {
		gorm.Model
		Uuid     string     `json:"uuid" gorm:"column:uuid;type:varchar(100);"`
		UserName string     `json:"user_name" gorm:"type:varchar(32);"`
		RealName string     `json:"real_name" gorm:"type:varchar(32);"`
		Password string     `json:"password" gorm:"type:varchar(128);"`
		Portrait string     `json:"portrait" gorm:"type:varchar(128);"`
		Gender   int        `json:"gender"`
		IDCard   string     `json:"id_card" gorm:"type:varchar(128);"`
		Addr     string     `json:"addr" gorm:"type:varchar(256);"`
		Phone    string     `json:"Phone" gorm:"type:varchar(128);"`
		Email    string     `json:"email" gorm:"type:varchar(128);"`
		IPAddr   string     `json:"ip_addr" gorm:"type:varchar(128);"`
		Status   int        `json:"status"` // 0 正常，1。冻结，2.注销
		Birth    time.Time  `json:"birth"`
		LastTime time.Time  `json:"last_time"`
		Roles    []AuthRole `gorm:"many2many:user_roles;"`
	}
)

// permission
type (
	AuthRole struct {
		gorm.Model
		Name        string           `json:"role_name" gorm:"type:varchar(64);"`
		Expound     string           `json:"expound" gorm:"type:varchar(128);"`
		Permissions []AuthPermission `gorm:"many2many:role_permissions;"`
	}

	AuthPermission struct {
		gorm.Model
		Name    string `json:"name" gorm:"type:varchar(64)"`
		Api     string `json:"lee_api" gorm:"type:varchar(128)"`
		Method  string `json:"method" gorm:"type:varchar(128)"`
		Expound string `json:"Expound" gorm:"type:varchar(256)"`
	}
)

// relationship
type (
	RelateBlue struct {
		gorm.Model
		UserId int
		PeerId int
	}
	RelateBlack struct {
		gorm.Model
		UserId int
		PeerId int
	}
	RelatePick struct {
		gorm.Model
		UserId int
		PeerId int
	}
)
