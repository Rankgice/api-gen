package utils

import (
	"database/sql"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

// ToCamelCase 将下划线命名的字符串转换为小驼峰命名
func ToCamelCase(s string) string {
	parts := strings.Split(s, "_")
	for i := 1; i < len(parts); i++ {
		parts[i] = cases.Title(language.English).String(parts[i])
	}
	return strings.Join(parts, "")
}

// ToPascalCase 将下划线命名的字符串转换为大驼峰命名
func ToPascalCase(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		parts[i] = cases.Title(language.English).String(parts[i])
	}
	return strings.Join(parts, "")
}
func ParseDatabaseName(dsn string) string {
	// 找到最后一个斜杠的位置
	lastSlashIndex := strings.LastIndex(dsn, "/")
	if lastSlashIndex == -1 {
		return ""
	}
	// 提取斜杠后面的部分
	databaseName := dsn[lastSlashIndex+1:]
	return databaseName
}

func MysqlTypeToGOType(mysqlType string) string {
	switch mysqlType {
	case "int", "tinyint", "smallint", "mediumint", "bigint":
		return "int64"
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext":
		return "string"
	case "float", "double":
		return "float64"
	case "decimal":
		return "string"
	case "date", "datetime", "timestamp":
		return "time.Time"
	case "boolean":
		return "bool"
	default:
		return "interface{}"
	}
}
func GenerateField(field, dataType, nullable, key string, val sql.NullString, extra, comment string) string {
	FieldName := ToPascalCase(field)
	fieldType := MysqlTypeToGOType(dataType)
	return fmt.Sprintf("\t%s\t%s\t`json:\"%s,optional\" form:\"%s,optional\"`\n", FieldName, fieldType, field, field)
}

//return fmt.Sprintf(`%s     %s  ``json:"%s,optional" form:"%s,optional"```, FieldName,fieldType,field,field)

// ReplaceTemplate 匹配字符串中形如 #{xxx} 的值，
// 并将replaceMap中对应的值填入其中
// 并且确保这些值不被单引号或双引号包裹。
// result返回的是填充后的结果字符串
// 时间复杂度 O(n)
func ReplaceTemplate(input string, replaceMap map[string]string) (result string) {
	nextEnd := 0
	inputRune := []rune(input)
	inputLen := len(inputRune)
	l := 0
	for i := 0; i < inputLen; i++ {
		curRune := inputRune[i]
		switch curRune {
		case '\'':
			if inputRune[l] == '\'' && l != i {
				l = i + 1
			} else if inputRune[l] != '"' && inputRune[l] != '{' {
				l = i
			}
		case '"':
			if inputRune[l] == '"' && l != i {
				l = i + 1
			} else if inputRune[l] != '\'' && inputRune[l] != '{' {
				l = i
			}
		case '{':
			if inputRune[l] != '\'' && inputRune[l] != '"' && i > 0 && inputRune[i-1] == '#' {
				l = i
			}
		case '}':
			if inputRune[l] == '{' && l > 0 && inputRune[l-1] == '#' {
				result += string(inputRune[nextEnd:l-1]) + replaceMap[string(inputRune[l+1:i])]
				nextEnd = i + 1
				l = i + 1
			}
		}
	}
	result += string(inputRune[nextEnd:inputLen])
	return result
}
