package models
import (
		"time"
		)
type Crgy_kz struct {
  CRKZ_GUID 	string    	`xorm:" not null  pk VARCHAR(100)"`
  GD_GUID   	string    	`xorm:"VARCHAR(100)"`
  PM_JZ     	string    	`xorm:"VARCHAR(255)"`
  SX_SJX    	string    	`xorm:"VARCHAR(255)"`
  SX_XJX    	string    	`xorm:"VARCHAR(255)"`
  GC        	float64   	`xorm:"DECIMAL(16,2)"`
  TD_TJ_XX  	string    	`xorm:"VARCHAR(10)"`
  CD_PZ     	string    	`xorm:"TEXT"`       
  JS_SS     	string    	`xorm:"TEXT"`       
  XZ_TD_TJ  	string    	`xorm:"TEXT"`       
  DW_JE     	float64   	`xorm:"DECIMAL(18,4)"`
  DJ        	float64   	`xorm:"DECIMAL(18,4)"`
  DJ_ZF_RQ  	time.Time 	`xorm:"DATETIME(6)"`
  ZF_XX     	string    	`xorm:"VARCHAR(10)"`
  ZF_XD_TS  	float64   	`xorm:"DECIMAL(10,0)"`
  TZQD_XX   	string    	`xorm:"VARCHAR(10)"`
  PT_XX     	string    	`xorm:"VARCHAR(255)"`
  PTYD_BL   	float64   	`xorm:"DECIMAL(8,2)"`
  PTYD_MJ   	float64   	`xorm:"DECIMAL(18,4)"`
  MAX_PT_JZ_MJ	float64   	`xorm:"DECIMAL(18,4)"`
  MIN_ZZ_TS 	float64   	`xorm:"DECIMAL(10,0)"`
  MIN_90_ZZ_TS	float64   	`xorm:"DECIMAL(10,0)"`
  MIN_90_ZZ_TS_BS	float64   	`xorm:"DECIMAL(16,2)"`
  ZZ_TX_YQ  	string    	`xorm:"TEXT"`       
  ZZ_PT_LX_XX	string    	`xorm:"VARCHAR(10)"`
  BC_ZZ_PT_LX_FS	string    	`xorm:"TEXT"`       
  PT_XM     	string    	`xorm:"TEXT"`       
  TD_YT_GG_CL_XX	string    	`xorm:"VARCHAR(6)"` 
  TD_ZR_TJ_XX	string    	`xorm:"VARCHAR(6)"` 
  JZW_SH_FS_XX	string    	`xorm:"VARCHAR(6)"` 
  CRR       	string    	`xorm:"TEXT"`       
  CRR_DZ    	string    	`xorm:"TEXT"`       
  CRR_DH    	string    	`xorm:"TEXT"`       
  CRR_CZ    	string    	`xorm:"TEXT"`       
  CRR_YB    	string    	`xorm:"VARCHAR(10)"`
  CRR_KFH   	string    	`xorm:"TEXT"`       
  CRR_ZH    	string    	`xorm:"TEXT"`       
  SRR       	string    	`xorm:"TEXT"`       
  SRR_DZ    	string    	`xorm:"TEXT"`       
  SRR_DH    	string    	`xorm:"TEXT"`       
  SRR_CZ    	string    	`xorm:"TEXT"`       
  SRR_YB    	string    	`xorm:"VARCHAR(50)"`
  SRR_KFH   	string    	`xorm:"TEXT"`       
  SRR_ZH    	string    	`xorm:"TEXT"`       
  ZLJ_BL    	float64   	`xorm:"DECIMAL(16,2)"`
  KG_YQ_WYJ_BL	float64   	`xorm:"DECIMAL(16,2)"`
  JG_YQ_WYJ_BL	float64   	`xorm:"DECIMAL(16,2)"`
  LY_WYJ_BL 	float64   	`xorm:"DECIMAL(16,2)"`
  JD_YQ_WYJ_BL	float64   	`xorm:"DECIMAL(16,2)"`
  ZY_JJ_XX  	string    	`xorm:"VARCHAR(10)"`
  ZY_ZC_DW  	string    	`xorm:"VARCHAR(100)"`
  FS        	float64   	`xorm:"DECIMAL(10,0)"`
  MFS       	float64   	`xorm:"DECIMAL(10,0)"`
  GYGG_ZD_GUID	string    	`xorm:"VARCHAR(100)"`
  ZF_JSY    	float64   	`xorm:"DECIMAL(18,4)"`
  PT_FY     	float64   	`xorm:"DECIMAL(18,4)"`
  BC_TK     	          	`xorm:"LONGTEXT"`   
  WGGZD_SM  	          	`xorm:"LONGTEXT"`   
  GYGG_ZDBH 	string    	`xorm:"VARCHAR(255)"`
  SRR_FS    	float64   	`xorm:"DECIMAL(10,0)"`
  CRR_FS    	float64   	`xorm:"DECIMAL(10,0)"`
  CR_NX     	string    	`xorm:"VARCHAR(255)"`
  CJGS_ZD_BH	string    	`xorm:"VARCHAR(100)"`
  CJGS_ZD_GUID	string    	`xorm:"VARCHAR(100)"`
  WGSZD_SM  	string    	`xorm:"TEXT"`       
}