package api

import (
	"database/sql"
	"fmt"
	"github.com/Rankgice/api-gen/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// RunApi api命令
func RunApi(cmd *cobra.Command, args []string) {
	url, _ := cmd.Flags().GetString("url")
	table, _ := cmd.Flags().GetString("table")
	dir, _ := cmd.Flags().GetString("dir")
	prefix, _ := cmd.Flags().GetString("prefix")
	home, _ := cmd.Flags().GetString("home")
	service, _ := cmd.Flags().GetString("service")

	//解析数据库名
	database := utils.ParseDatabaseName(url)
	if service == "" {
		service = utils.ToCamelCase(database)
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
		tableName    = utils.ToCamelCase(table)
		TableName    = utils.ToPascalCase(table)
		TableInfo    string
		tableComment string
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
	//生成api结构体字段
	for _, fieldInfo := range filedList {
		TableInfo += utils.GenerateApiField(fieldInfo.Field, fieldInfo.DataType, fieldInfo.Nullable,
			fieldInfo.Key, fieldInfo.DefaultVal, fieldInfo.Extra, fieldInfo.Comment)
	}
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
		"prefix":       prefix,
		"serviceName":  service,
		"tableName":    tableName,
		"TableName":    TableName,
		"TableInfo":    TableInfo[:len(TableInfo)-1],
		"tableComment": tableComment,
	}

	var data []byte
	var commonData []byte
	var serviceData []byte
	if home == "" {
		// 读取嵌入的模板文件
		data, err = utils.Asset("tpl/api.tpl")
		if err != nil {
			log.Fatalf("Error reading embedded template: %v", err)
			return
		}
		commonData, err = utils.Asset("tpl/api_common.tpl")
		if err != nil {
			log.Fatalf("Error reading embedded template: %v", err)
			return
		}
		serviceData, err = utils.Asset("tpl/api_service.tpl")
		if err != nil {
			log.Fatalf("Error reading embedded template: %v", err)
		}
	} else {
		// 读取用户指定的模板文件
		data, err = os.ReadFile(home + "/api.tpl")
		if err != nil {
			log.Fatalln("Error Reading file:", err)
			return
		}
		commonData, err = os.ReadFile(home + "/api_common.tpl")
		if err != nil {
			log.Fatalln("Error Reading file:", err)
			return
		}
		serviceData, err = os.ReadFile(home + "/api_service.tpl")
		if err != nil {
			log.Fatalln("Error Reading file:", err)
		}
	}
	//根据模板进行替换,生成API文件
	result := utils.ReplaceTemplate(string(data), replaceMap)
	fileDir := dir + "/" + tableName
	fileName := fmt.Sprintf("%s.api", tableName)
	utils.GenApiFile(fileDir, fileName, result)
	//生成common.api
	utils.GenApiFile(dir, "common.api", string(commonData))
	//生成#{service}.api
	serviceResult := utils.ReplaceTemplate(string(serviceData), replaceMap)
	utils.GenApiFile(dir, fmt.Sprintf("%s.api", service), serviceResult)
}
