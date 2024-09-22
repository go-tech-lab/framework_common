package util

import (
	"fmt"
	"reflect"
	"strings"
)

func List2Markdown(list []interface{}, ignoreTableHeaders ...bool) string {
	ignoreTableHeader := false
	if len(ignoreTableHeaders) > 0 {
		ignoreTableHeader = ignoreTableHeaders[0]
	}
	fieldNum := reflect.TypeOf(list[0]).NumField()
	tableContent := strings.Builder{}
	//组件表头
	tableHeader := "| "
	tableHeaderPartition := "| "
	for i := 0; i < fieldNum; i++ {
		targetField := reflect.TypeOf(list[0]).Field(i)
		tableHeader += " " + targetField.Name + " |"
		tableHeaderPartition += " " + " :-- |"
	}
	if !ignoreTableHeader {
		tableContent.WriteString(tableHeader)
		tableContent.WriteString("\n")
		tableContent.WriteString(tableHeaderPartition)
		tableContent.WriteString("\n")
	}
	for _, item := range list {
		tableRow := "| "
		for i := 0; i < fieldNum; i++ {
			fieldValue := reflect.ValueOf(item).Field(i).Interface()
			fieldString := fmt.Sprintf("%v", fieldValue)
			tableRow += " " + fieldString + " |"
		}
		tableContent.WriteString(tableRow)
		tableContent.WriteString("\n")
	}
	return tableContent.String()
}
