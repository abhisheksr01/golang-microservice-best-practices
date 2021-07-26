package service

import (
	"fmt"
	"time"
)

var Now = time.Now().Hour()

func Greet(now time.Time, userName string) string {
	if currentHour := now.Hour(); currentHour >= 4 && currentHour < 12 {
		return fmt.Sprintf("Hi %s, Good Morning.", userName)
	} else if currentHour >= 12 && currentHour < 17 {
		return fmt.Sprintf("Hi %s, Good Afternoon.", userName)
	} else if currentHour >= 17 && currentHour < 20 {
		return fmt.Sprintf("Hi %s, Good Evening.", userName)
	}
	return fmt.Sprintf("Hi %s, Good Night.", userName)
}
