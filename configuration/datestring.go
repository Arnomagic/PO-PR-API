package configuration

import (
	"fmt"
	"time"
)

func DateTimeString(text_Fdate string) string {
	if text_Fdate == "" {
		return ""
	}
	layout := "2006-01-02T15:04:05Z"
	str, err := time.Parse(layout, text_Fdate)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return str.Format("2006-01-02 15:04:05")
}
