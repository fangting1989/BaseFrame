package models

type Cjgs_zd struct {
	CJGS_ZD_ID   int64   `xorm:"not nullBIGINT(20)"`
	LAND_ID      string  `xorm:"CHAR(10)"`
	CJGS_ZD_GUID string  `xorm:"VARCHAR(50)"`
	CJGS_GUID    string  `xorm:"VARCHAR(50)"`
	SR_DW        string  `xorm:"VARCHAR(100)"`
	XM_MC        string  `xorm:"VARCHAR(100)"`
	TD_ZL        string  `xorm:"VARCHAR(150)"`
	TD_MJ        float64 `xorm:"DECIMAL(18,6)"`
	TD_YT        string  `xorm:"VARCHAR(50)"`
	TD_SY_TJ     string  `xorm:"VARCHAR(4000)"`
	CR_NX        string  `xorm:"VARCHAR(50)"`
	CJ_JG        float64 `xorm:"DECIMAL(18,6)"`
	ZD_BH        string  `xorm:"VARCHAR(50)"`
	HAVE_YCDK    string  `xorm:"VARCHAR(50)"`
}
