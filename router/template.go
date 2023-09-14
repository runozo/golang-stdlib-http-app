package router

import "html/template"

var defaultFuncs = template.FuncMap{
	"defTitle": func(ip interface{}) string {
		v, ok := ip.(string)
		if !ok || (ok && v == "") {
			return "Tasks - Golang Std HTTP"
		}
		return v
	},
}
var templateFiles = []string{
	"./web/templates/base.gohtml",
}

// tmplLayout appends the given files to the templateFiles slice.
//
// It takes a variadic parameter `files` of type `string` which represents
// the files to be appended to the `templateFiles` slice.
//
// It returns a slice of strings.
func tmplLayout(files ...string) []string {
	return append(templateFiles, files...)
}
