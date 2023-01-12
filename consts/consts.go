package consts

const (
	STATUS_CODE_200 = 200
	STATUS_CODE_400 = 400
	STATUS_CODE_401 = 401
	STATUS_CODE_404 = 404
	STATUS_CODE_500 = 500
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
