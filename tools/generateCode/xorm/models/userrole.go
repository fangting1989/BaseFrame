package models
			import (
			"time"
			)
			type userrole struct {
			  RoleID    	int       	`xorm:"INT(11)"`    
  UserID    	int       	`xorm:"INT(11)"`    
  UserRoleID	int       	`xorm:"not null pk INT(11)"`
}