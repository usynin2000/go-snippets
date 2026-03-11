package main

import (
	"fmt"
	"reflect"
)

type FieldInfo struct {
	Field string
	Label string
	Help  string
}

type Form struct {
	Username string `label:"Username" help:"Enter your username"`
	Age 	 int 	`label:"Age" help:"Enter your age"`
}

func ParseTags(v interface{}) []FieldInfo {
	t := reflect.TypeOf(v)
	var fields []FieldInfo

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fields = append(fields, FieldInfo{
			Field: f.Name,
			Label: f.Tag.Get("label"),
			Help: f.Tag.Get("help"),
		})
	}
	return fields
}

func main() {
	info := ParseTags(Form{})
	for _, f := range info {
		fmt.Printf("%s -> %s (%s)\n", f.Field, f.Label, f.Help)
	}
}

// This way you can build dynamic forms, UI, documentation and etc.
