package config

import (
	"bytes"
	"html/template"
)

type help struct {
	TypesGrouped
	Detailed bool
}

const helpTemplate = `
{{$d := .Detailed}}{{range .TypesGrouped}}{{$g := .Name}}{{if $g}}group "{{.Name}}":
{{end}}
{{range .Types}}{{if $g}}	{{end}}{{.Name}} ({{.Type}}) [{{if .Aliases}}{{range $index, $element := .Aliases}}{{if $index}}, {{end}}{{ $element }}{{end}}{{end}}] - {{.Description}} {{if ne $d true}} - Default: "{{.Default}}"{{end}}
{{if $d}}{{.Documentation}}
		Default: "{{.Default}}"

{{end}}{{if $d}}
{{end}}{{end}}{{if ne $d true}}
{{end}}{{end}}`

func (c *Configs) GetHelp(name, group string, detailed bool) (text string,
	err error) {

	var buf []byte
	b := bytes.NewBuffer(buf)
	switch {
	case name == "" && group == "":
		g := help{c.GetGroups(), detailed}
		tmpl, _ := template.New("help").Parse(helpTemplate)
		err := tmpl.Execute(b, g)
		if log.E.Chk(err) {
			return "", err
		}
		text = b.String()
	case name == "" && group != "":
	case name != "" && group == "":
	}
	return
}
