package models
			import (
			"time"
			)
			type rolemodule struct {
			  RoleModuleID	int       	`xorm:"not null pk INT(11)"`
  RoleID    	int       	`xorm:"INT(11)"`    
  ModuleID  	int       	`xorm:"INT(11)"`    
  ModulePermission	string    	`xorm:"VARCHAR(200)"`
}