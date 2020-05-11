package util

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
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

func FileHasID(filename string) bool {
	pattern := ".+-[0-9]{8}[0-9]+"
	match, _ := regexp.MatchString(pattern, filename)
	return match
}

func AddFileID(filename string) string {
	id := GetID()
	// if extension exists, add ID before extension
	ext := filepath.Ext(filename)
	filenameNoExt := strings.Replace(filename, ext, "", -1)
	return fmt.Sprintf("%s-%s%s", filenameNoExt, id, ext)
}
