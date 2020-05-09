package util

import (
	"fmt"
	"time"
)

func GetID() string {
	t := time.Now()

	year := t.Year()
	month := int(t.Month())
	day := t.Day()
	hour := t.Hour()
	minute := t.Minute()
	second := t.Second()

	return fmt.Sprintf("%d%d%d%d%d%d", year, month, day, hour, minute, second)
}
