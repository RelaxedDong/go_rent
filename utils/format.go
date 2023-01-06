package utils

import (
	"fmt"
	"strings"
)

func FormatString(template string, items map[string]interface{}) string {
	// node: 未包含在items里面的数据，则不会进行替换
	for key, value := range items {
		template = strings.ReplaceAll(template, fmt.Sprintf("{%v}", key), fmt.Sprintf("%v", value))
	}
	return template
}
