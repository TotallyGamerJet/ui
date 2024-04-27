package components

import (
	"bytes"
	"html/template"
)

var buttonTemplate = template.Must(template.ParseFS(files, "button.gohtml"))

type ButtonProps struct {
	Text     string
	Disabled bool
}

func BasicButton(text string, disabled bool) template.HTML {
	b := new(bytes.Buffer)
	if err := buttonTemplate.Execute(b, ButtonProps{Text: text, Disabled: disabled}); err != nil {
		return ErrorHtml("BasicButton", err)
	}
	return template.HTML(b.String())
}

func RoundedButton(text string) template.HTML {
	b := new(bytes.Buffer)
	if err := buttonTemplate.ExecuteTemplate(b, "RoundedButton", ButtonProps{Text: text}); err != nil {
		return ErrorHtml("BasicButton", err)
	}
	return template.HTML(b.String())
}
