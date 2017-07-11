package models
			import (
			"time"
			)
			type module struct {
			  ModuleID  	int       	`xorm:"not null pk INT(11)"`
  ModuleTypeID	int       	`xorm:"INT(11)"`    
  ModuleName	string    	`xorm:"VARCHAR(50)"`
  ModuleIndex	int       	`xorm:"INT(11)"`    
  Valid     	int       	`xorm:"INT(11)"`    
  ModuleIcon	string    	`xorm:"VARCHAR(50)"`
  ModuleUrl 	string    	`xorm:"VARCHAR(50)"`
  ModulePermission	string    	`xorm:"VARCHAR(50)"`
  ModuleDesc	string    	`xorm:"VARCHAR(200)"`
}