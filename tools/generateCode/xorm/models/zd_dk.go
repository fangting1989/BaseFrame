package models
import (
		"time"
		)
type zd_dk struct {
  ZD_DK_ID  	int64     	`xorm:"not nullBIGINT(20)"`
  JZDSX_GUID	string    	`xorm:"VARCHAR(50)"`
  ZD_GUID   	string    	`xorm:"VARCHAR(50)"`
  DKBH      	string    	`xorm:"VARCHAR(50)"`
  DKMC      	string    	`xorm:"VARCHAR(100)"`
  JZDS      	float64   	`xorm:"DECIMAL(10,0)"`
  DKMJ      	float64   	`xorm:"DECIMAL(18,4)"`
  DKYT      	string    	`xorm:"VARCHAR(100)"`
  TFH       	string    	`xorm:"VARCHAR(50)"`
}