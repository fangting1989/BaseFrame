package models
type Companyinfo struct {
  CompanyInfoID	int64     	`xorm:" not null  pk BIGINT(20)"`
  Company_ID	int64     	`xorm:"BIGINT(20)"` 
  DSZ       	string    	`xorm:"VARCHAR(20)"`
  DSZLXFS   	string    	`xorm:"VARCHAR(50)"`
  DSRYMC    	string    	`xorm:"VARCHAR(50)"`
  JSRYMC    	string    	`xorm:"VARCHAR(50)"`
  ZJLMC     	string    	`xorm:"VARCHAR(50)"`
  GDJBXX    	string    	`xorm:"VARCHAR(50)"`
}