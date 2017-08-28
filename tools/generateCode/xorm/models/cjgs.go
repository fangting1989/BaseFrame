package models
import (
"time"
)
type Cjgs struct {
  CJGS_GUID 	string    	`xorm:" not null  pk VARCHAR(100)"`
  GSBT      	string    	`xorm:"VARCHAR(200)"`
  JZRQ      	time.Time 	`xorm:"DATETIME(6)"`
  GSNR      	string    	`xorm:"LONGTEXT"`   
  XZQ_DM    	string    	`xorm:" not null VARCHAR(9)"`
  SBR       	string    	`xorm:"VARCHAR(50)"`
  XM_CJ     	float64   	`xorm:"DECIMAL(10,0)"`
  XM_ZT     	string    	`xorm:"VARCHAR(50)"`
  CREATE_DATE	time.Time 	`xorm:"DATETIME(6)"`
  CREATE_USER	string    	`xorm:"VARCHAR(50)"`
  DELETE_DATE	time.Time 	`xorm:"DATETIME(6)"`
  SH_ZT     	string    	`xorm:"VARCHAR(50)"`
  FB_SJ     	time.Time 	`xorm:"DATETIME(6)"`
  GY_FS     	string    	`xorm:"VARCHAR(5)"` 
  MODIFY_DATE	time.Time 	`xorm:"DATETIME(6)"`
  SF_TG     	float64   	`xorm:"DECIMAL(10,0)"`
  WL_BZ     	string    	`xorm:"VARCHAR(2)"` 
  SH_SJ     	time.Time 	`xorm:"DATETIME(6)"`
  OLD_CJGS_GUID	string    	`xorm:"VARCHAR(100)"`
  CH_SJ     	time.Time 	`xorm:"DATETIME(6)"`
  SHR       	string    	`xorm:"VARCHAR(50)"`
  XGYY      	string    	`xorm:"TEXT"`       
  TB_ZT     	float64   	`xorm:"DECIMAL(10,0)"`
  GS_DW     	string    	`xorm:"VARCHAR(100)"`
  GS_SJ_S   	time.Time 	`xorm:"DATETIME(6)"`
  GS_SJ_E   	time.Time 	`xorm:"DATETIME(6)"`
  ZD_SL     	float64   	`xorm:"DECIMAL(10,0)"`
  ZPG_SJ_S  	time.Time 	`xorm:"DATETIME(6)"`
  ZPG_SJ_E  	time.Time 	`xorm:"DATETIME(6)"`
  ZPG_LX    	float64   	`xorm:"DECIMAL(10,0)"`
  GS_LX     	float64   	`xorm:"DECIMAL(10,0)"`
  LX_DW     	string    	`xorm:"VARCHAR(100)"`
  DW_DZ     	string    	`xorm:"VARCHAR(100)"`
  YZ_BM     	string    	`xorm:"VARCHAR(10)"`
  LX_DH     	string    	`xorm:"VARCHAR(50)"`
  LXR       	string    	`xorm:"VARCHAR(50)"`
  EMAIL     	string    	`xorm:"VARCHAR(100)"`
  SB_SJ     	time.Time 	`xorm:"DATETIME(6)"`
  CX_YY     	string    	`xorm:"TEXT"`       
  BG_NR     	string    	`xorm:"TEXT"`       
  ISNEW     	string    	`xorm:"VARCHAR(5)"` 
  GS_BH     	string    	`xorm:"VARCHAR(100)"`
}