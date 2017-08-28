package models
type Qybg struct {
  QYBGID    	int64     	`xorm:" not null  pk BIGINT(20)"`
  Company_ID	int64     	`xorm:"BIGINT(20)"` 
  FILENAME  	string    	`xorm:"VARCHAR(200)"`
  PATH      	string    	`xorm:"VARCHAR(100)"`
}