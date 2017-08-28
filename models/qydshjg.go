package models
type Qydshjg struct {
  QYDSHJGID 	int64     	`xorm:" not null  pk BIGINT(20)"`
  Company_ID	int64     	`xorm:"BIGINT(20)"` 
  DSMC      	string    	`xorm:"VARCHAR(50)"`
  ZW        	string    	`xorm:"VARCHAR(20)"`
  XL        	string    	`xorm:"VARCHAR(20)"`
  NL        	string    	`xorm:"VARCHAR(20)"`
  REMARK    	string    	`xorm:"VARCHAR(500)"`
}