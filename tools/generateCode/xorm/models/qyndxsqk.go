package models
import (
		"time"
		)
type Qyndxsqk struct {
  QYNDXSQKID	int64     	`xorm:" not null  pk BIGINT(20)"`
  Company_ID	int64     	`xorm:"BIGINT(20)"` 
  BBLX      	int       	`xorm:"TINYINT(4)"` 
  BBLXMC    	string    	`xorm:"VARCHAR(50)"`
  XSND      	int       	`xorm:"INT(4)"`     
  XSMJ      	string    	`xorm:"VARCHAR(200)"`
  XSTS      	string    	`xorm:"VARCHAR(200)"`
  XSJE      	string    	`xorm:"VARCHAR(200)"`
  REMARK    	string    	`xorm:"VARCHAR(500)"`
  QYMC      	string    	`xorm:"VARCHAR(200)"`
}