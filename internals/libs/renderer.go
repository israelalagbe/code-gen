package libs

import (
	"bytes"
	"text/template"
)

func RenderTemplate(path string, data interface{}) string {
	tmpl := ReadFile(path)

	t, err := template.New("test").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, data)
	if err != nil {
		panic(err)
	}

	return tpl.String()
}
