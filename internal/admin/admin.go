package admin

import (
	"strconv"
	"time"
)

type admins []string

func ThisCurrentWeek(admins admins) string {
	cwn := currentWeekNumber()
	return thisWeek(admins, cwn)
}

func thisWeek(admins admins, week int) string {
	return "The *Administrator Of The Week (AOTW)* number *" + strconv.Itoa(week) + "*: <@" + admins[week%len(admins)] + ">"
}

func currentWeekNumber() int {
	currentTime := time.Now().UTC()
	_, week := currentTime.ISOWeek()
	return week
}
