package models
			import (
			"time"
			)
			type smsmessage struct {
			  Smsmessageid	int64     	`xorm:"not null pk BIGINT(20)"`
  Content   	string    	`xorm:"VARCHAR(255)"`
  Phone     	string    	`xorm:"VARCHAR(20)"`
  Sendtime  	time.Time 	`xorm:"DATETIME"`   
  State     	int       	`xorm:"TINYINT(4)"` 
  Type      	int       	`xorm:"TINYINT(4)"` 
}