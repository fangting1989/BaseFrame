package models
import (
		"time"
		)
type Project struct {
  PROJECTID 	int64     	`xorm:" not null  pk BIGINT(20)"`
  XMMC      	string    	`xorm:"VARCHAR(100)"`
  XMGS      	string    	`xorm:"VARCHAR(100)"`
  SSQY      	string    	`xorm:"VARCHAR(100)"`
  PROVINCE  	string    	`xorm:"VARCHAR(50)"`
  CITY      	string    	`xorm:"VARCHAR(50)"`
  COUNTY    	string    	`xorm:"VARCHAR(50)"`
  LPDZ      	string    	`xorm:"VARCHAR(100)"`
  SYNX      	string    	`xorm:"VARCHAR(50)"`
  ZJMJ      	float64   	`xorm:"DECIMAL(18,3)"`
  RJL       	float64   	`xorm:"DECIMAL(18,3)"`
  LHL       	float64   	`xorm:"DECIMAL(18,3)"`
  GHHS      	int       	`xorm:"INT(11)"`    
  ZXZT      	string    	`xorm:"VARCHAR(50)"`
  WYLX      	string    	`xorm:"VARCHAR(50)"`
  WYGS      	string    	`xorm:"VARCHAR(100)"`
  WYF       	float64   	`xorm:"DECIMAL(18,3)"`
  KPSJ      	time.Time 	`xorm:"DATE"`       
  XSZT      	string    	`xorm:"VARCHAR(50)"`
  SLBDZ     	string    	`xorm:"VARCHAR(100)"`
  SLBDH     	string    	`xorm:"VARCHAR(100)"`
  YSZH      	string    	`xorm:"VARCHAR(100)"`
  YSZSLSJ   	time.Time 	`xorm:"DATE"`       
  YSZTS     	int       	`xorm:"INT(11)"`    
  YSZMJ     	float64   	`xorm:"DECIMAL(18,3)"`
  ZZMJ      	float64   	`xorm:"DECIMAL(18,3)"`
  SBMJ      	float64   	`xorm:"DECIMAL(18,3)"`
  YSKSTS    	int       	`xorm:"INT(11)"`    
  KSMJ      	float64   	`xorm:"DECIMAL(18,3)"`
  CJYF      	string    	`xorm:"VARCHAR(50)"`
  XSTS      	string    	`xorm:"VARCHAR(50)"`
  XSMJ      	float64   	`xorm:"DECIMAL(18,3)"`
  XSJE      	float64   	`xorm:"DECIMAL(18,3)"`
  XSJJ      	float64   	`xorm:"DECIMAL(18,3)"`
  LPZMC     	string    	`xorm:"VARCHAR(100)"`
  YSTS      	int       	`xorm:"INT(11)"`    
  CJKSTS    	int       	`xorm:"INT(11)"`    
}