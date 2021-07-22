{{/*
Selector labels
*/}}
{{- define "mysql.selectorLabels" -}}
app: mysql
tier: mysql
{{- end }}

{{/*
Mysql labels
*/}}
{{- define "mysql.labels" -}}
{{ include "mysql.selectorLabels" . }}
{{ include "common.labels" . }}
{{- end }}