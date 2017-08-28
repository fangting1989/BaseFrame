package models
import (
"time"
)
type Gzqy struct {
  GZQY_ID   	int64     	`xorm:" not null  pk BIGINT(20)"`
  Company_id	int64     	`xorm:"BIGINT(20)"` 
  Member_ID 	int64     	`xorm:"BIGINT(20)"` 
  GZSJ      	time.Time 	`xorm:"DATETIME"`   
  YXZT      	int       	`xorm:"INT(11)"`    
}