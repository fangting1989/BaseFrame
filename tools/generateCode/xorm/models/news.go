package models
			import (
			"time"
			)
			type news struct {
			  Newsid    	int       	`xorm:"not null pk INT(11)"`
  Newstitle 	string    	`xorm:"VARCHAR(255)"`
  Newsubtitle	string    	`xorm:"VARCHAR(255)"`
  Newscontent	string    	`xorm:"TEXT"`       
  Newsimgs  	string    	`xorm:"VARCHAR(255)"`
  Newsdesc  	string    	`xorm:"VARCHAR(255)"`
  Addtime   	time.Time 	`xorm:"DATETIME"`   
}