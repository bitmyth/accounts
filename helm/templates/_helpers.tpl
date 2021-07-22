{{/*
Selector labels
*/}}
{{- define "api.selectorLabels" -}}
app: api
tier: backend
{{- end }}

{{/*
labels
*/}}
{{- define "api.labels" -}}
{{ include "api.selectorLabels" . }}
{{ include "common.labels" . }}
{{- end }}
{{/*

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "api.chart" -}}
{{ printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

Common labels
*/}}
{{- define "common.labels" -}}
helm.sh/chart: {{ include "api.chart" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}
