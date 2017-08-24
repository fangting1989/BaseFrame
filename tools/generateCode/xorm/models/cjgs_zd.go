package models
import (
		"time"
		)
type Cjgs_zd struct {
  CJGS_ZD_GUID	string    	`xorm:" not null  pk VARCHAR(100)"`
  CJGS_GUID 	string    	`xorm:"VARCHAR(100)"`
  SR_DW     	string    	`xorm:"VARCHAR(200)"`
  XM_MC     	string    	`xorm:"VARCHAR(500)"`
  TD_ZL     	string    	`xorm:"TEXT"`       
  TD_MJ     	float64   	`xorm:"DECIMAL(18,6)"`
  TD_YT     	string    	`xorm:"VARCHAR(50)"`
  TD_SY_TJ  	string    	`xorm:"TEXT"`       
  CR_NX     	string    	`xorm:"VARCHAR(50)"`
  CJ_JG     	float64   	`xorm:"DECIMAL(18,6)"`
  ZD_BH     	string    	`xorm:"VARCHAR(100)"`
  BZ        	string    	`xorm:"TEXT"`       
  HAVE_YCDK 	string    	`xorm:"VARCHAR(50)"`
  GYGG_ZD_GUID	string    	`xorm:"VARCHAR(100)"`
}