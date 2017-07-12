package models
import (
		"time"
		)
type tdjbxx struct {
  TDJBXX_ID 	int64     	`xorm:"not nullBIGINT(20)"`
  ZDDXZQ_ID 	int64     	`xorm:"BIGINT(20)"` 
  ZD_DK_ID  	int64     	`xorm:"BIGINT(20)"` 
  ZD_JZDZB_ID	int64     	`xorm:"BIGINT(20)"` 
  ZD_GUID   	string    	`xorm:"VARCHAR(50)"`
  STATE     	int       	`xorm:"TINYINT(4)"` 
  ZD_BH     	string    	`xorm:"VARCHAR(50)"`
  DKMC      	string    	`xorm:"VARCHAR(100)"`
  Province  	string    	`xorm:"VARCHAR(50)"`
  City      	string    	`xorm:"VARCHAR(50)"`
  County    	string    	`xorm:"VARCHAR(50)"`
  TD_ZL     	string    	`xorm:"VARCHAR(250)"`
  TD_YT     	string    	`xorm:"VARCHAR(50)"`
  MJ        	float64   	`xorm:"DECIMAL(20,4)"`
  JZMJ      	float64   	`xorm:"DECIMAL(20,4)"`
  CR_NX     	float64   	`xorm:"DECIMAL(10,0)"`
  MAX_RJL   	float64   	`xorm:"DECIMAL(10,2)"`
  MAX_LHL   	float64   	`xorm:"DECIMAL(10,2)"`
  MAX_JZ_MD 	float64   	`xorm:"DECIMAL(10,2)"`
  MAX_JZ_XG 	float64   	`xorm:"DECIMAL(10,2)"`
  JC_SS     	string    	`xorm:"VARCHAR(100)"`
}