package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/Rankgice/api-gen/utils"
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
	"log"
	"os"
)

var sqlUrl = flag.String("url", "", "The data source of database,like \"root:password@tcp(127.0.0.1:3306)/database")
var table = flag.String("table", "", "The table or table globbing patterns in the database")
var dir = flag.String("dir", "./desc/", "The target dir")
var prefix = flag.String("prefix", "/prefix", "This is the API file routing prefix")
var home = flag.String("home", "", "The apiGen home path of the template")

func main() {
	// 解析命令行参数
	flag.Parse()
	if *sqlUrl == "" {
		log.Fatalln("Error: --url flag is required")
	} else if *table == "" {
		log.Fatalln("Error: --table flag is required")
	}
	if (*dir)[len(*dir)-1] != '/' {
		*dir += "/"
	}
	// 数据库连接信息
	dsn := *sqlUrl + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("Error connect to mysql", err)
	}
	defer db.Close()
	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatalln("Error connect to database", err)
	}
	database := utils.ParseDatabaseName(*sqlUrl)
	// 查询数据
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
	`, database, *table)
	if err != nil {
		log.Fatalln("Error query data:", err)
		return
	}
	defer rows.Close()

	tableName := utils.ToCamelCase(*table)
	TableName := utils.ToPascalCase(*table)
	var TableInfo string

	for rows.Next() {
		var field, dataType, nullable, key, extra, comment string
		var defaultVal sql.NullString
		err := rows.Scan(&field, &dataType, &nullable, &key, &defaultVal, &extra, &comment)
		if err != nil {
			log.Fatalln("Error resolving database data:", err)
		}
		TableInfo += utils.GenerateField(field, dataType, nullable, key, defaultVal, extra, comment)
	}

	replaceMap := map[string]string{
		"prefix":    *prefix,
		"tableName": tableName,
		"TableName": TableName,
		"TableInfo": TableInfo[:len(TableInfo)-1],
	}
	var data []byte
	if *home == "" {
		// 读取嵌入的模板文件
		data, err = Asset("tpl/api.tpl")
		if err != nil {
			log.Fatalf("Error reading embedded template: %v", err)
		}
	} else {
		// 读取用户指定的模板文件
		data, err = os.ReadFile(*home)
		if err != nil {
			log.Fatalln("Error Reading file:", err)
			return
		}
	}
	//根据模板进行替换,生成API文件
	result := utils.ReplaceTemplate(string(data), replaceMap)
	fileName := fmt.Sprintf(*dir+"%s.api", tableName)
	//检查上级目录是否存在，不存在则创建
	err = os.MkdirAll(*dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	// 检查文件是否存在
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// 文件不存在，创建并写入数据
		err := os.WriteFile(fileName, []byte(result), 0644)
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
		log.Fatalln("API File already exists. ignored")
	}
}
