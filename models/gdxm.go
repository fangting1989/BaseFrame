package models

import (
	"time"
)

type Gdxm struct {
	GDXM_ID       int64     `xorm:"not null BIGINT(20)"`
	CJGS_ZD_ID    int64     `xorm:"BIGINT(20)"`
	GD_GUID       string    `xorm:"VARCHAR(50)"`
	ZD_GUID       string    `xorm:"VARCHAR(50)"`
	XZQ_DM        string    `xorm:"VARCHAR(10)"`
	TD_ZL         string    `xorm:"VARCHAR(200)"`
	ZD_BH         string    `xorm:"VARCHAR(100)"`
	XM_MC         string    `xorm:"VARCHAR(500)"`
	XM_XH         float64   `xorm:"DECIMAL(5,0)"`
	TZZT_XZ       string    `xorm:"VARCHAR(10)"`
	HY_FL         string    `xorm:"VARCHAR(10)"`
	TD_JB         string    `xorm:"VARCHAR(30)"`
	GD_ZMJ        float64   `xorm:"DECIMAL(18,6)"`
	GY_FS         string    `xorm:"VARCHAR(10)"`
	GY_MJ         float64   `xorm:"DECIMAL(18,6)"`
	XZ_MJ         float64   `xorm:"DECIMAL(18,6)"`
	PC_MJ         float64   `xorm:"DECIMAL(18,6)"`
	GY_WLD_MJ     float64   `xorm:"DECIMAL(18,6)"`
	CL_MJ         float64   `xorm:"DECIMAL(18,6)"`
	CL_XZCR_MJ    float64   `xorm:"DECIMAL(18,6)"`
	CL_GBXZCR_MJ  float64   `xorm:"DECIMAL(18,6)"`
	CL_QT_MJ      float64   `xorm:"DECIMAL(18,6)"`
	CL_XQ         float64   `xorm:"DECIMAL(18,6)"`
	TD_YT         string    `xorm:"VARCHAR(6)"`
	CR_NX         string    `xorm:"VARCHAR(255)"`
	YT_SM         string    `xorm:"VARCHAR(255)"`
	QD_RQ         time.Time `xorm:"DATE"`
	JD_SJ         time.Time `xorm:"DATE"`
	DG_SJ         time.Time `xorm:"DATE"`
	JG_SJ         time.Time `xorm:"DATE"`
	BG_JD_SJ      time.Time `xorm:"DATE"`
	BG_DG_SJ      time.Time `xorm:"DATE"`
	BG_JG_SJ      time.Time `xorm:"DATE"`
	SJ_JD_SJ      time.Time `xorm:"DATE"`
	SJ_DG_SJ      time.Time `xorm:"DATE"`
	SJ_JG_SJ      time.Time `xorm:"DATE"`
	PZ_WH         string    `xorm:"VARCHAR(100)"`
	PZ_RQ         time.Time `xorm:"DATE"`
	PZ_JG         string    `xorm:"VARCHAR(100)"`
	KFJS_QX       float64   `xorm:"DECIMAL(8,2)"`
	JE            float64   `xorm:"DECIMAL(18,4)"`
	GDZC_TZE      float64   `xorm:"DECIMAL(18,4)"`
	TZ_QD         float64   `xorm:"DECIMAL(18,4)"`
	ZT_JZ_XZ      string    `xorm:"VARCHAR(255)"`
	FS_JZ_XZ      string    `xorm:"VARCHAR(255)"`
	JZ_MJ         float64   `xorm:"DECIMAL(18,4)"`
	MIN_JZ_MD     float64   `xorm:"DECIMAL(8,2)"`
	MIN_JZ_MD_TAG string    `xorm:"VARCHAR(10)"`
	MAX_JZ_MD     float64   `xorm:"DECIMAL(8,2)"`
	MAX_JZ_MD_TAG string    `xorm:"VARCHAR(10)"`
	MIN_RJL       float64   `xorm:"DECIMAL(8,2)"`
	MIN_RJL_TAG   string    `xorm:"VARCHAR(10)"`
	MAX_RJL       float64   `xorm:"DECIMAL(8,2)"`
	MAX_RJL_TAG   string    `xorm:"VARCHAR(10)"`
	MIN_LHL       float64   `xorm:"DECIMAL(8,2)"`
	MIN_LHL_TAG   string    `xorm:"VARCHAR(10)"`
	MAX_LHL       float64   `xorm:"DECIMAL(8,2)"`
	MAX_LHL_TAG   string    `xorm:"VARCHAR(10)"`
	MIN_JZ_GD     float64   `xorm:"DECIMAL(8,2)"`
	MIN_JZ_GD_TAG string    `xorm:"VARCHAR(10)"`
	MAX_JZ_GD     float64   `xorm:"DECIMAL(8,2)"`
	MAX_JZ_GD_TAG string    `xorm:"VARCHAR(10)"`
	QT_TD_YQ      string    `xorm:"VARCHAR(2000)"`
	KFTZ_ZE       float64   `xorm:"DECIMAL(18,3)"`
	ZD_SZ_QY      string    `xorm:"VARCHAR(50)"`
	HT_LX         string    `xorm:"VARCHAR(10)"`
	PARAID        float64   `xorm:"DECIMAL(2,0)"`
	KFLY_BZ       string    `xorm:"VARCHAR(5)"`
	WKFLY_YY      string    `xorm:"VARCHAR(3000)"`
	SF_SH         float64   `xorm:"DECIMAL(10,0)"`
	XZ_ZT         float64   `xorm:"DECIMAL(10,0)"`
	QY_SS_LX      string    `xorm:"VARCHAR(50)"`
}
