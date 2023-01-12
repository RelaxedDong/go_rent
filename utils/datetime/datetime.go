package datetime

import "time"

func DateTimeToStr(t time.Time) (timeStr string) {
	return t.Format("2006-01-02 15:04:05")
}

//def datetime_to_str(date, date_format=FORMAT_DATETIME, process_none=False):
//    """
//    convert {@see datetime} into date string ('2011-01-12')
//    """
//    if process_none and date is None:
//        return ''
//    return date.strftime(date_format)
