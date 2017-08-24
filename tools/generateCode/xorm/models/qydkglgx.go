package models
import (
		"time"
		)
type Qydkglgx struct {
  QYDKGLGXID	int64     	`xorm:" not null  pk BIGINT(20)"`
  Company_ID	int64     	`xorm:"BIGINT(20)"` 
  GD_GUID   	string    	`xorm:"VARCHAR(50)"`
  Getlanddate	time.Time 	`xorm:"DATETIME"`   
  Isvalid   	int       	`xorm:"TINYINT(4)"` 
  Remark    	string    	`xorm:"VARCHAR(2000)"`
}