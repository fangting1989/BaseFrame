package models
import (
		"time"
		)
type gygg struct {
  GYGG_ID   	int64     	`xorm:"not nullBIGINT(20)"`
  LAND_ID   	string    	`xorm:"CHAR(10)"`   
  GYGG_GUID 	string    	`xorm:"VARCHAR(50)"`
  GYGG_BT   	string    	`xorm:"VARCHAR(255)"`
  JZ_RQ     	time.Time 	`xorm:"DATE"`       
  GYGG_NR   	[]uint8   	`xorm:"BLOB"`       
  XZQ_DM    	string    	`xorm:"VARCHAR(9)"` 
  SH_ZT     	string    	`xorm:"VARCHAR(2)"` 
  FB_SJ     	time.Time 	`xorm:"DATE"`       
  JZ_RQ_END 	time.Time 	`xorm:"DATE"`       
  GG_LX     	string    	`xorm:"VARCHAR(50)"`
  SH_SJ     	time.Time 	`xorm:"DATE"`       
  GG_BH     	string    	`xorm:"VARCHAR(50)"`
  LQ_WJ_SJ_S	time.Time 	`xorm:"DATE"`       
  LQ_WJ_SJ_E	time.Time 	`xorm:"DATE"`       
  LQ_WJ_DD  	string    	`xorm:"VARCHAR(300)"`
  TJ_SQ_SJ_S	time.Time 	`xorm:"DATE"`       
  TJ_SQ_SJ_E	time.Time 	`xorm:"DATE"`       
  TJ_SQ_DD  	string    	`xorm:"VARCHAR(300)"`
  JN_BZJ_JZ_SJ	time.Time 	`xorm:"DATE"`       
  QR_TB_ZG_SJ	time.Time 	`xorm:"DATE"`       
  ZB_SJ_S   	time.Time 	`xorm:"DATE"`       
  ZB_SJ_E   	time.Time 	`xorm:"DATE"`       
  CR_SJ     	time.Time 	`xorm:"DATE"`       
  CR_DD     	string    	`xorm:"VARCHAR(300)"`
  KB_SJ     	time.Time 	`xorm:"DATE"`       
  QT_SX     	string    	`xorm:"VARCHAR(4000)"`
  DK_JJ     	string    	`xorm:"VARCHAR(500)"`
}