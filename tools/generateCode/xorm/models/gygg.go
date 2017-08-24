package models
import (
		"time"
		)
type Gygg struct {
  GYGG_GUID 	string    	`xorm:" not null  pk VARCHAR(100)"`
  GYGG_BT   	string    	`xorm:"VARCHAR(255)"`
  JZ_RQ     	time.Time 	`xorm:"DATETIME(6)"`
  XZQ_DM    	string    	`xorm:" not null VARCHAR(9)"`
  SBR       	string    	`xorm:"VARCHAR(50)"`
  XM_CJ     	float64   	`xorm:"DECIMAL(10,0)"`
  XM_ZT     	string    	`xorm:"VARCHAR(6)"` 
  CREATE_DATE	time.Time 	`xorm:"DATETIME(6)"`
  CREATE_USER	string    	`xorm:"VARCHAR(50)"`
  DELETE_DATE	time.Time 	`xorm:"DATETIME(6)"`
  SH_ZT     	string    	`xorm:"VARCHAR(2)"` 
  FB_SJ     	time.Time 	`xorm:"DATETIME(6)"`
  JZ_RQ_END 	time.Time 	`xorm:"DATETIME(6)"`
  GG_LX     	string    	`xorm:"VARCHAR(50)"`
  OLD_GYGG_GUID	string    	`xorm:"VARCHAR(100)"`
  QT_GG_LX  	string    	`xorm:"VARCHAR(50)"`
  MODIFY_DATE	time.Time 	`xorm:"DATETIME(6)"`
  SF_TG     	float64   	`xorm:"DECIMAL(10,0)"`
  WL_BZ     	string    	`xorm:"VARCHAR(2)"` 
  SH_SJ     	time.Time 	`xorm:"DATETIME(6)"`
  OLD_GYGG_BT	string    	`xorm:"TEXT"`       
  CH_SJ     	time.Time 	`xorm:"DATETIME(6)"`
  SHR       	string    	`xorm:"VARCHAR(50)"`
  TB_ZT     	float64   	`xorm:"DECIMAL(10,0)"`
  CR_DW     	string    	`xorm:"VARCHAR(100)"`
  GG_BH     	string    	`xorm:"VARCHAR(50)"`
  PZ_JG     	string    	`xorm:"VARCHAR(100)"`
  LQ_WJ_SJ_S	time.Time 	`xorm:"DATETIME(6)"`
  LQ_WJ_SJ_E	time.Time 	`xorm:"DATETIME(6)"`
  LQ_WJ_DD  	string    	`xorm:"TEXT"`       
  TJ_SQ_SJ_S	time.Time 	`xorm:"DATETIME(6)"`
  TJ_SQ_SJ_E	time.Time 	`xorm:"DATETIME(6)"`
  TJ_SQ_DD  	string    	`xorm:"TEXT"`       
  JN_BZJ_JZ_SJ	time.Time 	`xorm:"DATETIME(6)"`
  QR_TB_ZG_SJ	time.Time 	`xorm:"DATETIME(6)"`
  ZB_SJ_S   	time.Time 	`xorm:"DATETIME(6)"`
  ZB_SJ_E   	time.Time 	`xorm:"DATETIME(6)"`
  CR_SJ     	time.Time 	`xorm:"DATETIME(6)"`
  CR_DD     	string    	`xorm:"TEXT"`       
  KB_SJ     	time.Time 	`xorm:"DATETIME(6)"`
  QT_SX     	string    	`xorm:"TEXT"`       
  LXR       	string    	`xorm:"VARCHAR(50)"`
  LX_DH     	string    	`xorm:"VARCHAR(100)"`
  LX_DZ     	string    	`xorm:"TEXT"`       
  KH_DW     	string    	`xorm:"VARCHAR(100)"`
  KHH       	string    	`xorm:"TEXT"`       
  YHZH      	string    	`xorm:"TEXT"`       
  SB_SJ     	time.Time 	`xorm:"DATETIME(6)"`
  ZD_SL     	float64   	`xorm:"DECIMAL(10,0)"`
  KB_DD     	string    	`xorm:"TEXT"`       
  ISNEW     	string    	`xorm:"VARCHAR(5)"` 
  SQ_FS     	string    	`xorm:"VARCHAR(4)"` 
  SQR_QT_TJ 	string    	`xorm:"TEXT"`       
  BSYJ      	float64   	`xorm:"DECIMAL(10,0)"`
  BG_NW     	string    	`xorm:"TEXT"`       
  CX_YY     	string    	`xorm:"TEXT"`       
  XGYY      	string    	`xorm:"TEXT"`       
  DK_JJ     	string    	`xorm:"TEXT"`       
  ZB_YZ     	string    	`xorm:"VARCHAR(50)"`
  GYGG_NR   	          	`xorm:"LONGTEXT"`   
}