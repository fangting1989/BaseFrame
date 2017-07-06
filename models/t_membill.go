package models
import (
  "time"
)
type T_membill struct {
  MemBillID 	int64     	`gorm:"column:MemBillID;NOT NULL;PRIMARY KEY;"`
  MemBillAmount	float64   	`gorm:"column:MemBillAmount;"`
  BillCostTime	time.Time 	`gorm:"column:BillCostTime;"`
  BillPayTime	time.Time 	`gorm:"column:BillPayTime;"`
  BillPayCode	string    	`gorm:"column:BillPayCode;"`
  BillPayType	int       	`gorm:"column:BillPayType;"`
  BillPoint 	float64   	`gorm:"column:BillPoint;"`
  MemID     	int64     	`gorm:"column:MemID;"`
  MemBillGuid	string    	`gorm:"column:MemBillGuid;"`
  PayState  	int       	`gorm:"column:PayState;"`
}