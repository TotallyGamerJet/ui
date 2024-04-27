package components

import (
	"html/template"
	"net/http"
)

var indexTemplate = template.Must(template.New("base.gohtml").
	Funcs(Components).
	ParseFS(files, "base.gohtml", "index.gohtml",
		/* ↓↓ This is only necessary if you want to call the template directly */
		"button.gohtml"))

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var viewData = map[string]any{
		"IsDisabled": func() bool {
			// Some logic here...
			return true
		}(),
		"MyViewData": map[string]any{
			"Text": "This button had data passed in",
		},
	}
	if err := indexTemplate.ExecuteTemplate(w, "base.gohtml", viewData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
