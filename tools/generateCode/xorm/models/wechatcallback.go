package models
			import (
			"time"
			)
			type wechatcallback struct {
			  Appid     	string    	`xorm:"VARCHAR(50)"`
  Attach    	string    	`xorm:"not null pk VARCHAR(100)"`
  Cash_fee  	string    	`xorm:"VARCHAR(25)"`
  Fee_type  	string    	`xorm:"VARCHAR(15)"`
  Is_subscribe	string    	`xorm:"VARCHAR(20)"`
  Mch_id    	string    	`xorm:"VARCHAR(20)"`
  Nonce_str 	string    	`xorm:"VARCHAR(50)"`
  Openid    	string    	`xorm:"VARCHAR(50)"`
  Out_trade_no	string    	`xorm:"VARCHAR(32)"`
  Result_code	string    	`xorm:"VARCHAR(50)"`
  Return_code	string    	`xorm:"VARCHAR(50)"`
  Total_fee 	string    	`xorm:"VARCHAR(25)"`
  Time_end  	string    	`xorm:"VARCHAR(50)"`
  Trade_type	string    	`xorm:"VARCHAR(50)"`
  Transaction_id	string    	`xorm:"VARCHAR(50)"`
}