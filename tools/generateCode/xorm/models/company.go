package models
import (
		"time"
		)
type Company struct {
  Company_ID	int64     	`xorm:" not null  pk BIGINT(20)"`
  Logo      	string    	`xorm:"VARCHAR(100)"`
  GPDM      	string    	`xorm:"VARCHAR(40)"`
  Qypf      	string    	`xorm:"VARCHAR(20)"`
  QYMC      	string    	`xorm:"VARCHAR(100)"`
  LXDZ      	string    	`xorm:"VARCHAR(200)"`
  LXDH      	string    	`xorm:"VARCHAR(20)"`
  QYYX      	string    	`xorm:"VARCHAR(40)"`
  QYWZ      	string    	`xorm:"VARCHAR(20)"`
  FRDB      	string    	`xorm:"VARCHAR(20)"`
  ZCZB      	string    	`xorm:"VARCHAR(20)"`
  ZCSJ      	time.Time 	`xorm:"DATE"`       
  QYZK      	string    	`xorm:"VARCHAR(40)"`
  GSZCH     	string    	`xorm:"VARCHAR(25)"`
  ZZJGDMZ   	string    	`xorm:"VARCHAR(25)"`
  TYXYDM    	string    	`xorm:"VARCHAR(25)"`
  NSRBSM    	string    	`xorm:"VARCHAR(25)"`
  HY        	string    	`xorm:"VARCHAR(25)"`
  YYQX      	string    	`xorm:"VARCHAR(500)"`
  ZCDZ      	string    	`xorm:"VARCHAR(500)"`
  JYFW      	string    	`xorm:"VARCHAR(500)"`
  HZRQ      	time.Time 	`xorm:"DATETIME"`   
  DJJG      	string    	`xorm:"VARCHAR(50)"`
  BZ        	string    	`xorm:"VARCHAR(500)"`
  SJGS      	int64     	`xorm:"BIGINT(20)"` 
  QYLX      	string    	`xorm:"VARCHAR(100)"`
  SSLX      	string    	`xorm:"VARCHAR(100)"`
  YWLX      	string    	`xorm:"VARCHAR(100)"`
  CPLX      	string    	`xorm:"VARCHAR(100)"`
}