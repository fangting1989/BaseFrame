package models

type Pageobject struct {
	PageNum  int64 `xorm:"not null BIGINT(20)"`
	PageSize int64 `xorm:"BIGINT(20)"`
}
