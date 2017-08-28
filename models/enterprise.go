package models
import (
"time"
)
type Enterprise struct {
  Enterprise_ID	int64     	`xorm:" not null  pk BIGINT(20)"`
  Logourl   	string    	`xorm:"VARCHAR(200)"`
  Legalperson	string    	`xorm:"VARCHAR(50)"`
  JobName   	string    	`xorm:"VARCHAR(50)"`
  Linktel   	string    	`xorm:"VARCHAR(50)"`
  Buscard   	string    	`xorm:"VARCHAR(50)"`
  Startdate 	time.Time 	`xorm:"DATE"`       
  Enddate   	time.Time 	`xorm:"DATE"`       
  Isvalid   	int       	`xorm:"TINYINT(1)"` 
  Validdate 	time.Time 	`xorm:"DATE"`       
  Remark    	string    	`xorm:"VARCHAR(500)"`
}