package ExcelBuilder

import (
	"fmt"
	"reflect"
)

type Builder struct {
}

type IBuildExcel interface {
	//写excel
	BuildExcel([]interface{}) error
}

//生成
func (b *Builder) BuildExcel(list []interface{}) error {
	for _, obj := range list {
		t := reflect.TypeOf(obj)
		v := reflect.ValueOf(obj)
		for i := 0; i < t.NumField(); i++ {
			fmt.Println(fmt.Sprintf("字段名称: %v, 字段值：%v", t.Field(i).Name, v.Field(i)))
		}
	}
	return nil
}
