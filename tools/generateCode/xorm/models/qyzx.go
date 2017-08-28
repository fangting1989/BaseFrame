package models
import (
"time"
)
type Qyzx struct {
  QYZXID    	int64     	`xorm:" not null  pk BIGINT(20)"`
  Company_ID	int64     	`xorm:"BIGINT(20)"` 
  ZXLB      	int       	`xorm:"TINYINT(4)"` 
  ZXBT      	string    	`xorm:"VARCHAR(250)"`
  ZXZBT     	string    	`xorm:"VARCHAR(250)"`
  ZXFBSJ    	time.Time 	`xorm:"DATE"`       
  ZXNR      	string    	`xorm:"MEDIUMTEXT"` 
  ZXFBR     	string    	`xorm:"VARCHAR(100)"`
  ZXLY      	string    	`xorm:"VARCHAR(100)"`
  ZXZT      	int       	`xorm:"TINYINT(4)"` 
  ZXZY      	string    	`xorm:"VARCHAR(200)"`
  ZXDZ      	string    	`xorm:"VARCHAR(200)"`
  REMARK    	string    	`xorm:"VARCHAR(500)"`
}