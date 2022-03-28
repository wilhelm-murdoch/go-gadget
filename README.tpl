# Packages
{{ range .packages }}
  * `{{ . }}`{{ end }}

# Methods
{{ range .functions }}{{ if .Exported }}* [{{ .Name }}](#{{.Name}}){{ end }}
{{ end }}

{{ range .functions }}{{ if and .Exported .Example }}## {{ .Name }}
{{ .Comment }}
{{ .Body }}
```go
{{ .Example }}
```
{{ end }}{{ end }}