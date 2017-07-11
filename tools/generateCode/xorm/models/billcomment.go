package models
			import (
			"time"
			)
			type billcomment struct {
			  BillCommentID	int64     	`xorm:"not null pk BIGINT(20)"`
  Content   	string    	`xorm:"TEXT"`       
  Point     	float64   	`xorm:"DECIMAL(6,2)"`
  ComPerson 	string    	`xorm:"VARCHAR(50)"`
  Mobile    	string    	`xorm:"VARCHAR(50)"`
  PersonCode	string    	`xorm:"VARCHAR(50)"`
  MemBillID 	int64     	`xorm:"BIGINT(20)"` 
  DealConent	string    	`xorm:"TEXT"`       
  State     	int       	`xorm:"INT(11)"`    
  Commentdate	time.Time 	`xorm:"DATETIME"`   
  Dealdate  	time.Time 	`xorm:"DATETIME"`   
  Pid       	int64     	`xorm:"BIGINT(20)"` 
}