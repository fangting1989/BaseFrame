package models
import (
		"time"
		)
type Gdxm_info struct {
  GDXM_INFOID	int64     	`xorm:" not null  pk BIGINT(20)"`
  GD_GUID   	string    	`xorm:"VARCHAR(50)"`
  ZD_GUID   	string    	`xorm:"VARCHAR(50)"`
  YJV       	float64   	`xorm:"DECIMAL(16,2)"`
  STATE     	int       	`xorm:"TINYINT(4)"` 
  ZD_BH     	string    	`xorm:"VARCHAR(500)"`
  DKMC      	string    	`xorm:"VARCHAR(500)"`
  Province  	string    	`xorm:"VARCHAR(50)"`
  City      	string    	`xorm:"VARCHAR(50)"`
  County    	string    	`xorm:"VARCHAR(50)"`
  TD_ZL     	string    	`xorm:"VARCHAR(500)"`
  TD_YT     	string    	`xorm:"VARCHAR(50)"`
  MJ        	float64   	`xorm:"DECIMAL(20,4)"`
  JZMJ      	float64   	`xorm:"DECIMAL(20,4)"`
  CR_NX     	string    	`xorm:"VARCHAR(500)"`
  MAX_RJL   	float64   	`xorm:"DECIMAL(10,2)"`
  MAX_LHL   	float64   	`xorm:"DECIMAL(10,2)"`
  MAX_JZ_MD 	float64   	`xorm:"DECIMAL(10,2)"`
  MAX_JZ_XG 	float64   	`xorm:"DECIMAL(10,2)"`
  JC_SS     	string    	`xorm:"VARCHAR(500)"`
  TD_XZ     	string    	`xorm:"VARCHAR(50)"`
  QSJ       	float64   	`xorm:"DECIMAL(16,2)"`
  QSLMDJ    	float64   	`xorm:"DECIMAL(16,2)"`
  CJRQ      	time.Time 	`xorm:"DATE"`       
  CJ_JG     	float64   	`xorm:"DECIMAL(16,2)"`
  LMDJ      	float64   	`xorm:"DECIMAL(16,2)"`
  SDDR      	int       	`xorm:"INT(11)"`    
  JGSJ      	time.Time 	`xorm:"DATE"`       
  DGSJ      	time.Time 	`xorm:"DATE"`       
  XZQ_DM    	string    	`xorm:"VARCHAR(255)"`
}