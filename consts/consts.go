package consts

const (
	STATUS_CODE_200 = 200
	STATUS_CODE_400 = 400
	STATUS_CODE_401 = 401
	STATUS_CODE_404 = 404
	STATUS_CODE_500 = 500
)

const (
	ErrorMsgHouseNotExists   = "房源不存在!"
	ErrorMsgAccountNotExists = "用户不存在!"
)

const (
	DefaultPageSize = 20
)

const (
	CHECKING  = "0"
	RENTING   = "1"
	DELETED   = "2"
	CHECKFAIL = "3"
)

var HouseStatusMap = map[string]string{
	CHECKING:  "审核中",
	RENTING:   "出租中",
	DELETED:   "已删除",
	CHECKFAIL: "审核失败",
}

var RentTypeMap = map[string]string{
	"0": "不限",
	"1": "整租",
	"2": "合租",
}

var ApartMentTypeMap = map[string]string{
	"1": "单间",
	"2": "一室一厅",
	"3": "两室一厅",
	"4": "两室两厅",
	"5": "三室一厅",
	"6": "三室两厅",
	"7": "四室一厅",
	"8": "四室两厅",
	"9": "其它",
}
var DEFAULT_NONE_IMAGE string = "https://rent-mini.obs.cn-north-4.myhuaweicloud.com/web/01bc0f59c9a9b0a8012053f85f066c.jpg%40260w_195h_1c_1e_1o_100sh.jpg?versionId=null"

var CityMapSorted = []string{"北京", "上海", "广州", "深圳", "杭州", "成都", "武汉", "长沙", "西安", "郑州", "天津", "厦门"}
var CityMap = map[string]string{
	"北京": "北京市",
	"上海": "上海市",
	"广州": "广东省",
	"深圳": "广东省",
	"杭州": "浙江省",
	"成都": "四川省",
	"武汉": "湖北省",
	"长沙": "湖南省",
	"郑州": "河南省",
	"西安": "陕西省",
	"天津": "天津省",
	"厦门": "福建省",
}

var FacilityMap = map[string]map[string]string{
	"bingxiang": {"name": "冰箱", "icon": "icon-bingxiang"},
	"duwei":     {"name": "独卫", "icon": "icon-matong"},
	"dianti":    {"name": "电梯", "icon": "icon-dianti"},
	"kongtiao":  {"name": "空调", "icon": "icon-kongtiao-"},
	"nuanqi":    {"name": "暖气", "icon": "icon-nuanqi-"},
	"yigui":     {"name": "衣柜", "icon": "icon-yigui"},
	"yangtai":   {"name": "阳台", "icon": "icon-yangtai"},
	"reshuiqi":  {"name": "热水", "icon": "icon-reshuiqi"},
	"zhufan":    {"name": "煮饭", "icon": "icon-Concise"},
	"wifi":      {"name": "无线网", "icon": "icon-wifi"},
	"weibolu":   {"name": "微波炉", "icon": "icon-weibolu-"},
	"xiyiji":    {"name": "洗衣机", "icon": "icon-xiyiji"},
}

var HouseCheckMessage string = "%s 发布了新房源，请注意审核!"
