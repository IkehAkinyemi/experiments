package main

import (
	"fmt"
	"regexp"
	"strings"
)

// prepareIndexingText takes care of inputs with dates and versions and makes them indexable
func prepareIndexingText(name string) string {
	versionsAndDates := []string{}
	versionRegex := regexp.MustCompile(`(v\d+\.?\d*\.?\d*\.?\d*(\.?\d*)*(.*beta)?(.*alpha)?)`)
	dateRegex := regexp.MustCompile(`(\b(0?[1-9]|[12]\d|30|31)[^\w\d\r\n:](0?[1-9]|1[0-2])[^\w\d\r\n:](\d{4}|\d{2})\b)|(\b(0?[1-9]|1[0-2])[^\w\d\r\n:](0?[1-9]|[12]\d|30|31)[^\w\d\r\n:](\d{4}|\d{2})\b)`)

	foundVersions := versionRegex.FindAllString(name, -1)
	if len(foundVersions) > 0 {
		cleanVer := foundVersions[0]
		if cleanVer[len(cleanVer)-1] == '.' {
			cleanVer = cleanVer[:len(cleanVer)-1]
		}
		versionsAndDates = append(versionsAndDates, cleanVer)
		name = strings.ReplaceAll(name, foundVersions[0], " ")
	}

	foundDates := dateRegex.FindAllString(name, -1)
	if len(foundDates) > 0 {
		cleanDate := foundDates[0]
		if cleanDate[len(cleanDate)-1] == '.' {
			cleanDate = cleanDate[:len(cleanDate)-1]
		}
		versionsAndDates = append(versionsAndDates, cleanDate)
		name = strings.ReplaceAll(name, foundDates[0], " ")
	}

	m := regexp.MustCompile(`[\&\'\"\:\*\?\~]`)
	n := regexp.MustCompile(`[\.\+\-\_\@\{\}\(\)\<\>]`)
	o := regexp.MustCompile(`\s\s+`)

	name = m.ReplaceAllString(name, "")
	name = n.ReplaceAllString(name, " ")
	name += " " + strings.Join(versionsAndDates, " ")
	name = o.ReplaceAllString(name, " ")

	return strings.TrimSpace(name)
}

func main() {
	fmt.Println(prepareIndexingText("Extract the first date found and assign it to the cleanDate variable."))
}