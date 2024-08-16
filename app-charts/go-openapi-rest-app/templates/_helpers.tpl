{{/*
Expand the name of the chart.
*/}}
{{- define "go-openapi-rest-app.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a fully qualified name.
*/}}
{{- define "go-openapi-rest-app.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := include "go-openapi-rest-app.name" . -}}
{{- if .Values.namespaceOverride -}}
{{- printf "%s-%s" .Values.namespaceOverride $name | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Get the name of the chart.
*/}}
{{- define "go-openapi-rest-app.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a name for the service.
*/}}
{{- define "go-openapi-rest-app.serviceName" -}}
{{- printf "%s-service" (include "go-openapi-rest-app.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a name for the deployment.
*/}}
{{- define "go-openapi-rest-app.deploymentName" -}}
{{- printf "%s-deployment" (include "go-openapi-rest-app.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- end -}}
