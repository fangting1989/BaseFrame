package models
			import (
			"time"
			)
			type membill struct {
			  MemBillID 	int64     	`xorm:"not null pk BIGINT(20)"`
  MemBillAmount	float64   	`xorm:"DECIMAL(10,2)"`
  BillCostTime	time.Time 	`xorm:"DATE"`       
  BillPayTime	time.Time 	`xorm:"DATE"`       
  BillPayCode	string    	`xorm:"VARCHAR(50)"`
  BillPayType	int       	`xorm:"INT(11)"`    
  BillPoint 	float64   	`xorm:"DECIMAL(10,2)"`
  MemID     	int64     	`xorm:"BIGINT(20)"` 
  MemBillGuid	string    	`xorm:"VARCHAR(40)"`
  PayState  	int       	`xorm:"INT(11)"`    
}