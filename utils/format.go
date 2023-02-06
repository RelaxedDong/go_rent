package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func FormatString(template string, items map[string]interface{}) string {
	// node: 未包含在items里面的数据，则不会进行替换
	for key, value := range items {
		template = strings.ReplaceAll(template, fmt.Sprintf("{%v}", key), fmt.Sprintf("%v", value))
	}
	return template
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
