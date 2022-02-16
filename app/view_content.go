package app

var ViewContentTemplate = `
{{ $a := "{{ define \"content\" }}" }}
{{- $a }}
<div>内容页</div>
{{ $b := "{{ end }}" }}
{{- $b }}
`