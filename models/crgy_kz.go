package models

import (
	"time"
)

type Crgy_kz struct {
	CRGY_KZ_ID      int64     `xorm:"not nullBIGINT(20)"`
	CRKZ_GUID       string    `xorm:"VARCHAR(50)"`
	GD_GUID         string    `xorm:"VARCHAR(50)"`
	PM_JZ           string    `xorm:"VARCHAR(255)"`
	SX_SJX          string    `xorm:"VARCHAR(255)"`
	SX_XJX          string    `xorm:"VARCHAR(255)"`
	GC              float64   `xorm:"DECIMAL(16,2)"`
	TD_TJ_XX        string    `xorm:"VARCHAR(10)"`
	CD_PZ           string    `xorm:"VARCHAR(1000)"`
	JS_SS           string    `xorm:"VARCHAR(1000)"`
	XZ_TD_TJ        string    `xorm:"VARCHAR(2000)"`
	DW_JE           float64   `xorm:"DECIMAL(18,4)"`
	DJ              float64   `xorm:"DECIMAL(18,4)"`
	DJ_ZF_RQ        time.Time `xorm:"DATE"`
	TZQD_XX         string    `xorm:"VARCHAR(10)"`
	PT_XX           string    `xorm:"VARCHAR(255)"`
	PTYD_BL         float64   `xorm:"DECIMAL(8,2)"`
	PTYD_MJ         float64   `xorm:"DECIMAL(18,4)"`
	MAX_PT_JZ_MJ    float64   `xorm:"DECIMAL(18,4)"`
	MIN_ZZ_TS       float64   `xorm:"DECIMAL(10,0)"`
	MIN_90_ZZ_TS    float64   `xorm:"DECIMAL(10,0)"`
	MIN_90_ZZ_TS_BS float64   `xorm:"DECIMAL(16,2)"`
	ZZ_TX_YQ        string    `xorm:"VARCHAR(255)"`
	ZZ_PT_LX_XX     string    `xorm:"VARCHAR(10)"`
	BC_ZZ_PT_LX_FS  string    `xorm:"VARCHAR(1000)"`
	PT_XM           string    `xorm:"VARCHAR(1000)"`
	TD_YT_GG_CL_XX  string    `xorm:"VARCHAR(6)"`
	TD_ZR_TJ_XX     string    `xorm:"VARCHAR(6)"`
	JZW_SH_FS_XX    string    `xorm:"VARCHAR(6)"`
	ZLJ_BL          float64   `xorm:"DECIMAL(16,2)"`
	GYGG_ZD_GUID    string    `xorm:"VARCHAR(50)"`
	PT_FY           float64   `xorm:"DECIMAL(18,4)"`
	BC_TK           string    `xorm:"VARCHAR(8000)"`
	WGGZD_SM        string    `xorm:"VARCHAR(200)"`
	GYGG_ZDBH       string    `xorm:"VARCHAR(255)"`
	CR_NX           string    `xorm:"VARCHAR(255)"`
}
