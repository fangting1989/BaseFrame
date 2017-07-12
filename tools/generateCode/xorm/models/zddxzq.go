package models
import (
		"time"
		)
type zddxzq struct {
  ZDDXZQ_ID 	int64     	`xorm:"not nullBIGINT(20)"`
  ZDDXZQ_GUID	string    	`xorm:"VARCHAR(50)"`
  ZD_GUID   	string    	`xorm:"VARCHAR(50)"`
  GD_GUID   	string    	`xorm:"VARCHAR(50)"`
  XZQ_DM    	string    	`xorm:"VARCHAR(50)"`
  ZD_MJ     	float64   	`xorm:"DECIMAL(18,6)"`
}