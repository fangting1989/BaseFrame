package models
import (
"time"
)
type Cgqk struct {
  CGQKID    	int64     	`xorm:" not null  pk BIGINT(200)"`
  Company_ID	int64     	`xorm:"BIGINT(20)"` 
  CSMC      	string    	`xorm:"VARCHAR(200)"`
  GSMC      	string    	`xorm:"VARCHAR(50)"`
  CGSJ      	time.Time 	`xorm:"DATE"`       
  CZBL      	string    	`xorm:"VARCHAR(50)"`
  RJCZ      	string    	`xorm:"VARCHAR(50)"`
  BTZFDDBR  	string    	`xorm:"VARCHAR(50)"`
  ZCZB      	float64   	`xorm:"DECIMAL(16,2)"`
  ZCSJ      	time.Time 	`xorm:"DATE"`       
  ZT        	string    	`xorm:"VARCHAR(50)"`
  REMARK    	string    	`xorm:"VARCHAR(200)"`
}