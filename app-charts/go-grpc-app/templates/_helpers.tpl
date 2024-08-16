{{/*
Expand the name of the chart.
*/}}
{{- define "go-grpc-app.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a fully qualified name.
*/}}
{{- define "go-grpc-app.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := include "go-grpc-app.name" . -}}
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
{{- define "go-grpc-app.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a name for the service.
*/}}
{{- define "go-grpc-app.serviceName" -}}
{{- printf "%s-service" (include "go-grpc-app.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a name for the deployment.
*/}}
{{- define "go-grpc-app.deploymentName" -}}
{{- printf "%s-deployment" (include "go-grpc-app.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- end -}}
