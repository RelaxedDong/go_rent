package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type AccountModel struct {
	// https://segmentfault.com/a/1190000041558761?sort=votes
	Id        int    `orm:"pk;auto"`
	OpenId    string `orm:"size(50);unique;column(openid)"`
	NickName  string `orm:"size(100);column(nickname)"`
	AvatarUrl string `orm:"size(255);column(avatar_url)"`
	Country   string `orm:"size(30);"`
	Province  string `orm:"size(30);"`
	City      string `orm:"size(30);"`
	Gender    string `orm:"size(1);default(0)"`
	Status    string `orm:"size(1);default(0)"`

	SessionKey time.Time `orm:"type(datetime);null"`
	Phone      string    `orm:"size(11);null;default()"`
	Wechat     string    `orm:"size(50);null;default()"`

	LastLogin   time.Time `orm:"type(datetime)"`
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedTime time.Time `orm:"auto_now;type(datetime)"`
}

func (u *AccountModel) TableName() string {
	return "account" // 自定义表名
}

// OrmManager 未初始化的标准变量定义格式
var OrmManager orm.Ormer

//Article -> article
//AuthUser -> auth_user
//Auth_User -> auth__user 两个下划线
//DB_AuthUser -> d_b__auth_user

func InitOrmConfig() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// todo: 提取配置文件
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/go_rent_backend?charset=utf8", 30)
}

func init() {
	InitOrmConfig()
	//创建表
	fmt.Println("run 创建表...")
	orm.RegisterModel(new(AccountModel))
	//生成表
	//自动创建表 参数二为是否开启创建表(如果值为ture时，表已经存在并且表中有值的情况下，它会先删除我们的表，然后重新创建)   参数三是否更新表
	orm.RunSyncdb("default", false, true)
	OrmManager = orm.NewOrm()
}
