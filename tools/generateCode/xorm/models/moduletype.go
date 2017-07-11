package models
			import (
			"time"
			)
			type moduletype struct {
			  ModuleTypeID	int       	`xorm:"not null pk INT(11)"`
  ModuleTypeName	string    	`xorm:"VARCHAR(50)"`
  ModuleTypeIndex	int       	`xorm:"INT(11)"`    
  ModuleTypePID	int       	`xorm:"INT(11)"`    
  Valid     	float64   	`xorm:"DECIMAL(1,0)"`
  ModuleTypeIcon	string    	`xorm:"VARCHAR(50)"`
  ModuleTypeDesc	string    	`xorm:"VARCHAR(200)"`
}