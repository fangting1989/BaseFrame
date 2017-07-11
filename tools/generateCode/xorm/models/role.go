package models
			import (
			"time"
			)
			type role struct {
			  RoleID    	int       	`xorm:"not null pk INT(11)"`
  RoleName  	string    	`xorm:"VARCHAR(50)"`
  RoleValid 	int       	`xorm:"INT(11)"`    
  RoleCreatedate	time.Time 	`xorm:"DATE"`       
  RoleDesc  	string    	`xorm:"VARCHAR(200)"`
}