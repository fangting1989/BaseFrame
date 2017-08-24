package models
import (
		"time"
		)
type File struct {
  File_id   	int64     	`xorm:" not null  pk BIGINT(20)"`
  NAME      	string    	`xorm:"VARCHAR(500)"`
  EXTENSION 	string    	`xorm:"VARCHAR(50)"`
  SIZE      	int       	`xorm:"INT(20)"`    
  STORAGE_PATH	string    	`xorm:"VARCHAR(500)"`
  Isvalid   	int       	`xorm:"TINYINT(1)"` 
  Createdate	int64     	`xorm:" not null BIGINT(20)"`
}