{{/*
Selector labels
*/}}
{{- define "web.selectorLabels" -}}
app: web
tier: frontend
{{- end }}

{{/*
Web labels
*/}}
{{- define "web.labels" -}}
{{ include "web.selectorLabels" . }}
{{ include "common.labels" . }}
{{- end }}

