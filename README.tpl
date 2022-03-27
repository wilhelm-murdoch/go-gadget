# Packages
{{ range . }}
* {{ .Name }} {{ end }}

# Methods
{{ range . }} {{ range .Files }} {{ range .Functions }}
* [{{ .Name }}](#{{ .Name }}) {{ end }} {{ end }} {{ end }}

{{ range . }} {{ range .Files }} {{ range .Functions }}
## {{ .Name }}
{{ .Comment }} {{ end }} {{ end }} {{ end }}