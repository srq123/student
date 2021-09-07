package datamodels

type Student struct {
	Id          int    ` xorm:"pk"`
	Age         int    ` xorm:"int"`
	UserName    string ` xorm:"DEFAULT '' varchar(45)"`
	StudentName string ` xorm:"DEFAULT '' varchar(45)"`
	Sex         string ` xorm:"DEFAULT '' varchar(45)"`
	SysCreated  string `xorm:"not null default 0 comment('创建时间') varchar(45)"`
	SysUpdated  string `xorm:"not null default 0 comment('修改时间') varchar(45)"`
}

func (Student) TableName() string {
	return "t_student"
}
