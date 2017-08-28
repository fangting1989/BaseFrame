package models
type Member struct {
  Member_ID 	int64     	`xorm:" not null  pk BIGINT(20)"`
  Enterprise_ID	string    	`xorm:"VARCHAR(100)"`
  JobName   	string    	`xorm:"VARCHAR(50)"`
  Age       	int       	`xorm:"INT(11)"`    
  PassPort  	string    	`xorm:"VARCHAR(50)"`
  Password  	string    	`xorm:"VARCHAR(50)"`
  Mobile    	string    	`xorm:"VARCHAR(50)"`
  Email     	string    	`xorm:"VARCHAR(20)"`
  Isvalid   	int       	`xorm:"TINYINT(1)"` 
  Nickname  	string    	`xorm:"VARCHAR(50)"`
  Remark    	string    	`xorm:"VARCHAR(500)"`
}