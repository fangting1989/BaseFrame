package models
import (
"time"
)
type Sqqy struct {
  SQQY_ID   	int64     	`xorm:" not null  pk BIGINT(20)"`
  Company_ID	string    	`xorm:"VARCHAR(100)"`
  Member_ID 	string    	`xorm:"VARCHAR(100)"`
  AuthorizedType	int       	`xorm:"TINYINT(1)"` 
  StartDate 	time.Time 	`xorm:"DATE"`       
  EndDate   	time.Time 	`xorm:"DATE"`       
  Remark    	string    	`xorm:"VARCHAR(500)"`
}