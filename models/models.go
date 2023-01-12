package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type BaseModel struct {
	// https://segmentfault.com/a/1190000041558761?sort=votes
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedTime time.Time `orm:"auto_now;type(datetime)"`
}

type AccountModel struct {
	BaseModel
	Id        int64  `orm:"pk;auto"`
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

	LastLogin time.Time `orm:"type(datetime)"`
}

type HouseModel struct {
	Id       int64  `orm:"pk;auto"`
	Title    string `orm:"size(150)"`
	Desc     string `orm:"type(text);null"`
	Address  string `orm:"size(150)"`
	Region   string `orm:"size(100)"`
	Province string `orm:"size(100)"`
	City     string `orm:"size(100)"`
	Subway   string `orm:"size(100);default();"`
	Bus      string `orm:"size(200);default();null"`
	// array json数组
	Facilities string        `orm:"type(text);"`
	Publisher  *AccountModel `orm:"rel(fk)"`
	// array json数组
	Imgs      string `orm:"type(text);"`
	Area      uint32 `orm:"null"`
	VideoUrl  string `orm:"size(200);null"`
	Storey    uint8  `orm:"null"`
	Longitude string
	Latitude  string
	//状态相关
	ShowPhone bool   `orm:"default(true)"`
	ViewCount uint32 `orm:"null"`
	// 0: 审核 1：出租 2：已删除 3：审核失败
	Status       string `orm:"size(1);default(0)"`
	FailReason   string `orm:"size(100);default();"`
	CanShortRent bool   `orm:"default(false)"`
	IsDelicate   bool   `orm:"default(false)"`
	// 类型（0不限，1整租，2合租）
	HouseType string `orm:"size(1);default(0)"`
	// 1-9 ['单间', '一室一厅', '两室一厅', '两室两厅', '三室一厅', '三室两厅', '四室一厅', '四室两厅', '其它']
	Apartment string `orm:"size(1);default(1)"`
	// 存放json数组
	ExtraInfo string `orm:"type(text);"`
	BaseModel
}

func (self *AccountModel) TableName() string {
	return "account" // 自定义表名
}

func (self *HouseModel) TableName() string {
	return "house"
}

// 创建索引
func (self *HouseModel) TableIndex() [][]string {
	return [][]string{
		{"title"},
	}
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
	orm.RegisterModel(new(HouseModel))
	//生成表
	//自动创建表 参数二为是否开启创建表(如果值为ture时，表已经存在并且表中有值的情况下，它会先删除我们的表，然后重新创建)   参数三是否更新表
	orm.RunSyncdb("default", false, true)
	OrmManager = orm.NewOrm()
}
