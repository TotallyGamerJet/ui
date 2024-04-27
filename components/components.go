package components

import (
	"embed"
	"html/template"
)

//go:embed *.gohtml
var files embed.FS

// Components maps from component names to functions that return their HTML.
var Components template.FuncMap = template.FuncMap{
	"BasicButton":   BasicButton,
	"RoundedButton": RoundedButton,
}
