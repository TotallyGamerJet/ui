package components

import (
	"html/template"
	"strings"
)

var errorHtml = template.Must(template.ParseFS(files, "error.gohtml"))

type ErrorHtmlProps struct {
	Name string
	Err  string
}

func ErrorHtml(name string, err error) template.HTML {
	b := new(strings.Builder)
	_ = errorHtml.Execute(b, ErrorHtmlProps{Name: name, Err: err.Error()})
	return template.HTML(b.String())
}
