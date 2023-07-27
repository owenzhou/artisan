package model

var Template = `//代码生成器生成
package {{ .packageName }}
{{ $imports := "" -}}
{{- if eq .importTime true -}}
{{- $imports = (printf "%s%s\r\n\t" $imports "\"time\"") -}}
{{- end -}}
{{- if eq .importSql true -}}
{{- $imports = (printf "%s%s\r\n\t" $imports "\"database/sql\"") -}}
{{- end -}}
{{- if ne $imports "" -}}
import (
	{{ $imports }}
)
{{- end }}

type {{.modelName}} struct {
	{{- range $index, $value := .fields}}
	{{ $value.field }}  {{ $value.type }}  {{ $value.tag }}
	{{- end}}
}
`
