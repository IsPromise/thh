package test

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
)

type Friend struct {
	Name string
}
type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func TestTmpl(_ *testing.T) {
	f1 := Friend{Name: "xiaofang"}
	f2 := Friend{Name: "wugui"}
	t := template.New("test")
	t = template.Must(t.Parse(
		`hello {{.UserName}}!
{{ range .Emails }}
an email {{ . }}
{{- end }}
{{ with .Friends }}
{{- range . }}
my friend name is {{.Name}}
{{- end }}
{{ end }}`))
	p := Person{UserName: "longshuai",
		Emails:  []string{"a1@qq.com", "a2@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	var b bytes.Buffer
	err := t.Execute(&b, p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b.String())
	}

}
