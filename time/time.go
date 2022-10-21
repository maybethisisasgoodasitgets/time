package time

import "time"

func GetNowUnixTime()string{

	return string(rune(time.Now().Unix()))
}