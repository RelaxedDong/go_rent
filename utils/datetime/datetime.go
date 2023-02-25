package datetime

import (
	"fmt"
	"time"
)

func TimeToStr(t time.Time) (timeStr string) {
	return t.Format("2006-01-02 15:04:05")
}

func GetLastLoginActiveText(lastLogin time.Time) string {
	Duration := time.Now().Sub(lastLogin)
	Minutes, Hour := Duration.Minutes(), Duration.Hours()
	switch {
	case Minutes < 20:
		return "刚刚来过"
	case Minutes < 60:
		return fmt.Sprintf("%d 分钟前来过", int64(Minutes))
	case Hour < 24:
		return fmt.Sprintf("%d 小时前来过", int64(Hour))
	default:
		return fmt.Sprintf("%d 天前来过", int64(Hour/24))
	}
}
