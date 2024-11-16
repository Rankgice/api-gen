package main

import (
	"github.com/Rankgice/api-gen/cmd"
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
)

//var sqlUrl = flag.String("url", "", "The data source of database,like \"root:password@tcp(127.0.0.1:3306)/database")
//var table = flag.String("table", "", "The table or table globbing patterns in the database")
//var dir = flag.String("dir", "./desc/", "The target dir")
//var prefix = flag.String("prefix", "/prefix", "This is the API file routing prefix")
//var home = flag.String("home", "", "The apiGen home path of the template")

func main() {
	// 执行根命令
	cmd.Execute()
}

//func main1() {
//	// 解析命令行参数
//	flag.Parse()
//	if *sqlUrl == "" {
//		log.Fatalln("Error: --url flag is required")
//		return
//	} else if *table == "" {
//		log.Fatalln("Error: --table flag is required")
//		return
//	}
//	//自动补全斜杠
//	if (*dir)[len(*dir)-1] != '/' {
//		*dir += "/"
//	}
//
//	// 声明变量
//	var (
//		tableName      = utils.ToCamelCase(*table)
//		TableName      = utils.ToPascalCase(*table)
//		TableInfo      string
//		TableInfoProto string
//		tableComment   string
//	)
//	// 数据库连接信息
//	db, err := utils.ConnectMysql(*sqlUrl + "?charset=utf8mb4&parseTime=True&loc=Local")
//	defer db.Close()
//	if err != nil {
//		return
//	}
//	//解析数据库名
//	database := utils.ParseDatabaseName(*sqlUrl)
//	// 查询表字段信息
//	filedList, err := utils.QueryTableFiled(db, database, *table)
//	if err != nil {
//		return
//	}
//	//生成api结构体字段
//	for i, fieldInfo := range filedList {
//		TableInfo += utils.GenerateApiField(fieldInfo.Field, fieldInfo.DataType, fieldInfo.Nullable,
//			fieldInfo.Key, fieldInfo.DefaultVal, fieldInfo.Extra, fieldInfo.Comment)
//		TableInfoProto += utils.GenerateProtoField(fieldInfo.Field, fieldInfo.DataType, fieldInfo.Nullable,
//			fieldInfo.Key, fieldInfo.DefaultVal, fieldInfo.Extra, fieldInfo.Comment, i)
//
//	}
//	TableInfoProtoPage := fmt.Sprintf("\tint64\tpage\t = %d;// 页码\n\tint64\tpageSize\t = %d;// 每页数量", len(filedList)+1, len(filedList)+2)
//	// 查询表注释
//	tableComment, err = utils.QueryTableComment(db, database, *table)
//	if err != nil {
//		return
//	}
//	// 将 xxx表 => xxx, 去掉最后的'表'字
//	tableCommentRunes := []rune(tableComment)
//	if len(tableCommentRunes) == 0 {
//		tableComment = "业务"
//	} else if tableCommentRunes[len(tableCommentRunes)-1] == '表' {
//		tableComment = string(tableCommentRunes[:len(tableCommentRunes)-1])
//	}
//	//封装替换map
//	replaceMap := map[string]string{
//		"prefix":             *prefix,
//		"databaseName":       database,
//		"tableName":          tableName,
//		"TableName":          TableName,
//		"TableInfo":          TableInfo[:len(TableInfo)-1],
//		"TableInfoProto":     TableInfoProto,
//		"TableInfoProtoPage": TableInfoProtoPage,
//		"tableComment":       tableComment,
//	}
//
//	var data []byte
//	if *home == "" {
//		// 读取嵌入的模板文件
//		data, err = Asset("tpl/api.tpl")
//		if err != nil {
//			log.Fatalf("Error reading embedded template: %v", err)
//			return
//		}
//	} else {
//		// 读取用户指定的模板文件
//		data, err = os.ReadFile(*home)
//		if err != nil {
//			log.Fatalln("Error Reading file:", err)
//			return
//		}
//	}
//	//根据模板进行替换,生成API文件
//	result := utils.ReplaceTemplate(string(data), replaceMap)
//	fileName := fmt.Sprintf(*dir+"%s.api", tableName)
//	//检查上级目录是否存在，不存在则创建
//	err = os.MkdirAll(*dir, 0755)
//	if err != nil {
//		fmt.Println("Error creating directory:", err)
//		return
//	}
//	// 检查文件是否存在
//	if _, err := os.Stat(fileName); os.IsNotExist(err) {
//		// 文件不存在，创建并写入数据
//		err := os.WriteFile(fileName, []byte(result), 0644)
//		if err != nil {
//			log.Fatalln("Error writing file:", err)
//			return
//		}
//		log.Println("Generate API File successfully.")
//	} else if err != nil {
//		// 其他错误
//		log.Fatalln("Error checking file:", err)
//	} else {
//		// 文件已存在
//		log.Fatalln("API File already exists. ignored")
//	}
//}
