package models
import (
"time"
)
type Qyydxsqk struct {
  QYYDXSQKID	int64     	`xorm:" not null  pk BIGINT(20)"`
  Company_ID	int64     	`xorm:"BIGINT(20)"` 
  PROVINCE  	string    	`xorm:"VARCHAR(255)"`
  CITY      	string    	`xorm:"VARCHAR(255)"`
  AREA      	string    	`xorm:"VARCHAR(255)"`
  Project   	string    	`xorm:"VARCHAR(255)"`
  CJSJ      	time.Time 	`xorm:"DATE"`       
  XSTS      	string    	`xorm:"VARCHAR(200)"`
  XSMJ      	string    	`xorm:"VARCHAR(200)"`
  XSJE      	string    	`xorm:"VARCHAR(200)"`
  XSJJ      	string    	`xorm:"VARCHAR(200)"`
  WYLX      	string    	`xorm:"VARCHAR(200)"`
  XSZT      	string    	`xorm:"VARCHAR(200)"`
  KPRQ      	time.Time 	`xorm:"DATE"`       
  JZMJ      	string    	`xorm:"VARCHAR(200)"`
  RJL       	string    	`xorm:"VARCHAR(200)"`
  LHL       	string    	`xorm:"VARCHAR(200)"`
  GHHS      	string    	`xorm:"VARCHAR(200)"`
  QYMC      	string    	`xorm:"VARCHAR(200)"`
}