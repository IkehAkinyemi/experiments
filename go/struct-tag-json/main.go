package main

import (
	"fmt"
	"reflect"
	"strings"
)

type MyStruct struct {
	UserID   int    `json:"user_id" validate:"required,min=1"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email"`
}

func GetJSONTags(s interface{}) []string {
	t := reflect.TypeOf(s)
	if t.Kind() != reflect.Struct {
		return nil
	}

	numFields := t.NumField()
	tags := make([]string, 0, numFields)

	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			jsonTag = strings.Split(jsonTag, ",")[0]
			tags = append(tags, jsonTag)
		}
	}

	return tags
}

func main() {
	myStruct := MyStruct{}
	tags := GetJSONTags(myStruct)
	fmt.Println(tags)
}
