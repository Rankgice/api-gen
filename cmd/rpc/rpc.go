package rpc

import (
	"database/sql"
	"fmt"
	"github.com/Rankgice/api-gen/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// RunRpc rpc命令
func RunRpc(cmd *cobra.Command, args []string) {
	url, _ := cmd.Flags().GetString("url")
	table, _ := cmd.Flags().GetString("table")
	dir, _ := cmd.Flags().GetString("dir")
	packageName, _ := cmd.Flags().GetString("package")
	home, _ := cmd.Flags().GetString("home")

	//解析数据库名
	database := utils.ParseDatabaseName(url)
	if packageName == "" {
		packageName = utils.ToCamelCase(database)
	}
	//自动去除斜杠
	if dir[len(dir)-1] == '/' {
		dir = dir[:len(dir)-1]
	}
	if len(home) > 0 && home[len(home)-1] == '/' {
		home = home[:len(home)-1]
	}

	// 声明变量
	var (
		tableName          = utils.ToCamelCase(table)
		TableName          = utils.ToPascalCase(table)
		TableInfo          string
		TableInfoProtoPage string
		tableComment       string
	)
	// 数据库连接信息
	db, err := utils.ConnectMysql(url + "?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	// 查询表字段信息
	filedList, err := utils.QueryTableFiled(db, database, table)
	if err != nil {
		return
	}
	//生成rpc结构体字段
	for i, fieldInfo := range filedList {
		TableInfo += utils.GenerateProtoField(fieldInfo.Field, fieldInfo.DataType, fieldInfo.Nullable,
			fieldInfo.Key, fieldInfo.DefaultVal, fieldInfo.Extra, fieldInfo.Comment, i)
	}
	//生成rpc分页字段信息
	TableInfoProtoPage += fmt.Sprintf("\tint64 page = %d; // 页码\n", len(filedList)+1)
	TableInfoProtoPage += fmt.Sprintf("\tint64 pageSize = %d; // 每页数量", len(filedList)+2)

	// 查询表注释
	tableComment, err = utils.QueryTableComment(db, database, table)
	if err != nil {
		return
	}
	// 将 xxx表 => xxx, 去掉最后的'表'字
	tableCommentRunes := []rune(tableComment)
	if len(tableCommentRunes) == 0 {
		tableComment = "业务"
	} else if tableCommentRunes[len(tableCommentRunes)-1] == '表' {
		tableComment = string(tableCommentRunes[:len(tableCommentRunes)-1])
	}

	//封装替换map
	replaceMap := map[string]string{
		"packageName":        packageName,
		"tableName":          tableName,
		"TableName":          TableName,
		"TableInfo":          TableInfo[:len(TableInfo)-1],
		"TableInfoProtoPage": TableInfoProtoPage,
		"tableComment":       tableComment,
	}

	var data []byte
	if home == "" {
		// 读取嵌入的模板文件
		data, err = utils.Asset("tpl/proto.tpl")
		if err != nil {
			log.Fatalf("Error reading embedded template: %v", err)
			return
		}
	} else {
		// 读取用户指定的模板文件
		data, err = os.ReadFile(home + "/proto.tpl")
		if err != nil {
			log.Fatalln("Error Reading file:", err)
			return
		}
	}
	//根据模板进行替换,生成API文件
	result := utils.ReplaceTemplate(string(data), replaceMap)
	utils.GenFile(dir, fmt.Sprintf("%s.proto", tableName), result)
}
