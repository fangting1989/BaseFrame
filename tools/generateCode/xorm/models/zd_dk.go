package models
type Zd_dk struct {
  JZDSX_GUID	int64     	`xorm:" not null  pk BIGINT(20)"`
  ZD_GUID   	string    	`xorm:"VARCHAR(50)"`
  DKBH      	string    	`xorm:"VARCHAR(50)"`
  DKMC      	string    	`xorm:"VARCHAR(100)"`
  JZDS      	float64   	`xorm:"DECIMAL(10,0)"`
  DKMJ      	float64   	`xorm:"DECIMAL(18,4)"`
  DKYT      	string    	`xorm:"VARCHAR(100)"`
  TFH       	string    	`xorm:"VARCHAR(50)"`
}