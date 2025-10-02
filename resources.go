// Copyright 2015-2025 Bleemeo
//
// bleemeo.com an infrastructure monitoring solution in the Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bleemeo

// A Resource represents a route to a model on the Bleemeo API.
type Resource = string

// Available resources on the Bleemeo API.
const (
	ResourceAccount               Resource = "v1/account/"
	ResourceAccountConfig         Resource = "v1/accountconfig/"
	ResourceAgent                 Resource = "v1/agent/"
	ResourceAgentConfig           Resource = "v1/agentconfig/"
	ResourceAgentFact             Resource = "v1/agentfact/"
	ResourceAgentType             Resource = "v1/agenttype/"
	ResourceApplication           Resource = "v1/application/"
	ResourceAuditLog              Resource = "v1/auditlog/"
	ResourceAWSIntegration        Resource = "v1/awsintegration/"
	ResourceContactsGroup         Resource = "v1/contactsgroup/"
	ResourceContainer             Resource = "v1/container/"
	ResourceDashboard             Resource = "v1/dashboard/"
	ResourceEvent                 Resource = "v1/event/"
	ResourceFlappyConfiguration   Resource = "v1/flappyconfiguration/"
	ResourceForecast              Resource = "v1/forecast/"
	ResourceForecastConfiguration Resource = "v1/forecastconfiguration/"
	ResourceGloutonConfigItem     Resource = "v1/gloutonconfigitem/"
	ResourceGloutonCrashReport    Resource = "v1/gloutoncrashreport/"
	ResourceGloutonDiagnostic     Resource = "v1/gloutondiagnostic/"
	ResourceHealthCheck           Resource = "v1/healthcheck/"
	ResourceIntegration           Resource = "v1/integration/"
	ResourceIntegrationTemplate   Resource = "v1/integrationtemplate/"
	ResourceDashboardLayout       Resource = "v1/dashboardlayout/"
	ResourceLimit                 Resource = "v1/limit/"
	ResourceMetric                Resource = "v1/metric/"
	ResourceMetricAnnotation      Resource = "v1/metricannotation/"
	ResourceMetricName            Resource = "v1/metricname/"
	ResourceMetricOperation       Resource = "v1/metricoperation/"
	ResourceMetricTemplateGroup   Resource = "v1/metrictemplategroup/"
	ResourceNotificationExecution Resource = "v1/notificationexecution/"
	ResourceNotificationRule      Resource = "v1/notificationrule/"
	ResourcePublicStatusPage      Resource = "v1/publicstatuspage/"
	ResourceRecordingRule         Resource = "v1/recordingrule/"
	ResourceReport                Resource = "v1/report/"
	ResourceReportConfig          Resource = "v1/reportconfig/"
	ResourceServerGroup           Resource = "v1/servergroup/"
	ResourceService               Resource = "v1/service/"
	ResourceSession               Resource = "v1/session/"
	ResourceSilence               Resource = "v1/silence/"
	ResourceSilenceRecurrent      Resource = "v1/silencerecurrent/"
	ResourceSlo                   Resource = "v1/slo/"
	ResourceTag                   Resource = "v1/tag/"
	ResourceUser                  Resource = "v1/user/"
	ResourceWidget                Resource = "v1/widget/"
	ResourceWidgetAnnotation      Resource = "v1/widgetannotation/"
)
