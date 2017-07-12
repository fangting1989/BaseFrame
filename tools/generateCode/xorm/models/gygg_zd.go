package models
import (
		"time"
		)
type gygg_zd struct {
  GYGG_ZD_ID	int64     	`xorm:"not nullBIGINT(20)"`
  LAND_ID   	string    	`xorm:"CHAR(10)"`   
  GYGG_ID   	int64     	`xorm:"BIGINT(20)"` 
  GYGG_ZD_GUID	string    	`xorm:"VARCHAR(50)"`
  GYGG_GUID 	string    	`xorm:"VARCHAR(50)"`
  ZD_BH     	string    	`xorm:"VARCHAR(50)"`
  MJ        	float64   	`xorm:"DECIMAL(18,3)"`
  TD_YT     	string    	`xorm:"VARCHAR(50)"`
  LP_ZT     	string    	`xorm:"VARCHAR(50)"`
  ZD_ZL     	string    	`xorm:"VARCHAR(200)"`
  MIN_RJL   	float64   	`xorm:"DECIMAL(18,3)"`
  MIN_RJL_TAG	string    	`xorm:"VARCHAR(10)"`
  MAX_RJL   	float64   	`xorm:"DECIMAL(18,3)"`
  MAX_RJL_TAG	string    	`xorm:"VARCHAR(10)"`
  MIN_LHL   	float64   	`xorm:"DECIMAL(18,3)"`
  MIN_LHL_TAG	string    	`xorm:"VARCHAR(10)"`
  MAX_LHL   	float64   	`xorm:"DECIMAL(18,3)"`
  MAX_LHL_TAG	string    	`xorm:"VARCHAR(10)"`
  MIN_JZ_MD 	float64   	`xorm:"DECIMAL(18,3)"`
  MIN_JZ_MD_TAG	string    	`xorm:"VARCHAR(10)"`
  MAX_JZ_MD 	float64   	`xorm:"DECIMAL(18,3)"`
  MAX_JZ_MD_TAG	string    	`xorm:"VARCHAR(10)"`
  TZ_QD     	float64   	`xorm:"DECIMAL(18,3)"`
  CR_NX     	float64   	`xorm:"DECIMAL(10,0)"`
  CR_BZJ    	float64   	`xorm:"DECIMAL(19,4)"`
  QSJ       	float64   	`xorm:"DECIMAL(19,4)"`
  JJ_FD     	float64   	`xorm:"DECIMAL(19,4)"`
  TD_XZ     	string    	`xorm:"VARCHAR(200)"`
  GP_SJ_S   	time.Time 	`xorm:"DATE"`       
  GP_SJ_E   	time.Time 	`xorm:"DATE"`       
  MIN_JZ_XG 	float64   	`xorm:"DECIMAL(18,3)"`
  MIN_JZ_XG_TAG	string    	`xorm:"VARCHAR(10)"`
  MAX_JZ_XG 	float64   	`xorm:"DECIMAL(18,3)"`
  MAX_JZ_XG_TAG	string    	`xorm:"VARCHAR(10)"`
  JC_SS     	string    	`xorm:"VARCHAR(100)"`
  QD_PZ     	string    	`xorm:"VARCHAR(200)"`
  XZ_TD_TJ  	string    	`xorm:"VARCHAR(200)"`
  RJL_GX    	string    	`xorm:"VARCHAR(50)"`
  LHL_GX    	string    	`xorm:"VARCHAR(50)"`
  JZ_MD_GX  	string    	`xorm:"VARCHAR(50)"`
  JZ_XG_GX  	string    	`xorm:"VARCHAR(50)"`
  QSJ_DW    	string    	`xorm:"VARCHAR(50)"`
  JJFD_DW   	string    	`xorm:"VARCHAR(50)"`
  BZ        	string    	`xorm:"VARCHAR(4000)"`
}