// Package gen 命令行的 gen 命令
package codemake

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/leancodebox/goose/fileopt"
	"github.com/leancodebox/goose/stropt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"text/template"
	"thh/arms/output"
)

type Model struct {
	StructName  string
	PackageName string
	ClientName  string
}

//go:embed tmpl
var tmplFS embed.FS

// CmdMake 说明 cobra 命令
var CmdMake = &cobra.Command{
	Use:   "make",
	Short: "Generate file and code",
}

func init() {
	// 注册 gen 的子命令
	CmdMake.AddCommand(
		CmdMakeCMD,
		CmdMakeRemoteService,
	)
	appendCommand(CmdMake)
}

// makeModelFromString 格式化用户输入的内容
func makeModelFromString(name string) Model {
	model := Model{}
	model.StructName = stropt.Singular(strcase.ToCamel(name))
	model.PackageName = stropt.Snake(model.StructName)
	model.ClientName = stropt.Camel(model.StructName)
	return model
}

func buildWithOutput(data map[string]any, filePath string, tmplPath string) {
	outputData := buildByTmpl(data, tmplPath)
	dirPath := filepath.Dir(filePath)
	if !fileopt.IsExist(filePath) {
		if err := os.MkdirAll(dirPath, 0666); err != nil {
		}
	}
	err := fileopt.Put([]byte(outputData), filePath)
	if err != nil {
		output.Exit(err.Error())
	}
	// 提示成功
	output.Success(fmt.Sprintf("[%s] created.", filePath))
}

func buildByTmpl(data map[string]any, tmplPath string) string {

	modelData, err := tmplFS.ReadFile(tmplPath)
	if err != nil {
		fmt.Println("err", err)
	}
	modelStub := string(modelData)

	var b bytes.Buffer
	t := template.New("test")
	t.Funcs(template.FuncMap{
		"AddOne": func(p int) int { return p + 1 },
	})
	t = template.Must(t.Parse(modelStub))

	err = t.Execute(&b, data)

	if err != nil {
		fmt.Println(err)
	}
	return b.String()
}
