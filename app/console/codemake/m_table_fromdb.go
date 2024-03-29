package codemake

import (
	"fmt"
	"github.com/leancodebox/goose/preferences"
	"regexp"
	"thh/app/bundles/eh"
	"thh/app/bundles/logging"

	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "migrate:table",
		Short: "导表助手",
		Long:  ".env 文件中 TARGET_DATABASE_URL(local) 为导入db地址， ORIGIN_DATABASE_URL(tmp) 为被导入db地址。会获取tmp中的表信息导入至local中",
		Run:   runMTableFromDb,
		//Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

func runMTableFromDb(_ *cobra.Command, _ []string) {
	// init
	dataSourceName := preferences.GetString("dbtool.originUrl")
	localSourceName := preferences.GetString("dbtool.targetUrl")
	localDb, err := gorm.Open(mysql.Open(localSourceName), &gorm.Config{PrepareStmt: false,
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 全局禁用表名复数
		Logger:         logger.Default})
	if logging.ErrIf(err) {
		return
	}
	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{PrepareStmt: false,
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 全局禁用表名复数
		Logger:         logger.Default})
	if logging.ErrIf(err) {
		return
	}

	rows, err := db.Raw("show tables").Rows()
	tbDesc := make(map[string]string)
	if logging.ErrIf(err) {
		return
	}
	for rows.Next() {
		var table string
		logging.ErrIf(rows.Scan(&table))
		tbDesc[table] = table
	}
	eh.PrIF(rows.Close())

	reg, _ := regexp.Compile("AUTO_INCREMENT=[0-9]+")
	//根据规则提取关键信息
	for tmpTableName, _ := range tbDesc {
		var list []ShowTable

		// Get table annotations.获取表注释
		sctSql := "show create table " + tmpTableName
		db.Raw(sctSql).Scan(&list)
		if len(list) == 1 {
			//list[0].CreateTable
			createSql := reg.ReplaceAllString(list[0].CreateTable, `AUTO_INCREMENT=1`)
			cErr := localDb.Exec(createSql).Error
			if cErr != nil {
				fmt.Println()
			}
		} else {
			fmt.Println("存在异常查询", sctSql)
		}
	}

	fmt.Println("build end")

}

type ShowTable struct {
	Table       string `gorm:"column:Table"`
	CreateTable string `gorm:"column:Create Table"`
}
