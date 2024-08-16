{{/*
Expand the name of the chart.
*/}}
{{- define "nodejs-grpc-app.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a fully qualified name.
*/}}
{{- define "nodejs-grpc-app.fullname" -}}
{{- if .Values.fullnameOverride -}}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" -}}
{{- else -}}
{{- $name := include "nodejs-grpc-app.name" . -}}
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
{{- define "nodejs-grpc-app.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a name for the service.
*/}}
{{- define "nodejs-grpc-app.serviceName" -}}
{{- printf "%s-service" (include "nodejs-grpc-app.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{/*
Create a name for the deployment.
*/}}
{{- define "nodejs-grpc-app.deploymentName" -}}
{{- printf "%s-deployment" (include "nodejs-grpc-app.fullname" .) | trunc 63 | trimSuffix "-" -}}
{{- end -}}
