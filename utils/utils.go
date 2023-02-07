package utils

import (
	"sort"
	"strings"
)

func InStringArray(value string, array []string) bool {
	for _, v := range array {
		if strings.HasPrefix(value, v) {
			return true
		}
		//if v == value {
		//	return true
		//}

	}
	return false
}

func IsInInt64Arrray(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func MapToNameValueList(DictData interface{}, reverseKeyValue bool, excludeKeys []string) interface{} {
	// todo: 这里就是给一个map添加两个key字段，实现有点儿怪...
	keys := []interface{}{}
	insertKeys := func(key string, value interface{}) {
		if reverseKeyValue {
			keys = append(keys, map[string]interface{}{"name": value, "value": key})
		} else {
			keys = append(keys, map[string]interface{}{"name": key, "value": value})
		}
	}
	switch ret := DictData.(type) {
	case map[string]interface{}:
		for key, value := range ret {
			if len(excludeKeys) > 0 && InStringArray(key, excludeKeys) {
				continue
			}
			insertKeys(key, value)
		}
	case map[string]string:
		for key, value := range ret {
			insertKeys(key, value)
		}
	}
	return keys
}

func GetSortedMapKeys(data map[string]string) []string {
	keys := make([]string, 0, len(data))
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func InterFaceListToStringList(data []interface{}) []string {
	array := make([]string, len(data))
	for i, v := range data {
		array[i] = v.(string)
	}
	return array
}
