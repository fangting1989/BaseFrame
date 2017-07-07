package models

type T_member struct {
	MemberID     int64   `gorm:"column:MemberID;NOT NULL;primary_key;"`
	MemName      string  `gorm:"column:MemName;"`
	MemPwd       string  `gorm:"column:MemPwd;"`
	MemMobile    string  `gorm:"column:MemMobile;"`
	MemOld       string  `gorm:"column:MemOld;"`
	MemAge       string  `gorm:"column:MemAge;"`
	MemMail      string  `gorm:"column:MemMail;"`
	WeChatOpenID string  `gorm:"column:WeChatOpenID;"`
	CostAmount   float64 `gorm:"column:CostAmount;"`
	CostNum      int     `gorm:"column:CostNum;"`
	WeChatImg    string  `gorm:"column:WeChatImg;"`
	WeChatName   string  `gorm:"column:WeChatName;"`
	MemSex       int     `gorm:"column:MemSex;"`
}

func (T_member) TableName() string {
	return "t_member"
}
