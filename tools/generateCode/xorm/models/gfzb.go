package models
import (
		"time"
		)
type Gfzb struct {
  GFZBID    	int64     	`xorm:" not null  pk BIGINT(20)"`
  Company_ID	int64     	`xorm:"BIGINT(20)"` 
  GDMC      	string    	`xorm:"VARCHAR(100)"`
  CGSL      	string    	`xorm:"VARCHAR(50)"`
  CGZB      	string    	`xorm:"VARCHAR(50)"`
  BDQK      	string    	`xorm:"VARCHAR(50)"`
  SJCZE     	float64   	`xorm:"DECIMAL(16,2)"`
  SJCZSJ    	string    	`xorm:"VARCHAR(50)"`
  SHCZFS    	string    	`xorm:"VARCHAR(50)"`
  REMARK    	string    	`xorm:"VARCHAR(500)"`
}