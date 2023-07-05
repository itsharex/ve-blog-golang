package easycode

import (
	"fmt"
	"github.com/ve-weiyi/ve-admin-store/server/global"
	"github.com/ve-weiyi/ve-admin-store/server/infra/testinit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"path"
	"testing"
)

// GEN 自动生成 GORM 模型结构体文件及使用示例 https://blog.csdn.net/Jeffid/article/details/126898000
const dsn = "root:mysql7914@(127.0.0.1:3306)/blog-v2?charset=utf8mb4&parseTime=True&loc=Local"

var db *gorm.DB

func init() {
	testinit.Init()
	log.SetFlags(log.LstdFlags | log.Llongfile)
	var err error
	// 连接数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "tb_",
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	log.Println("mysql connection done")
}

func TestPlate(t *testing.T) {
	out := path.Join(global.GetRuntimeRoot(), "server/api", "blog")
	//out := path.Join("./autocode_template", "test")

	cfg := Config{
		db:             nil,
		Replace:        true,
		ReplaceCommon:  true,
		GenerateCommon: true,
		OutPath:        out,
		OutFileNS: func(tableName string) (fileName string) {
			return fmt.Sprintf("tb_%v", tableName)
		},
		IgnoreMap: map[string]string{
			//tmpl.KeyRouter:     "",
			//tmpl.KeyController: "",
			//tmpl.KeyService:    "",
			//tmpl.KeyRepository: "",
			//tmpl.KeyModel:      "",
		},
	}
	typeInt := "int"
	// 自定义字段的数据类型
	// 统一数字类型为int64,兼容protobuf
	dataMap := map[string]func(columnType gorm.ColumnType) (dataType string){
		//"tinyint":    func(columnType gorm.ColumnType) (dataType string) { return typeInt },
		"smallint":  func(columnType gorm.ColumnType) (dataType string) { return typeInt },
		"mediumint": func(columnType gorm.ColumnType) (dataType string) { return typeInt },
		"bigint":    func(columnType gorm.ColumnType) (dataType string) { return typeInt },
		"int":       func(columnType gorm.ColumnType) (dataType string) { return typeInt },
		//"datetime":  func(columnType gorm.ColumnType) (dataType string) { return "*time.Time" },
	}
	cfg.WithDataTypeMap(dataMap)
	cfg.WithJSONTagNameStrategy(func(columnName string) (tagContent string) {
		//toStringField := "time"
		//if strings.Contains(columnName, toStringField) {
		// return columnName + "\" example:\"2022-11-16T16:00:00.000Z"
		//}
		return columnName
	})

	gen := NewGenerator(cfg)
	gen.UseDB(db)
	//gen.InitPackage("hello")
	//gen.ApplyMetas(gen.GenerateMetasFromSchema())
	//gen.ApplyMetas(gen.GenerateMetasFromTable("comment", "评论"))
	//gen.ApplyMetas(gen.GenerateMetasFromTable("friend_link", "友链"))
	//gen.ApplyMetas(gen.GenerateMetasFromTable("photo", "相片"))
	//gen.ApplyMetas(gen.GenerateMetasFromTable("photo_album", "相册"))
	//gen.ApplyMetas(gen.GenerateMetasFromTable("talk", "说说"))
	//gen.ApplyMetas(gen.GenerateMetasFromTable("user_account", "用户账号信息"))
	//gen.ApplyMetas(gen.GenerateMetasFromTable("user_information", "用户信息"))
	//gen.ApplyMetas(gen.GenerateMetasFromTable("user_login_history", "用户登录历史"))
	//gen.ApplyMetas(gen.GenerateMetasFromTable("role", "角色"))
	//gen.ApplyMetas(gen.GenerateMetasFromTable("menu", "菜单"))
	gen.ApplyMetas(gen.GenerateMetasFromTable("upload", "文件上传"))
	//gen.GenerateCommonFile("upload", "文件上传")

	//gen.RollBack()
	gen.Execute()
}
