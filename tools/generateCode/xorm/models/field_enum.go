package models
type Field_enum struct {
  FLD_ENUM_ID	int64     	`xorm:" not null  pk BIGINT(20)"`
  ENUM_GROUP_GUID	int64     	`xorm:"BIGINT(20)"` 
  PARENT_ENUM_VALUE	string    	`xorm:"VARCHAR(255)"`
  ENUM_VALUE	string    	`xorm:"VARCHAR(255)"`
  ENUM_NAME 	string    	`xorm:"VARCHAR(255)"`
  ENUM_LEVEL	string    	`xorm:"VARCHAR(255)"`
  ENUM_DESC 	string    	`xorm:"VARCHAR(255)"`
  ENUM_STATUS	string    	`xorm:"VARCHAR(255)"`
  USE_FREQUENCY	string    	`xorm:"VARCHAR(255)"`
  ENUM_ORDER	string    	`xorm:"VARCHAR(255)"`
  ENUM_VALUEOTHER	string    	`xorm:"VARCHAR(255)"`
  ENUM_NAMEOTHER	string    	`xorm:"VARCHAR(255)"`
  ENUM_VALUEOTHER1	string    	`xorm:"VARCHAR(255)"`
}