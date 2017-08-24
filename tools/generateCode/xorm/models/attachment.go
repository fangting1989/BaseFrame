package models
import (
		"time"
		)
type Attachment struct {
  Attachment_id	int64     	`xorm:" not null  pk BIGINT(20)"`
  BUSINESS_TYPE	string    	`xorm:"VARCHAR(100)"`
  BUSINESS_ID	int64     	`xorm:"BIGINT(20)"` 
  BUSINESS_CATEGORY	string    	`xorm:"VARCHAR(100)"`
  File_id   	int64     	`xorm:"BIGINT(20)"` 
  Createdate	time.Time 	`xorm:"DATETIME"`   
  Isvalid   	int       	`xorm:"TINYINT(4)"` 
  Md5code   	string    	`xorm:"VARCHAR(100)"`
}