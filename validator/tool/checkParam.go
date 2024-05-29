package tool

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

/*
说明：结构体属性校验，根据tag为check的值进行校验
格式：<类型：string | uint | int， 校验>
string：正则表达式 | 特殊字符
int：范围 1-100，默认全闭，即[1,100]
uint：范围 1-100，默认全闭，即[1,100]
*/

func CheckParams(obj any) error {
	var objType = reflect.TypeOf(obj)
	// 遍历结构体所有属性
	for i := 0; i < objType.NumField(); i++ {
		// 获取属性名和属性值
		fieldName := objType.Field(i).Name
		data := reflect.ValueOf(obj).Field(i)
		// 判断是否有tag为check的值，没有则跳过
		validate, flag := objType.Field(i).Tag.Lookup("check")
		if !flag {
			continue
		}
		// 根据check的值进行校验
		err := check(fieldName, data, validate)
		if err != nil {
			return err
		}
	}
	return nil
}

func check(fieldName string, data reflect.Value, str string) error {
	// 解析tag值，确保tag值格式正确
	var splits = strings.Split(str, ",")
	if len(splits) != 2 {
		return errors.New(fmt.Sprintf("校验字段:%s，tag `check` 格式错误", fieldName))
	}

	// 根据校验类型选择校验方法
	var tp = splits[0]
	switch tp {
	case "string":
		return checkString(fieldName, data.String(), splits[1])
	case "int":
		return checkInt(fieldName, data.Int(), splits[1])
	case "uint":
		return checkUint(fieldName, data.Uint(), splits[1])
	default:
		return errors.New(fmt.Sprintf("校验字段:%s，暂只支持string | int | uint 类型的校验", fieldName))
	}
}

func checkString(fieldName string, data string, regex string) error {
	// 判断特殊字段
	// required：不能为空
	if regex == "required" {
		if data == "" {
			return errors.New(fmt.Sprintf("校验字段:%s，不能为空", fieldName))
		}
	}
	// email：邮箱
	if regex == "email" {
		regex = `^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$`
	}
	// phone：手机号
	if regex == "phone" {
		regex = `^1[3|4|5|7|8][0-9]\d{8}$`
	}

	// 正则表达式校验
	match := regexp.MustCompile(regex).MatchString(data)
	if !match {
		return errors.New(fmt.Sprintf("校验字段:%s，数据：%s，格式错误", fieldName, data))
	}
	return nil
}

func checkInt(fieldName string, data int64, regex string) error {
	// 格式1-100
	var splits = strings.Split(regex, "-")
	if len(splits) != 2 {
		return errors.New(fmt.Sprintf("校验字段:%s，格式错误,举例：int,1-100", fieldName))
	}
	var minV, _ = strconv.Atoi(splits[0])
	var maxV, _ = strconv.Atoi(splits[1])
	if data < int64(minV) || data > int64(maxV) {
		return errors.New(fmt.Sprintf("校验字段:%s，数据：%d，不在范围(%d,%d)内", fieldName, data, minV, maxV))
	}
	return nil
}

func checkUint(fieldName string, data uint64, regex string) error {
	// 格式1-100
	var splits = strings.Split(regex, "-")
	if len(splits) != 2 {
		return errors.New(fmt.Sprintf("校验字段:%s，格式错误,举例：uint,1-100", fieldName))
	}
	var minV, _ = strconv.Atoi(splits[0])
	var maxV, _ = strconv.Atoi(splits[1])
	if data < uint64(minV) || data > uint64(maxV) {
		return errors.New(fmt.Sprintf("校验字段:%s，数据：%d，不在范围(%d,%d)内", fieldName, data, minV, maxV))
	}
	return nil
}
