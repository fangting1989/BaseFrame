package models
			import (
			"time"
			)
			type member struct {
			  MemberID  	int64     	`xorm:"not null pk BIGINT(20)"`
  MemName   	string    	`xorm:"VARCHAR(50)"`
  MemPwd    	string    	`xorm:"VARCHAR(50)"`
  MemMobile 	string    	`xorm:"VARCHAR(20)"`
  MemOld    	string    	`xorm:"VARCHAR(50)"`
  MemAge    	string    	`xorm:"VARCHAR(50)"`
  MemMail   	string    	`xorm:"VARCHAR(50)"`
  WeChatOpenID	string    	`xorm:"VARCHAR(50)"`
  CostAmount	float64   	`xorm:"DECIMAL(10,2)"`
  CostNum   	int       	`xorm:"INT(11)"`    
  WeChatImg 	string    	`xorm:"VARCHAR(200)"`
  WeChatName	string    	`xorm:"VARCHAR(80)"`
  MemSex    	int       	`xorm:"INT(11)"`    
}