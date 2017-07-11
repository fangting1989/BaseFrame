package models
			import (
			"time"
			)
			type imgs struct {
			  Tableid   	int       	`xorm:"not nullINT(11)"`
  Tablename 	string    	`xorm:"VARCHAR(255)"`
  Imgpath   	string    	`xorm:"VARCHAR(255)"`
  Addtime   	time.Time 	`xorm:"DATETIME"`   
}