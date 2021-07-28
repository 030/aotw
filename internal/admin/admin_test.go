package admin

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThisWeek(t *testing.T) {
	m := make(map[string]int)
	m["piet"] = 0
	m["klaas"] = 1
	m["jan"] = 2
	m["mo"] = 3
	m["piet"] = 4
	m["klaas"] = 5
	m["jan"] = 6
	m["mo"] = 7

	admins := admins{"piet", "klaas", "jan", "mo"}

	for admin, weekNumber := range m {
		assert.Equal(t, "The *Administrator Of The Week (AOTW)* number *"+strconv.Itoa(weekNumber)+"*: <@"+admin+">", thisWeek(admins, weekNumber))
	}
}
