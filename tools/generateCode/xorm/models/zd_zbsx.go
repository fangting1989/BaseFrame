package models
import (
		"time"
		)
type zd_zbsx struct {
  ZD_ZBSX_ID	int64     	`xorm:"not nullBIGINT(20)"`
  ZB_GUID   	string    	`xorm:"VARCHAR(50)"`
  ZD_GUID   	string    	`xorm:"VARCHAR(50)"`
  ZBX       	string    	`xorm:"VARCHAR(50)"`
  TYLX      	string    	`xorm:"VARCHAR(50)"`
  JDFD      	string    	`xorm:"VARCHAR(50)"`
  DH        	string    	`xorm:"VARCHAR(50)"`
  JLDW      	string    	`xorm:"VARCHAR(20)"`
  JD        	float64   	`xorm:"DECIMAL(18,4)"`
}