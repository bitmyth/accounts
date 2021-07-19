{{/*
Selector labels
*/}}
{{- define "web.selectorLabels" -}}
app: web
tier: frontend
{{- end }}

{{/*
Common labels
*/}}
{{- define "web.labels" -}}
helm.sh/chart: {{ include "api.chart" . }}
{{ include "web.selectorLabels" . }}
{{- include "common.labels" . }}
{{- end }}

