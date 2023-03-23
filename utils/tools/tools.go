package tools

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"golang.org/x/net/html"
	"log"
	"os"
	"reflect"
	"strings"
)

// IndexOf 查询指定元素在数组中的位置
func IndexOf[T string](source []T, target T) int {
	for index, item := range source {
		if item == target {
			return index
		}
	}
	return -1
}

// QueryColData 获取指定二维数组索引列的数据
func QueryColData(table [][]string, colIndex int) []string {
	data := make([]string, 0)
	for _, row := range table {
		for i := 0; i < len(row); i++ {
			if i == colIndex {
				data = append(data, row[i])
			}
		}
	}
	return data
}

// FindAttrVal 查询指定属性的值
func FindAttrVal(node *html.Node, attrName string) string {
	value, hasAttr := "", false
	for _, attr := range node.Attr {
		if attr.Key == attrName {
			hasAttr = true
			value = attr.Val
		}
	}
	if !hasAttr {
		log.Fatalln("找不到指定属性")
	}
	return value
}

// GetPinyinFirstWord 获取拼音首字母
func GetPinyinFirstWord(words string) string {
	result := ""
	pyArgs := pinyin.NewArgs()
	for _, word := range pinyin.Pinyin(words, pyArgs) {
		for _, w := range word {
			result += w[:1]
		}
	}
	return result
}

// HumpToUnderline 驼峰转下划线
func HumpToUnderline(word string) string {
	result := ""
	for i, w := range word {
		s := strings.ToLower(string(w))
		if i == 0 {
			result += s
			continue
		}
		if 'A' <= w && w <= 'Z' {
			result += "_" + s
		} else {
			result += s
		}
	}
	return result
}

// FileIsExist 判断文件是否存在
func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// ReadFile 读取文件
func ReadFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("读取文件出错 : %s", err)
	}
	return data
}

// WriteFile 写文件
func WriteFile(path string, data []byte) {
	err := os.WriteFile(path, data, 0666)
	if err != nil {
		log.Fatalf("写文件出错 : %s", err)
	}
}

// MakeUpdateSqlAndArgs 构建更新语句和参数
func MakeUpdateSqlAndArgs(data interface{}) (string, []interface{}) {
	rType := reflect.TypeOf(data).Elem()
	rValue := reflect.ValueOf(data).Elem()

	pkgAndName := strings.Split(rType.String(), ".")
	tableName := HumpToUnderline(pkgAndName[len(pkgAndName)-1])

	operCols := make([]string, 0) // 操作列
	operArgs := make([]interface{}, 0)
	whereCols := make([]string, 0) // 条件列
	whereArgs := make([]interface{}, 0)

	for i := 0; i < rType.NumField(); i++ {
		field := rType.Field(i)
		value := rValue.Field(i)

		fieldNamePrefix := field.Name[0]
		if fieldNamePrefix >= 'A' && fieldNamePrefix <= 'Z' {
			dbTag := field.Tag.Get("db")
			if dbTag == "WHERE" {
				whereCols = append(whereCols, field.Name)
				whereArgs = append(whereArgs, value.Interface())
			} else {
				operCols = append(operCols, field.Name)
				operArgs = append(operArgs, value.Interface())
			}
		}
	}

	operSep, whereSep := ", ", " AND " // SQL间隔符
	var sql = ""
	var args = make([]interface{}, 0)

	if len(whereCols) != 0 && len(operCols) != 0 {
		sql += "UPDATE " + tableName
		sql += " SET "
		for _, col := range operCols {
			sql += fmt.Sprintf("%s = ?%s", HumpToUnderline(col), operSep)
		}
		sql = sql[:len(sql)-len(operSep)]
		sql += " WHERE "
		for _, col := range whereCols {
			sql += fmt.Sprintf("%s = ?%s", HumpToUnderline(col), whereSep)
		}
		sql = sql[:len(sql)-len(whereSep)]
	}

	args = append(args, operArgs...)
	args = append(args, whereArgs...)

	return sql, args
}
