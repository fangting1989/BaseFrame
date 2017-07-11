package models
			import (
			"time"
			)
			type prod struct {
			  Prodid    	int       	`xorm:"not null pk INT(11)"`
  Prodname  	string    	`xorm:"VARCHAR(255)"`
  Prodenname	string    	`xorm:"VARCHAR(255)"`
  Prodshortinfo	string    	`xorm:"VARCHAR(255)"`
  Prodinfo  	string    	`xorm:"TEXT"`       
  Prodkeyword	string    	`xorm:"VARCHAR(255)"`
  Prodprice 	float64   	`xorm:"DECIMAL(10,2)"`
  Prodcount 	string    	`xorm:"VARCHAR(255)"`
  Prodtags  	string    	`xorm:"VARCHAR(255)"`
  Proddesc  	string    	`xorm:"VARCHAR(255)"`
  Addtime   	time.Time 	`xorm:"DATETIME"`   
  State     	int       	`xorm:"INT(11)"`    
  Wechatshare	string    	`xorm:"VARCHAR(255)"`
}