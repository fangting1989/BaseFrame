package models

import (
	"time"
)

type TBillcomment struct {
	Billcommentid int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Content       string    `xorm:"TEXT"`
	Point         string    `xorm:"DECIMAL(6,2)"`
	Comperson     string    `xorm:"VARCHAR(50)"`
	Mobile        string    `xorm:"VARCHAR(50)"`
	Personcode    string    `xorm:"VARCHAR(50)"`
	Membillid     int64     `xorm:"BIGINT(20)"`
	Dealconent    string    `xorm:"TEXT"`
	State         int       `xorm:"INT(11)"`
	Commentdate   time.Time `xorm:"DATETIME"`
	Dealdate      time.Time `xorm:"DATETIME"`
	Pid           int64     `xorm:"BIGINT(20)"`
}

type TImgs struct {
	Tableid   int       `xorm:"not null index(index_id_name) INT(11)"`
	Tablename string    `xorm:"index(index_id_name) VARCHAR(255)"`
	Imgpath   string    `xorm:"VARCHAR(255)"`
	Addtime   time.Time `xorm:"DATETIME"`
}

type TMember struct {
	Memberid     int64  `xorm:"not null pk autoincr BIGINT(20)"`
	Memname      string `xorm:"VARCHAR(50)"`
	Mempwd       string `xorm:"VARCHAR(50)"`
	Memmobile    string `xorm:"VARCHAR(20)"`
	Memold       string `xorm:"VARCHAR(50)"`
	Memage       string `xorm:"VARCHAR(50)"`
	Memmail      string `xorm:"VARCHAR(50)"`
	Wechatopenid string `xorm:"VARCHAR(50)"`
	Costamount   string `xorm:"DECIMAL(10,2)"`
	Costnum      int    `xorm:"INT(11)"`
	Wechatimg    string `xorm:"VARCHAR(200)"`
	Wechatname   string `xorm:"VARCHAR(80)"`
	Memsex       int    `xorm:"INT(11)"`
}

type TMembill struct {
	Membillid     int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Membillamount string    `xorm:"DECIMAL(10,2)"`
	Billcosttime  time.Time `xorm:"DATE"`
	Billpaytime   time.Time `xorm:"DATE"`
	Billpaycode   string    `xorm:"VARCHAR(50)"`
	Billpaytype   int       `xorm:"INT(11)"`
	Billpoint     string    `xorm:"DECIMAL(10,2)"`
	Memid         int64     `xorm:"BIGINT(20)"`
	Membillguid   string    `xorm:"VARCHAR(40)"`
	Paystate      int       `xorm:"INT(11)"`
}

type TModule struct {
	Moduleid         int    `xorm:"not null pk autoincr INT(11)"`
	Moduletypeid     int    `xorm:"INT(11)"`
	Modulename       string `xorm:"VARCHAR(50)"`
	Moduleindex      int    `xorm:"INT(11)"`
	Valid            int    `xorm:"INT(11)"`
	Moduleicon       string `xorm:"VARCHAR(50)"`
	Moduleurl        string `xorm:"VARCHAR(50)"`
	Modulepermission string `xorm:"VARCHAR(50)"`
	Moduledesc       string `xorm:"VARCHAR(200)"`
}

type TModuletype struct {
	Moduletypeid    int    `xorm:"not null pk autoincr INT(11)"`
	Moduletypename  string `xorm:"VARCHAR(50)"`
	Moduletypeindex int    `xorm:"INT(11)"`
	Moduletypepid   int    `xorm:"INT(11)"`
	Valid           string `xorm:"DECIMAL(1)"`
	Moduletypeicon  string `xorm:"VARCHAR(50)"`
	Moduletypedesc  string `xorm:"VARCHAR(200)"`
}

type TNews struct {
	Newsid      int       `xorm:"not null INT(11)"`
	Newstitle   string    `xorm:"VARCHAR(255)"`
	Newsubtitle string    `xorm:"VARCHAR(255)"`
	Newscontent string    `xorm:"TEXT"`
	Newsimgs    string    `xorm:"VARCHAR(255)"`
	Newsdesc    string    `xorm:"VARCHAR(255)"`
	Addtime     time.Time `xorm:"DATETIME"`
}

type TProd struct {
	Prodid        int       `xorm:"not null pk autoincr INT(11)"`
	Prodname      string    `xorm:"VARCHAR(255)"`
	Prodenname    string    `xorm:"VARCHAR(255)"`
	Prodshortinfo string    `xorm:"VARCHAR(255)"`
	Prodinfo      string    `xorm:"TEXT"`
	Prodkeyword   string    `xorm:"VARCHAR(255)"`
	Prodprice     string    `xorm:"DECIMAL(10,2)"`
	Prodcount     string    `xorm:"VARCHAR(255)"`
	Prodtags      string    `xorm:"VARCHAR(255)"`
	Proddesc      string    `xorm:"VARCHAR(255)"`
	Addtime       time.Time `xorm:"DATETIME"`
	State         int       `xorm:"INT(11)"`
	Wechatshare   string    `xorm:"VARCHAR(255)"`
}

type TRole struct {
	Roleid         int       `xorm:"not null pk autoincr INT(11)"`
	Rolename       string    `xorm:"VARCHAR(50)"`
	Rolevalid      int       `xorm:"INT(11)"`
	Rolecreatedate time.Time `xorm:"DATE"`
	Roledesc       string    `xorm:"VARCHAR(200)"`
}

type TRolemodule struct {
	Rolemoduleid     int    `xorm:"not null pk autoincr INT(11)"`
	Roleid           int    `xorm:"INT(11)"`
	Moduleid         int    `xorm:"INT(11)"`
	Modulepermission string `xorm:"VARCHAR(200)"`
}

type TSmsmessage struct {
	Smsmessageid int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Content      string    `xorm:"VARCHAR(255)"`
	Phone        string    `xorm:"VARCHAR(20)"`
	Sendtime     time.Time `xorm:"DATETIME"`
	State        int       `xorm:"TINYINT(4)"`
	Type         int       `xorm:"TINYINT(4)"`
}

type TUser struct {
	Userid      int    `xorm:"not null pk autoincr INT(11)"`
	Username    string `xorm:"VARCHAR(50)"`
	Usermobile  string `xorm:"VARCHAR(20)"`
	Userstatus  string `xorm:"DECIMAL(1)"`
	Createdate  string `xorm:"VARCHAR(30)"`
	Usermail    string `xorm:"VARCHAR(50)"`
	Useraddress string `xorm:"VARCHAR(200)"`
	Userpwd     string `xorm:"VARCHAR(50)"`
}

type TUserrole struct {
	Roleid     int `xorm:"INT(11)"`
	Userid     int `xorm:"INT(11)"`
	Userroleid int `xorm:"not null pk autoincr INT(11)"`
}

type TWechatcallback struct {
	Appid         string `xorm:"VARCHAR(50)"`
	Attach        string `xorm:"not null pk VARCHAR(100)"`
	CashFee       string `xorm:"VARCHAR(25)"`
	FeeType       string `xorm:"VARCHAR(15)"`
	IsSubscribe   string `xorm:"VARCHAR(20)"`
	MchId         string `xorm:"VARCHAR(20)"`
	NonceStr      string `xorm:"VARCHAR(50)"`
	Openid        string `xorm:"VARCHAR(50)"`
	OutTradeNo    string `xorm:"VARCHAR(32)"`
	ResultCode    string `xorm:"VARCHAR(50)"`
	ReturnCode    string `xorm:"VARCHAR(50)"`
	TotalFee      string `xorm:"VARCHAR(25)"`
	TimeEnd       string `xorm:"VARCHAR(50)"`
	TradeType     string `xorm:"VARCHAR(50)"`
	TransactionId string `xorm:"VARCHAR(50)"`
}
