package utils

import (
	"database/sql"
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"log"
	"os"
	"strings"
)

// TableFieldInfo 表字段信息
type TableFieldInfo struct {
	Field      string         `json:"field"`
	DataType   string         `json:"dataType"`
	Nullable   string         `json:"nullable"`
	Key        string         `json:"key"`
	DefaultVal sql.NullString `json:"defaultVal"`
	Extra      string         `json:"extra"`
	Comment    string         `json:"comment"`
}

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

// ParseDatabaseName 解析数据库名称
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

// MysqlTypeToApiType 将MySQL数据类型转换为Api数据类型
func MysqlTypeToApiType(mysqlType string) string {
	switch mysqlType {
	case "int", "tinyint", "smallint", "mediumint", "bigint":
		return "int64"
	case "varchar", "char", "json", "text", "tinytext", "mediumtext", "longtext":
		return "string"
	case "float", "double", "decimal":
		return "float64"
	case "date", "datetime", "timestamp":
		return "string"
	case "boolean":
		return "bool"
	default:
		return "interface{}"
	}
}

// MysqlTypeToProtoType 将MySQL数据类型转换为Proto数据类型
func MysqlTypeToProtoType(mysqlType string) string {
	switch mysqlType {
	case "int", "tinyint", "smallint", "mediumint", "bigint":
		return "int64"
	case "varchar", "char", "json", "text", "tinytext", "mediumtext", "longtext":
		return "string"
	case "float", "double", "decimal":
		return "double"
	case "date", "datetime", "timestamp":
		return "string"
	case "boolean":
		return "bool"
	default:
		return "string"
	}
}

// GenerateProtoField 生成api结构体字段
func GenerateProtoField(field, dataType, nullable, key string, val sql.NullString, extra, comment string, index int) string {
	FieldName := ToPascalCase(field)
	fieldType := MysqlTypeToProtoType(dataType)
	return fmt.Sprintf("\t%s\t%s\t = %d;// %s\n", fieldType, FieldName, index+1, comment)
}

// GenerateApiField 生成api结构体字段
func GenerateApiField(field, dataType, nullable, key string, val sql.NullString, extra, comment string) string {
	FieldName := ToPascalCase(field)
	fieldType := MysqlTypeToApiType(dataType)
	return fmt.Sprintf("\t%s\t%s\t`json:\"%s,optional\" form:\"%s,optional\"` // %s\n", FieldName, fieldType, field, field, comment)
}

// ConnectMysql 连接mysql
func ConnectMysql(dsn string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("Error connect to mysql", err)
		return
	}
	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatalln("Error connect to database", err)
		return
	}
	return
}

// QueryTableFiled 查询表字段
func QueryTableFiled(db *sql.DB, databaseName, tableName string) (tableFieldInfoList []TableFieldInfo, err error) {
	tableFieldInfoList = make([]TableFieldInfo, 0)
	// 查询表字段
	rows, err := db.Query(`
		SELECT 
			COLUMN_NAME AS "field",
			DATA_TYPE AS "dataType",
			IS_NULLABLE AS "nullable",   
			COLUMN_KEY AS "key",        
			COLUMN_DEFAULT AS "defaultVal", 
			EXTRA AS "extra",               
			COLUMN_COMMENT AS "comment"  
		FROM 
			INFORMATION_SCHEMA.COLUMNS
		WHERE 
			TABLE_SCHEMA = ? 
			AND TABLE_NAME = ?
			ORDER BY ORDINAL_POSITION
		`, databaseName, tableName)
	if err != nil {
		log.Fatalln("Error query data:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var tableFieldInfo TableFieldInfo
		err = rows.Scan(&(tableFieldInfo.Field), &(tableFieldInfo.DataType),
			&(tableFieldInfo.Nullable), &(tableFieldInfo.Key), &(tableFieldInfo.DefaultVal),
			&(tableFieldInfo.Extra), &(tableFieldInfo.Comment))
		if err != nil {
			log.Fatalln("Error resolving database data:", err)
			return
		}
		tableFieldInfoList = append(tableFieldInfoList, tableFieldInfo)
	}
	return
}

// QueryTableComment 查询表注释
func QueryTableComment(db *sql.DB, databaseName, tableName string) (tableComment string, err error) {
	err = db.QueryRow(`
		SELECT TABLE_COMMENT 
		FROM INFORMATION_SCHEMA.TABLES
		WHERE 
			TABLE_SCHEMA = ? 
			AND TABLE_NAME = ?
		`, databaseName, tableName).Scan(&tableComment)
	if err != nil {
		log.Fatalln("Error querying table comment: ", err)
		return
	}
	return
}

func GenFile(dir, fileName, data string) {
	//检查上级目录是否存在，不存在则创建
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	absFileName := dir + "/" + fileName
	// 检查文件是否存在
	if _, err := os.Stat(absFileName); os.IsNotExist(err) {
		// 文件不存在，创建并写入数据
		err := os.WriteFile(absFileName, []byte(data), 0644)
		if err != nil {
			log.Fatalln("Error writing file:", err)
			return
		}
		log.Println("Generate API File successfully.")
	} else if err != nil {
		// 其他错误
		log.Fatalln("Error checking file:", err)
	} else {
		// 文件已存在
		log.Println("API File already exists. ignored")
	}
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
		case '{':
			if i > 0 && inputRune[i-1] == '#' {
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
