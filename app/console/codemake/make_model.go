package codemake

import (
	"fmt"

	"github.com/leancodebox/goose/fileopt"
	"github.com/leancodebox/goose/preferences"
	"github.com/leancodebox/goose/stropt"
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "make:modelFromStruct",
		Short: "modelFromStruct",
		Run:   makeModel,
		//Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

func makeModel(_ *cobra.Command, _ []string) {

	//var modelName string
	//
	//var fieldName string
	//var typeName string
	//var cList []genColumns
	//
	//cList = append(cList, genColumns{Field: "id", Type: "bigint unsigned", Key: "PRI", Desc: "", Null: "NO", Default: GetStringP("0")})
	//
	//fmt.Println("请输入要创建的 model 名")
	//_, _ = fmt.Scanln(&modelName)
	//fmt.Println(modelName)
	//
	//for {
	//	fmt.Println("输入字段")
	//	_, _ = fmt.Scanln(&fieldName)
	//	if strings.TrimSpace(fieldName) == "" {
	//		fmt.Println("未家检测到新的字段")
	//		break
	//	}
	//	fieldName = str.Snake(fieldName)
	//	fmt.Println("选择类型")
	//	_, _ = fmt.Scanln(&typeName)
	//	cList = append(cList, genColumns{
	//		Field: fieldName,
	//		Type:  typeName,
	//		Key:   "",
	//		Desc:  "",
	//		Null:  "NO",
	//	})
	//}
	//
	//fmt.Println(cList)

	// 抓取历史
	//outPutModel("FTwitterSpiderHis", []genColumns{
	//	{Field: "id", Type: "bigint unsigned", Key: "PRI", Desc: "主键", Null: "NO", Default: nil},
	//	{Field: "params", Type: "text", Key: "", Desc: "请求参数", Null: "NO", Default: nil},
	//	{Field: "context", Type: "text", Key: "", Desc: "", Null: "", Default: nil},
	//	{Field: "type", Type: "varchar(255)", Key: "", Desc: "", Null: "NO", Default: GetStringP("follow")},
	//	{Field: "create_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	//})
	//
	//outPutModel("FTwitterUser", []genColumns{
	//	{Field: "id", Type: "bigint unsigned", Key: "PRI", Desc: "主键", Null: "NO", Default: nil},
	//	{Field: "rest_id", Type: "varchar(255)", Key: "MUL", Desc: "用户id", Null: "NO", Default: GetStringP("")},
	//	{Field: "screen_name", Type: "varchar(255)", Key: "MUL", Desc: "用户id", Null: "NO", Default: GetStringP("")},
	//	{Field: "desc", Type: "text", Key: "", Desc: "", Null: "", Default: nil},
	//	{Field: "create_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	//})
	//
	//// 用户历史记录表，记录用户历史上的数据，每次抓取都会记录
	//outPutModel("FTwitterUserHis", []genColumns{
	//	{Field: "id", Type: "bigint unsigned", Key: "PRI", Desc: "主键", Null: "NO", Default: nil},
	//	{Field: "rest_id", Type: "varchar(255)", Key: "MUL", Desc: "用户id", Null: "NO", Default: GetStringP("")},
	//	{Field: "screen_name", Type: "varchar(255)", Key: "MUL", Desc: "用户id", Null: "NO", Default: GetStringP("")},
	//	{Field: "desc", Type: "text", Key: "", Desc: "", Null: "", Default: GetStringP("")},
	//	{Field: "create_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	//})
	//
	//outPutModel("FTwitterTweet", []genColumns{
	//	{Field: "id", Type: "bigint unsigned", Key: "PRI", Desc: "主键", Null: "NO", Default: nil},
	//	{Field: "screen_name", Type: "varchar(255)", Key: "MUL", Desc: "", Null: "NO", Default: nil},
	//	{Field: "conversation_id", Type: "varchar(255)", Key: "MUL", Desc: "", Null: "NO", Default: nil},
	//	{Field: "context", Type: "text", Key: "", Desc: "", Null: "", Default: nil},
	//	{Field: "create_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	//})

	//outPutModel("PhoneLocation", []genColumns{
	//	{Field: "id", Type: "bigint unsigned", Key: "PRI", Desc: "主键", Null: "NO", Default: nil},
	//	{Field: "pref", Type: "varchar(255)", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "phone", Type: "varchar(255)", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "province", Type: "varchar(255)", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "city", Type: "varchar(255)", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "isp", Type: "varchar(255)", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "post_code", Type: "varchar(255)", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "city_code", Type: "varchar(255)", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "area_code", Type: "varchar(255)", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//})

	//outPutModel("Articles", []genColumns{
	//	{Field: "id", Type: "bigint unsigned", Key: "PRI", Desc: "主键", Null: "NO", Default: nil},
	//	{Field: "content", Type: "text", Key: "", Desc: "", Null: "", Default: GetStringP("")},
	//	{Field: "user_id", Type: "bigint", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "create_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	//	{Field: "update_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	//})
	//
	//outPutModel("Comment", []genColumns{
	//	{Field: "id", Type: "bigint unsigned", Key: "PRI", Desc: "主键", Null: "NO", Default: nil},
	//	{Field: "article_id", Type: "bigint", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "content", Type: "text", Key: "", Desc: "", Null: "", Default: GetStringP("")},
	//	{Field: "user_id", Type: "bigint", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "create_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	//	{Field: "update_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	//})
	//
	//outPutModel("Post", []genColumns{
	//	{Field: "id", Type: "bigint unsigned", Key: "PRI", Desc: "主键", Null: "NO", Default: nil},
	//	{Field: "title", Type: "bigint", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "slug", Type: "text", Key: "", Desc: "", Null: "", Default: GetStringP("")},
	//	{Field: "summary", Type: "bigint", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "content", Type: "bigint", Key: "", Desc: "", Null: "NO", Default: GetStringP("0")},
	//	{Field: "create_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	//	{Field: "update_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	//})
	//outPutModel("FTwitterFilterUser", []genColumns{
	//	{Field: "id", Type: "bigint unsigned", Key: "PRI", Desc: "主键", Null: "NO", Default: nil},
	//	{Field: "screen_name", Type: "varchar(255)", Key: "MUL", Desc: "", Null: "NO", Default: nil},
	//	{Field: "create_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	//	{Field: "deleted_at", Type: "datetime", Key: "", Desc: "", Null: "", Default: nil},
	//})
	//outPutModel("TodoTask", []genColumns{
	//	{Field: "task_id", Type: "bigint unsigned", Key: "PRI", Desc: "主键", Null: "NO", Default: nil},
	//	{Field: "task_name", Type: "varchar(50)", Key: "", Desc: "任务名", Null: "NO", Default: nil},
	//	{Field: "task_description", Type: "varchar(1024)", Key: "任务描述", Desc: "", Null: "NO", Default: nil},
	//	{Field: "status", Type: "int", Key: "", Desc: "任务状态（0：未完成，1：已完成）", Null: "NO", Default: nil},
	//	{Field: "create_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	//	{Field: "deadline", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: nil},
	//	{Field: "weight", Type: "int", Key: "", Desc: "任务权重（用于优先级排序）", Null: "NO", Default: nil},
	//	{Field: "paused", Type: "int", Key: "", Desc: "任务暂停状态（0：未暂停，1：已暂停）", Null: "NO", Default: nil},
	//})

	outPutModel("ToolWecomGroupInfo", []genColumns{
		{Field: "id", Type: "bigint unsigned", Key: "PRI", Desc: "主键", Null: "NO", Default: nil},
		{Field: "corp_id", Type: "varchar(255)", Key: "corpId", Desc: "", Null: "NO", Default: nil},
		{Field: "name", Type: "varchar(1024)", Key: "", Desc: "name", Null: "NO", Default: nil},
		{Field: "chat_id", Type: "varchar(255)", Key: "chatId", Desc: "", Null: "NO", Default: nil},
		{Field: "status", Type: "int", Key: "", Desc: "0,1", Null: "NO", Default: nil},
		{Field: "create_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
		{Field: "update_time", Type: "datetime", Key: "MUL", Desc: "", Null: "NO", Default: GetStringP("CURRENT_TIMESTAMP")},
	})
}

func outPutModel(modelName string, list []genColumns) {
	outputRoot := preferences.GetString("dbtool.output", "./storage/model/")
	outputRoot = `./storage/model/`
	modelPath := modelName
	modelEntityPath := outputRoot + modelPath + "/" + modelPath + ".go"
	connectPath := outputRoot + modelPath + "/" + modelPath + "_connect.go"
	repPath := outputRoot + modelPath + "/" + modelPath + "_rep.go"

	modelStr, connectStr, repStr := buildModelContent(stropt.Snake(modelName), list)

	fmt.Println(modelStr, connectStr, repStr)
	fmt.Println(modelEntityPath)
	fmt.Println(connectPath)
	fmt.Println(repPath)
	fileopt.PutContent(modelEntityPath, modelStr)
	fileopt.PutContent(connectPath, connectStr)
	fileopt.IsExistOrCreate(repPath, repStr)
}

func GetStringP(value string) *string {
	return &value
}
