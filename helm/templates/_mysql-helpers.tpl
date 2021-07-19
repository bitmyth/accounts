{{/*
Selector labels
*/}}
{{- define "mysql.selectorLabels" -}}
app: mysql
tier: mysql
{{- end }}

{{/*
Common labels
*/}}
{{- define "mysql.labels" -}}
helm.sh/chart: {{ include "api.chart" . }}
{{ include "mysql.selectorLabels" . }}
{{- include "common.labels" . }}
{{- end }}