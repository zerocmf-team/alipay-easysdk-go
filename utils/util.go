/**
** @创建时间: 2022/6/21 23:28
** @作者　　: return
** @描述　　:
 */

package utils

import (
	"reflect"
	"strings"
)

/**
 * @Author return <1140444693@qq.com>
 * @Description 递归结构体转map[string]string
 * @Date 2022/6/22 0:52:15
 * @Param
 * @return
 **/

func ReflectPtr(ptr interface{}, tag string) (json map[string]string) {

	if json == nil {
		json = make(map[string]string, 0)
	}

	t := reflect.TypeOf(ptr)
	v := reflect.ValueOf(ptr)

	kd := t.Kind()

	if kd == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	num := v.NumField()

	for i := 0; i < num; i++ {
		if v.Field(i).Kind() == reflect.Struct || (v.Field(i).Kind() == reflect.Ptr && v.Field(i).Elem().Kind() == reflect.Struct) {
			jsonChildren := ReflectPtr(v.Field(i).Interface(), "sign")
			for k, v := range jsonChildren {
				json[k] = v
			}
		} else {
			tags := t.Field(i).Tag.Get(tag)
			if tags == "" {
				continue
			}
			tagArr := strings.Split(tags, ",")
			tag := tagArr[0]
			omitempty := false
			if len(tagArr) > 1 {
				if tagArr[1] == "omitempty" {
					omitempty = true
				}
			}
			val := v.Field(i).String()
			if omitempty == true && val == "" {
				continue
			}
			json[tag] = val
		}
	}
	return
}
