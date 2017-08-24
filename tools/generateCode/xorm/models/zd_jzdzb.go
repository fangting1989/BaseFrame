package models
import (
		"time"
		)
type Zd_jzdzb struct {
  ZD_JZDZB_ID	int64     	`xorm:" not null  pk BIGINT(20)"`
  JZD_GUID  	string    	`xorm:"VARCHAR(50)"`
  ZD_GUID   	string    	`xorm:"VARCHAR(50)"`
  JZDSX_GUID	string    	`xorm:"VARCHAR(50)"`
  DKBH      	string    	`xorm:"VARCHAR(50)"`
  JZDH      	string    	`xorm:"VARCHAR(50)"`
  QH        	string    	`xorm:"VARCHAR(50)"`
  FSBC      	float64   	`xorm:"DECIMAL(18,6)"`
  BZ        	string    	`xorm:"VARCHAR(200)"`
  JMZB      	string    	`xorm:"VARCHAR(255)"`
  XZB       	float64   	`xorm:"DECIMAL(18,6)"`
  YZB       	float64   	`xorm:"DECIMAL(18,6)"`
  JZDORDER  	float64   	`xorm:"DECIMAL(10,0)"`
}