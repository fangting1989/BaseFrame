package models
import (
		"time"
		)
type Rzdt struct {
  RZDTID    	int64     	`xorm:" not null  pk BIGINT(20)"`
  Company_ID	int64     	`xorm:"BIGINT(20)"` 
  QYMC      	string    	`xorm:"VARCHAR(200)"`
  SJ        	time.Time 	`xorm:"DATE"`       
  RZLX      	string    	`xorm:"VARCHAR(200)"`
  GZ        	float64   	`xorm:"DECIMAL(16,2)"`
  JE        	float64   	`xorm:"DECIMAL(16,2)"`
  BL        	string    	`xorm:"VARCHAR(200)"`
  TZF       	string    	`xorm:"VARCHAR(200)"`
  XWLY      	string    	`xorm:"VARCHAR(2000)"`
}