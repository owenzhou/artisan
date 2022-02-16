package app

var ViewContentTemplate = `
{{ $a := "{{ define \"content\" }}" }}
{{- $a }}
<div>{{ $c := "{{ .title }}" }}{{ $c }}</div>
{{ $b := "{{ end }}" }}
{{- $b }}
`