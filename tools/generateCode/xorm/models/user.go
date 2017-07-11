package models
			import (
			"time"
			)
			type user struct {
			  UserID    	int       	`xorm:"not null pk INT(11)"`
  UserName  	string    	`xorm:"VARCHAR(50)"`
  UserMobile	string    	`xorm:"VARCHAR(20)"`
  UserStatus	float64   	`xorm:"DECIMAL(1,0)"`
  CreateDate	string    	`xorm:"VARCHAR(30)"`
  UserMail  	string    	`xorm:"VARCHAR(50)"`
  UserAddress	string    	`xorm:"VARCHAR(200)"`
  UserPwd   	string    	`xorm:"VARCHAR(50)"`
}