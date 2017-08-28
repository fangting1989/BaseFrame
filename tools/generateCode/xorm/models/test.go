package models
type Test struct {
  ID        	int64     	`xorm:" not null  pk BIGINT(20)"`
  NAME      	string    	`xorm:"VARCHAR(255)"`
}