// Copyright 2015-2024 Bleemeo
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

type ArgumentTypeEnum int

const (
	ArgumentTypeEnum_None  ArgumentTypeEnum = 0
	ArgumentTypeEnum_Float ArgumentTypeEnum = 1
)

type AuditLogCategoryEnum int

const (
	AuditLogCategoryEnum_Account     AuditLogCategoryEnum = 0
	AuditLogCategoryEnum_Agent       AuditLogCategoryEnum = 1
	AuditLogCategoryEnum_Monitor     AuditLogCategoryEnum = 2
	AuditLogCategoryEnum_Dashboard   AuditLogCategoryEnum = 3
	AuditLogCategoryEnum_Integration AuditLogCategoryEnum = 4
	// Notification Rule
	AuditLogCategoryEnum_Notification_Rule AuditLogCategoryEnum = 5
	AuditLogCategoryEnum_Silence           AuditLogCategoryEnum = 6
	// AWS Integration
	AuditLogCategoryEnum_AWS_Integration AuditLogCategoryEnum = 7
	// Mobile App
	AuditLogCategoryEnum_Mobile_App AuditLogCategoryEnum = 8
	AuditLogCategoryEnum_Sandbox    AuditLogCategoryEnum = 9
	AuditLogCategoryEnum_Report     AuditLogCategoryEnum = 10
	AuditLogCategoryEnum_User       AuditLogCategoryEnum = 11
	AuditLogCategoryEnum_Plan       AuditLogCategoryEnum = 12
	// Payment Mean
	AuditLogCategoryEnum_Payment_Mean AuditLogCategoryEnum = 13
	AuditLogCategoryEnum_Reservations AuditLogCategoryEnum = 14
	// Contacts Group
	AuditLogCategoryEnum_Contacts_Group AuditLogCategoryEnum = 15
	// Server Group
	AuditLogCategoryEnum_Server_Group AuditLogCategoryEnum = 16
	AuditLogCategoryEnum_Invoice      AuditLogCategoryEnum = 17
)

type AuditLogTypeEnum int

const (
	AuditLogTypeEnum_Create   AuditLogTypeEnum = 0
	AuditLogTypeEnum_Edit     AuditLogTypeEnum = 1
	AuditLogTypeEnum_Connect  AuditLogTypeEnum = 2
	AuditLogTypeEnum_Shutdown AuditLogTypeEnum = 3
	AuditLogTypeEnum_Timeout  AuditLogTypeEnum = 4
	AuditLogTypeEnum_Upgrade  AuditLogTypeEnum = 5
	// Auto-upgrade
	AuditLogTypeEnum_Auto_upgrade AuditLogTypeEnum = 6
	AuditLogTypeEnum_Delete       AuditLogTypeEnum = 7
	AuditLogTypeEnum_Undelete     AuditLogTypeEnum = 8
	AuditLogTypeEnum_Purge        AuditLogTypeEnum = 9
	AuditLogTypeEnum_Deactivate   AuditLogTypeEnum = 10
	AuditLogTypeEnum_Reactivate   AuditLogTypeEnum = 11
	AuditLogTypeEnum_Info         AuditLogTypeEnum = 12
)

type CountryEnum string

const (
	CountryEnum_AF CountryEnum = "AF"
	CountryEnum_AX CountryEnum = "AX"
	CountryEnum_AL CountryEnum = "AL"
	CountryEnum_DZ CountryEnum = "DZ"
	CountryEnum_AS CountryEnum = "AS"
	CountryEnum_AD CountryEnum = "AD"
	CountryEnum_AO CountryEnum = "AO"
	CountryEnum_AI CountryEnum = "AI"
	CountryEnum_AQ CountryEnum = "AQ"
	CountryEnum_AG CountryEnum = "AG"
	CountryEnum_AR CountryEnum = "AR"
	CountryEnum_AM CountryEnum = "AM"
	CountryEnum_AW CountryEnum = "AW"
	CountryEnum_AU CountryEnum = "AU"
	CountryEnum_AT CountryEnum = "AT"
	CountryEnum_AZ CountryEnum = "AZ"
	CountryEnum_BS CountryEnum = "BS"
	CountryEnum_BH CountryEnum = "BH"
	CountryEnum_BD CountryEnum = "BD"
	CountryEnum_BB CountryEnum = "BB"
	CountryEnum_BY CountryEnum = "BY"
	CountryEnum_BE CountryEnum = "BE"
	CountryEnum_BZ CountryEnum = "BZ"
	CountryEnum_BJ CountryEnum = "BJ"
	CountryEnum_BM CountryEnum = "BM"
	CountryEnum_BT CountryEnum = "BT"
	CountryEnum_BO CountryEnum = "BO"
	CountryEnum_BQ CountryEnum = "BQ"
	CountryEnum_BA CountryEnum = "BA"
	CountryEnum_BW CountryEnum = "BW"
	CountryEnum_BV CountryEnum = "BV"
	CountryEnum_BR CountryEnum = "BR"
	CountryEnum_IO CountryEnum = "IO"
	CountryEnum_BN CountryEnum = "BN"
	CountryEnum_BG CountryEnum = "BG"
	CountryEnum_BF CountryEnum = "BF"
	CountryEnum_BI CountryEnum = "BI"
	CountryEnum_CV CountryEnum = "CV"
	CountryEnum_KH CountryEnum = "KH"
	CountryEnum_CM CountryEnum = "CM"
	CountryEnum_CA CountryEnum = "CA"
	CountryEnum_KY CountryEnum = "KY"
	CountryEnum_CF CountryEnum = "CF"
	CountryEnum_TD CountryEnum = "TD"
	CountryEnum_CL CountryEnum = "CL"
	CountryEnum_CN CountryEnum = "CN"
	CountryEnum_CX CountryEnum = "CX"
	CountryEnum_CC CountryEnum = "CC"
	CountryEnum_CO CountryEnum = "CO"
	CountryEnum_KM CountryEnum = "KM"
	CountryEnum_CG CountryEnum = "CG"
	CountryEnum_CD CountryEnum = "CD"
	CountryEnum_CK CountryEnum = "CK"
	CountryEnum_CR CountryEnum = "CR"
	CountryEnum_CI CountryEnum = "CI"
	CountryEnum_HR CountryEnum = "HR"
	CountryEnum_CU CountryEnum = "CU"
	CountryEnum_CW CountryEnum = "CW"
	CountryEnum_CY CountryEnum = "CY"
	CountryEnum_CZ CountryEnum = "CZ"
	CountryEnum_DK CountryEnum = "DK"
	CountryEnum_DJ CountryEnum = "DJ"
	CountryEnum_DM CountryEnum = "DM"
	CountryEnum_DO CountryEnum = "DO"
	CountryEnum_EC CountryEnum = "EC"
	CountryEnum_EG CountryEnum = "EG"
	CountryEnum_SV CountryEnum = "SV"
	CountryEnum_GQ CountryEnum = "GQ"
	CountryEnum_ER CountryEnum = "ER"
	CountryEnum_EE CountryEnum = "EE"
	CountryEnum_SZ CountryEnum = "SZ"
	CountryEnum_ET CountryEnum = "ET"
	CountryEnum_FK CountryEnum = "FK"
	CountryEnum_FO CountryEnum = "FO"
	CountryEnum_FJ CountryEnum = "FJ"
	CountryEnum_FI CountryEnum = "FI"
	CountryEnum_FR CountryEnum = "FR"
	CountryEnum_GF CountryEnum = "GF"
	CountryEnum_PF CountryEnum = "PF"
	CountryEnum_TF CountryEnum = "TF"
	CountryEnum_GA CountryEnum = "GA"
	CountryEnum_GM CountryEnum = "GM"
	CountryEnum_GE CountryEnum = "GE"
	CountryEnum_DE CountryEnum = "DE"
	CountryEnum_GH CountryEnum = "GH"
	CountryEnum_GI CountryEnum = "GI"
	CountryEnum_GR CountryEnum = "GR"
	CountryEnum_GL CountryEnum = "GL"
	CountryEnum_GD CountryEnum = "GD"
	CountryEnum_GP CountryEnum = "GP"
	CountryEnum_GU CountryEnum = "GU"
	CountryEnum_GT CountryEnum = "GT"
	CountryEnum_GG CountryEnum = "GG"
	CountryEnum_GN CountryEnum = "GN"
	CountryEnum_GW CountryEnum = "GW"
	CountryEnum_GY CountryEnum = "GY"
	CountryEnum_HT CountryEnum = "HT"
	CountryEnum_HM CountryEnum = "HM"
	CountryEnum_VA CountryEnum = "VA"
	CountryEnum_HN CountryEnum = "HN"
	CountryEnum_HK CountryEnum = "HK"
	CountryEnum_HU CountryEnum = "HU"
	CountryEnum_IS CountryEnum = "IS"
	CountryEnum_IN CountryEnum = "IN"
	CountryEnum_ID CountryEnum = "ID"
	CountryEnum_IR CountryEnum = "IR"
	CountryEnum_IQ CountryEnum = "IQ"
	CountryEnum_IE CountryEnum = "IE"
	CountryEnum_IM CountryEnum = "IM"
	CountryEnum_IL CountryEnum = "IL"
	CountryEnum_IT CountryEnum = "IT"
	CountryEnum_JM CountryEnum = "JM"
	CountryEnum_JP CountryEnum = "JP"
	CountryEnum_JE CountryEnum = "JE"
	CountryEnum_JO CountryEnum = "JO"
	CountryEnum_KZ CountryEnum = "KZ"
	CountryEnum_KE CountryEnum = "KE"
	CountryEnum_KI CountryEnum = "KI"
	CountryEnum_KW CountryEnum = "KW"
	CountryEnum_KG CountryEnum = "KG"
	CountryEnum_LA CountryEnum = "LA"
	CountryEnum_LV CountryEnum = "LV"
	CountryEnum_LB CountryEnum = "LB"
	CountryEnum_LS CountryEnum = "LS"
	CountryEnum_LR CountryEnum = "LR"
	CountryEnum_LY CountryEnum = "LY"
	CountryEnum_LI CountryEnum = "LI"
	CountryEnum_LT CountryEnum = "LT"
	CountryEnum_LU CountryEnum = "LU"
	CountryEnum_MO CountryEnum = "MO"
	CountryEnum_MG CountryEnum = "MG"
	CountryEnum_MW CountryEnum = "MW"
	CountryEnum_MY CountryEnum = "MY"
	CountryEnum_MV CountryEnum = "MV"
	CountryEnum_ML CountryEnum = "ML"
	CountryEnum_MT CountryEnum = "MT"
	CountryEnum_MH CountryEnum = "MH"
	CountryEnum_MQ CountryEnum = "MQ"
	CountryEnum_MR CountryEnum = "MR"
	CountryEnum_MU CountryEnum = "MU"
	CountryEnum_YT CountryEnum = "YT"
	CountryEnum_MX CountryEnum = "MX"
	CountryEnum_FM CountryEnum = "FM"
	CountryEnum_MD CountryEnum = "MD"
	CountryEnum_MC CountryEnum = "MC"
	CountryEnum_MN CountryEnum = "MN"
	CountryEnum_ME CountryEnum = "ME"
	CountryEnum_MS CountryEnum = "MS"
	CountryEnum_MA CountryEnum = "MA"
	CountryEnum_MZ CountryEnum = "MZ"
	CountryEnum_MM CountryEnum = "MM"
	CountryEnum_NA CountryEnum = "NA"
	CountryEnum_NR CountryEnum = "NR"
	CountryEnum_NP CountryEnum = "NP"
	CountryEnum_NL CountryEnum = "NL"
	CountryEnum_NC CountryEnum = "NC"
	CountryEnum_NZ CountryEnum = "NZ"
	CountryEnum_NI CountryEnum = "NI"
	CountryEnum_NE CountryEnum = "NE"
	CountryEnum_NG CountryEnum = "NG"
	CountryEnum_NU CountryEnum = "NU"
	CountryEnum_NF CountryEnum = "NF"
	CountryEnum_KP CountryEnum = "KP"
	CountryEnum_MK CountryEnum = "MK"
	CountryEnum_MP CountryEnum = "MP"
	CountryEnum_NO CountryEnum = "NO"
	CountryEnum_OM CountryEnum = "OM"
	CountryEnum_PK CountryEnum = "PK"
	CountryEnum_PW CountryEnum = "PW"
	CountryEnum_PS CountryEnum = "PS"
	CountryEnum_PA CountryEnum = "PA"
	CountryEnum_PG CountryEnum = "PG"
	CountryEnum_PY CountryEnum = "PY"
	CountryEnum_PE CountryEnum = "PE"
	CountryEnum_PH CountryEnum = "PH"
	CountryEnum_PN CountryEnum = "PN"
	CountryEnum_PL CountryEnum = "PL"
	CountryEnum_PT CountryEnum = "PT"
	CountryEnum_PR CountryEnum = "PR"
	CountryEnum_QA CountryEnum = "QA"
	CountryEnum_RE CountryEnum = "RE"
	CountryEnum_RO CountryEnum = "RO"
	CountryEnum_RU CountryEnum = "RU"
	CountryEnum_RW CountryEnum = "RW"
	CountryEnum_BL CountryEnum = "BL"
	CountryEnum_SH CountryEnum = "SH"
	CountryEnum_KN CountryEnum = "KN"
	CountryEnum_LC CountryEnum = "LC"
	CountryEnum_MF CountryEnum = "MF"
	CountryEnum_PM CountryEnum = "PM"
	CountryEnum_VC CountryEnum = "VC"
	CountryEnum_WS CountryEnum = "WS"
	CountryEnum_SM CountryEnum = "SM"
	CountryEnum_ST CountryEnum = "ST"
	CountryEnum_SA CountryEnum = "SA"
	CountryEnum_SN CountryEnum = "SN"
	CountryEnum_RS CountryEnum = "RS"
	CountryEnum_SC CountryEnum = "SC"
	CountryEnum_SL CountryEnum = "SL"
	CountryEnum_SG CountryEnum = "SG"
	CountryEnum_SX CountryEnum = "SX"
	CountryEnum_SK CountryEnum = "SK"
	CountryEnum_SI CountryEnum = "SI"
	CountryEnum_SB CountryEnum = "SB"
	CountryEnum_SO CountryEnum = "SO"
	CountryEnum_ZA CountryEnum = "ZA"
	CountryEnum_GS CountryEnum = "GS"
	CountryEnum_KR CountryEnum = "KR"
	CountryEnum_SS CountryEnum = "SS"
	CountryEnum_ES CountryEnum = "ES"
	CountryEnum_LK CountryEnum = "LK"
	CountryEnum_SD CountryEnum = "SD"
	CountryEnum_SR CountryEnum = "SR"
	CountryEnum_SJ CountryEnum = "SJ"
	CountryEnum_SE CountryEnum = "SE"
	CountryEnum_CH CountryEnum = "CH"
	CountryEnum_SY CountryEnum = "SY"
	CountryEnum_TW CountryEnum = "TW"
	CountryEnum_TJ CountryEnum = "TJ"
	CountryEnum_TZ CountryEnum = "TZ"
	CountryEnum_TH CountryEnum = "TH"
	CountryEnum_TL CountryEnum = "TL"
	CountryEnum_TG CountryEnum = "TG"
	CountryEnum_TK CountryEnum = "TK"
	CountryEnum_TO CountryEnum = "TO"
	CountryEnum_TT CountryEnum = "TT"
	CountryEnum_TN CountryEnum = "TN"
	CountryEnum_TR CountryEnum = "TR"
	CountryEnum_TM CountryEnum = "TM"
	CountryEnum_TC CountryEnum = "TC"
	CountryEnum_TV CountryEnum = "TV"
	CountryEnum_UG CountryEnum = "UG"
	CountryEnum_UA CountryEnum = "UA"
	CountryEnum_AE CountryEnum = "AE"
	CountryEnum_GB CountryEnum = "GB"
	CountryEnum_UM CountryEnum = "UM"
	CountryEnum_US CountryEnum = "US"
	CountryEnum_UY CountryEnum = "UY"
	CountryEnum_UZ CountryEnum = "UZ"
	CountryEnum_VU CountryEnum = "VU"
	CountryEnum_VE CountryEnum = "VE"
	CountryEnum_VN CountryEnum = "VN"
	CountryEnum_VG CountryEnum = "VG"
	CountryEnum_VI CountryEnum = "VI"
	CountryEnum_WF CountryEnum = "WF"
	CountryEnum_EH CountryEnum = "EH"
	CountryEnum_YE CountryEnum = "YE"
	CountryEnum_ZM CountryEnum = "ZM"
	CountryEnum_ZW CountryEnum = "ZW"
)

type CurrentStatusEnum int

const (
	CurrentStatusEnum_ok       CurrentStatusEnum = 0
	CurrentStatusEnum_warning  CurrentStatusEnum = 1
	CurrentStatusEnum_critical CurrentStatusEnum = 2
	CurrentStatusEnum_unknown  CurrentStatusEnum = 3
)

type DeactivatedReasonEnum int

const (
	// I don't use the solution anymore
	DeactivatedReasonEnum_I_don_t_use_the_solution_anymore DeactivatedReasonEnum = 0
	// The tool setup is too complicated
	DeactivatedReasonEnum_The_tool_setup_is_too_complicated DeactivatedReasonEnum = 1
	// Doesn't meet my expectations
	DeactivatedReasonEnum_Doesn_t_meet_my_expectations DeactivatedReasonEnum = 2
	// I never have downtime, no need monitoring :)
	DeactivatedReasonEnum_I_never_have_downtime_no_need_monitoring DeactivatedReasonEnum = 3
	// Another solution suits me better
	DeactivatedReasonEnum_Another_solution_suits_me_better DeactivatedReasonEnum = 4
	// One essential feature is missing
	DeactivatedReasonEnum_One_essential_feature_is_missing DeactivatedReasonEnum = 5
	DeactivatedReasonEnum_Other                            DeactivatedReasonEnum = 6
)

type FCMDeviceTypeEnum string

const (
	FCMDeviceTypeEnum_ios     FCMDeviceTypeEnum = "ios"
	FCMDeviceTypeEnum_android FCMDeviceTypeEnum = "android"
	FCMDeviceTypeEnum_web     FCMDeviceTypeEnum = "web"
)

type GloutonConfigItemTypeEnum int

const (
	GloutonConfigItemTypeEnum_any    GloutonConfigItemTypeEnum = 0
	GloutonConfigItemTypeEnum_int    GloutonConfigItemTypeEnum = 1
	GloutonConfigItemTypeEnum_float  GloutonConfigItemTypeEnum = 2
	GloutonConfigItemTypeEnum_bool   GloutonConfigItemTypeEnum = 3
	GloutonConfigItemTypeEnum_string GloutonConfigItemTypeEnum = 4
	// list string
	GloutonConfigItemTypeEnum_list_string GloutonConfigItemTypeEnum = 10
	// list int
	GloutonConfigItemTypeEnum_list_int GloutonConfigItemTypeEnum = 11
	// map string string
	GloutonConfigItemTypeEnum_map_string_string GloutonConfigItemTypeEnum = 20
	// map string int
	GloutonConfigItemTypeEnum_map_string_int GloutonConfigItemTypeEnum = 21
	GloutonConfigItemTypeEnum_thresholds     GloutonConfigItemTypeEnum = 30
	GloutonConfigItemTypeEnum_services       GloutonConfigItemTypeEnum = 31
	// name instances
	GloutonConfigItemTypeEnum_name_instances GloutonConfigItemTypeEnum = 32
	// Blackbox targets
	GloutonConfigItemTypeEnum_Blackbox_targets GloutonConfigItemTypeEnum = 33
	// Prometheus targets
	GloutonConfigItemTypeEnum_Prometheus_targets GloutonConfigItemTypeEnum = 34
	// SNMP targets
	GloutonConfigItemTypeEnum_SNMP_targets GloutonConfigItemTypeEnum = 35
	// log inputs
	GloutonConfigItemTypeEnum_log_inputs GloutonConfigItemTypeEnum = 36
)

type GloutonDiagnosticTypeEnum int

const (
	GloutonDiagnosticTypeEnum_Crash GloutonDiagnosticTypeEnum = 0
	// On demand
	GloutonDiagnosticTypeEnum_On_demand GloutonDiagnosticTypeEnum = 1
)

type GraphEnum int

const (
	// Line Chart
	GraphEnum_Line_Chart GraphEnum = 0
	// Stacked Area Chart
	GraphEnum_Stacked_Area_Chart GraphEnum = 1
	// Pie Chart
	GraphEnum_Pie_Chart GraphEnum = 2
	GraphEnum_Gauge     GraphEnum = 3
	// Status History Chart
	GraphEnum_Status_History_Chart GraphEnum = 4
	// Metric Value
	GraphEnum_Metric_Value GraphEnum = 5
	GraphEnum_Status       GraphEnum = 6
	// SNMP Status
	GraphEnum_SNMP_Status GraphEnum = 7
	GraphEnum_Text        GraphEnum = 8
	GraphEnum_Image       GraphEnum = 9
	// Heatmap Status
	GraphEnum_Heatmap_Status GraphEnum = 10
	// Bar Chart
	GraphEnum_Bar_Chart GraphEnum = 11
)

type GraphSubtypeEnum int

const (
	// Status Round
	GraphSubtypeEnum_Status_Round GraphSubtypeEnum = 0
	// Status Smiley
	GraphSubtypeEnum_Status_Smiley GraphSubtypeEnum = 1
	// Status Image
	GraphSubtypeEnum_Status_Image GraphSubtypeEnum = 2
	// Heatmap with Value
	GraphSubtypeEnum_Heatmap_with_Value GraphSubtypeEnum = 100
	// Heatmap with Status
	GraphSubtypeEnum_Heatmap_with_Status GraphSubtypeEnum = 101
)

type HealthCheckCategoryEnum int

const (
	// Bleemeo Account
	HealthCheckCategoryEnum_Bleemeo_Account HealthCheckCategoryEnum = 1
	// Bleemeo Server
	HealthCheckCategoryEnum_Bleemeo_Server HealthCheckCategoryEnum = 2
	HealthCheckCategoryEnum_Server         HealthCheckCategoryEnum = 3
)

type HealthCheckNameEnum int

const (
	// 2FA for users
	HealthCheckNameEnum_2FA_for_users HealthCheckNameEnum = 101
	// K8S integration
	HealthCheckNameEnum_K8S_integration HealthCheckNameEnum = 102
	// Report creation
	HealthCheckNameEnum_Report_creation HealthCheckNameEnum = 103
	// Application creation
	HealthCheckNameEnum_Application_creation HealthCheckNameEnum = 104
	// Upgrade plan
	HealthCheckNameEnum_Upgrade_plan HealthCheckNameEnum = 105
	// Reaching account limits
	HealthCheckNameEnum_Reaching_account_limits HealthCheckNameEnum = 106
	// Passwords age
	HealthCheckNameEnum_Passwords_age HealthCheckNameEnum = 107
	// Suggest reservation
	HealthCheckNameEnum_Suggest_reservation HealthCheckNameEnum = 108
	// Glouton auto-upgrade
	HealthCheckNameEnum_Glouton_auto_upgrade HealthCheckNameEnum = 201
	// Duplicated agent
	HealthCheckNameEnum_Duplicated_agent HealthCheckNameEnum = 202
	// Agent outdated
	HealthCheckNameEnum_Agent_outdated HealthCheckNameEnum = 203
	// Configuration warning
	HealthCheckNameEnum_Configuration_warning HealthCheckNameEnum = 204
	// Service configuration needed
	HealthCheckNameEnum_Service_configuration_needed HealthCheckNameEnum = 205
	// NRPE detected
	HealthCheckNameEnum_NRPE_detected HealthCheckNameEnum = 206
	// Reaching agent limits
	HealthCheckNameEnum_Reaching_agent_limits HealthCheckNameEnum = 207
	// Smartctl install
	HealthCheckNameEnum_Smartctl_install HealthCheckNameEnum = 208
	// IPMI install
	HealthCheckNameEnum_IPMI_install HealthCheckNameEnum = 209
	// Swap needed
	HealthCheckNameEnum_Swap_needed HealthCheckNameEnum = 301
	// CPU usage could lead to server upgrade/downgrade
	HealthCheckNameEnum_CPU_usage_could_lead_to_server_upgrade_downgrade HealthCheckNameEnum = 302
	// Mem usage could lead to server upgrade/downgrade
	HealthCheckNameEnum_Mem_usage_could_lead_to_server_upgrade_downgrade HealthCheckNameEnum = 303
	// Machine Learning: prediction of memory/disk full
	HealthCheckNameEnum_Machine_Learning_prediction_of_memory_disk_full HealthCheckNameEnum = 304
	// Pending security upgrades
	HealthCheckNameEnum_Pending_security_upgrades HealthCheckNameEnum = 305
	// SSL: checks on cypher authorized and protocols versions
	HealthCheckNameEnum_SSL_checks_on_cypher_authorized_and_protocols_versions HealthCheckNameEnum = 306
)

type IconEnum int

const (
	// bi bi-hdd-stack
	IconEnum_bi_bi_hdd_stack IconEnum = 1
	// bi bi-laptop
	IconEnum_bi_bi_laptop IconEnum = 2
	// bi bi-database
	IconEnum_bi_bi_database IconEnum = 3
	// bi bi-wifi
	IconEnum_bi_bi_wifi IconEnum = 4
	// bi bi-printer
	IconEnum_bi_bi_printer IconEnum = 5
	// bi bi-phone
	IconEnum_bi_bi_phone IconEnum = 6
	// bi bi-hdd-network
	IconEnum_bi_bi_hdd_network IconEnum = 7
	// bi bi-hdd
	IconEnum_bi_bi_hdd IconEnum = 8
	// bi bi-ethernet
	IconEnum_bi_bi_ethernet IconEnum = 9
	// bi bi-plug
	IconEnum_bi_bi_plug IconEnum = 10
	// bi bi-life-preserver
	IconEnum_bi_bi_life_preserver IconEnum = 11
)

type IntegrationTemplateNameEnum string

const (
	IntegrationTemplateNameEnum_email              IntegrationTemplateNameEnum = "email"
	IntegrationTemplateNameEnum_Slack_Webhook      IntegrationTemplateNameEnum = "Slack Webhook"
	IntegrationTemplateNameEnum_Slack              IntegrationTemplateNameEnum = "Slack"
	IntegrationTemplateNameEnum_mobile_application IntegrationTemplateNameEnum = "mobile application"
	IntegrationTemplateNameEnum_webhook            IntegrationTemplateNameEnum = "webhook"
	IntegrationTemplateNameEnum_PagerDuty          IntegrationTemplateNameEnum = "PagerDuty"
	IntegrationTemplateNameEnum_Twilio             IntegrationTemplateNameEnum = "Twilio"
	IntegrationTemplateNameEnum_VictorOps          IntegrationTemplateNameEnum = "VictorOps"
	IntegrationTemplateNameEnum_Opsgenie           IntegrationTemplateNameEnum = "Opsgenie"
	IntegrationTemplateNameEnum_MessageBird        IntegrationTemplateNameEnum = "MessageBird"
	IntegrationTemplateNameEnum_Teams              IntegrationTemplateNameEnum = "Teams"
	IntegrationTemplateNameEnum_OVH_SMS            IntegrationTemplateNameEnum = "OVH SMS"
	IntegrationTemplateNameEnum_Telegram           IntegrationTemplateNameEnum = "Telegram"
)

type IntervalEnum int

const (
	IntervalEnum_1h IntervalEnum = 0
	IntervalEnum_6h IntervalEnum = 2
	IntervalEnum_1d IntervalEnum = 3
	IntervalEnum_1w IntervalEnum = 4
	IntervalEnum_1m IntervalEnum = 5
	IntervalEnum_1y IntervalEnum = 6
)

type LastDisconnectionReasonEnum int

const (
	// Agent clean shutdown
	LastDisconnectionReasonEnum_Agent_clean_shutdown LastDisconnectionReasonEnum = 1
	// Agent timeout
	LastDisconnectionReasonEnum_Agent_timeout LastDisconnectionReasonEnum = 2
	// Agent auto-upgrade
	LastDisconnectionReasonEnum_Agent_auto_upgrade LastDisconnectionReasonEnum = 3
	// Agent upgrade
	LastDisconnectionReasonEnum_Agent_upgrade LastDisconnectionReasonEnum = 4
)

type LimitTypeEnum int

const (
	LimitTypeEnum_Server  LimitTypeEnum = 1
	LimitTypeEnum_Monitor LimitTypeEnum = 2
	// Recording Rule
	LimitTypeEnum_Recording_Rule LimitTypeEnum = 3
	// Recording Rule metric
	LimitTypeEnum_Recording_Rule_metric LimitTypeEnum = 4
	// Agent Type
	LimitTypeEnum_Agent_Type LimitTypeEnum = 10
)

type NullEnum string

const (
	NullEnum_NULL NullEnum = "null"
)

type PeriodEnum int

const (
	PeriodEnum_Weekly  PeriodEnum = 0
	PeriodEnum_Monthly PeriodEnum = 1
)

type ReportEnum int

const (
	// No report
	ReportEnum_No_report ReportEnum = 0
	// Only weekly report
	ReportEnum_Only_weekly_report ReportEnum = 1
	// Full report
	ReportEnum_Full_report ReportEnum = 2
)

type SeverityEnum int

const (
	SeverityEnum_ok       SeverityEnum = 0
	SeverityEnum_warning  SeverityEnum = 1
	SeverityEnum_critical SeverityEnum = 2
)

type SourceEnum int

const (
	SourceEnum_unknown SourceEnum = 0
	SourceEnum_default SourceEnum = 1
	SourceEnum_file    SourceEnum = 2
	SourceEnum_env     SourceEnum = 3
	SourceEnum_api     SourceEnum = 4
)

type StatusEnum int

const (
	StatusEnum_OK       StatusEnum = 0
	StatusEnum_Warning  StatusEnum = 1
	StatusEnum_Critical StatusEnum = 2
	StatusEnum_Unknown  StatusEnum = 3
	StatusEnum_Info     StatusEnum = 10
)

type TagTypeEnum int

const (
	// Is automatic
	TagTypeEnum_Is_automatic TagTypeEnum = 0
	// Created by Glouton
	TagTypeEnum_Created_by_Glouton TagTypeEnum = 1
	// Created by frontend
	TagTypeEnum_Created_by_frontend TagTypeEnum = 2
	// No type
	TagTypeEnum_No_type TagTypeEnum = 10
)

type TargetTypeEnum int

const (
	TargetTypeEnum_none  TargetTypeEnum = 0
	TargetTypeEnum_url   TargetTypeEnum = 6
	TargetTypeEnum_email TargetTypeEnum = 7
	TargetTypeEnum_user  TargetTypeEnum = 9
	// phone number
	TargetTypeEnum_phone_number TargetTypeEnum = 8
	TargetTypeEnum_string       TargetTypeEnum = 1
)

type TimezoneEnum string

const (
	TimezoneEnum_Africa_Abidjan                 TimezoneEnum = "Africa/Abidjan"
	TimezoneEnum_Africa_Accra                   TimezoneEnum = "Africa/Accra"
	TimezoneEnum_Africa_Addis_Ababa             TimezoneEnum = "Africa/Addis_Ababa"
	TimezoneEnum_Africa_Algiers                 TimezoneEnum = "Africa/Algiers"
	TimezoneEnum_Africa_Asmara                  TimezoneEnum = "Africa/Asmara"
	TimezoneEnum_Africa_Bamako                  TimezoneEnum = "Africa/Bamako"
	TimezoneEnum_Africa_Bangui                  TimezoneEnum = "Africa/Bangui"
	TimezoneEnum_Africa_Banjul                  TimezoneEnum = "Africa/Banjul"
	TimezoneEnum_Africa_Bissau                  TimezoneEnum = "Africa/Bissau"
	TimezoneEnum_Africa_Blantyre                TimezoneEnum = "Africa/Blantyre"
	TimezoneEnum_Africa_Brazzaville             TimezoneEnum = "Africa/Brazzaville"
	TimezoneEnum_Africa_Bujumbura               TimezoneEnum = "Africa/Bujumbura"
	TimezoneEnum_Africa_Cairo                   TimezoneEnum = "Africa/Cairo"
	TimezoneEnum_Africa_Casablanca              TimezoneEnum = "Africa/Casablanca"
	TimezoneEnum_Africa_Ceuta                   TimezoneEnum = "Africa/Ceuta"
	TimezoneEnum_Africa_Conakry                 TimezoneEnum = "Africa/Conakry"
	TimezoneEnum_Africa_Dakar                   TimezoneEnum = "Africa/Dakar"
	TimezoneEnum_Africa_Dar_es_Salaam           TimezoneEnum = "Africa/Dar_es_Salaam"
	TimezoneEnum_Africa_Djibouti                TimezoneEnum = "Africa/Djibouti"
	TimezoneEnum_Africa_Douala                  TimezoneEnum = "Africa/Douala"
	TimezoneEnum_Africa_El_Aaiun                TimezoneEnum = "Africa/El_Aaiun"
	TimezoneEnum_Africa_Freetown                TimezoneEnum = "Africa/Freetown"
	TimezoneEnum_Africa_Gaborone                TimezoneEnum = "Africa/Gaborone"
	TimezoneEnum_Africa_Harare                  TimezoneEnum = "Africa/Harare"
	TimezoneEnum_Africa_Johannesburg            TimezoneEnum = "Africa/Johannesburg"
	TimezoneEnum_Africa_Juba                    TimezoneEnum = "Africa/Juba"
	TimezoneEnum_Africa_Kampala                 TimezoneEnum = "Africa/Kampala"
	TimezoneEnum_Africa_Khartoum                TimezoneEnum = "Africa/Khartoum"
	TimezoneEnum_Africa_Kigali                  TimezoneEnum = "Africa/Kigali"
	TimezoneEnum_Africa_Kinshasa                TimezoneEnum = "Africa/Kinshasa"
	TimezoneEnum_Africa_Lagos                   TimezoneEnum = "Africa/Lagos"
	TimezoneEnum_Africa_Libreville              TimezoneEnum = "Africa/Libreville"
	TimezoneEnum_Africa_Lome                    TimezoneEnum = "Africa/Lome"
	TimezoneEnum_Africa_Luanda                  TimezoneEnum = "Africa/Luanda"
	TimezoneEnum_Africa_Lubumbashi              TimezoneEnum = "Africa/Lubumbashi"
	TimezoneEnum_Africa_Lusaka                  TimezoneEnum = "Africa/Lusaka"
	TimezoneEnum_Africa_Malabo                  TimezoneEnum = "Africa/Malabo"
	TimezoneEnum_Africa_Maputo                  TimezoneEnum = "Africa/Maputo"
	TimezoneEnum_Africa_Maseru                  TimezoneEnum = "Africa/Maseru"
	TimezoneEnum_Africa_Mbabane                 TimezoneEnum = "Africa/Mbabane"
	TimezoneEnum_Africa_Mogadishu               TimezoneEnum = "Africa/Mogadishu"
	TimezoneEnum_Africa_Monrovia                TimezoneEnum = "Africa/Monrovia"
	TimezoneEnum_Africa_Nairobi                 TimezoneEnum = "Africa/Nairobi"
	TimezoneEnum_Africa_Ndjamena                TimezoneEnum = "Africa/Ndjamena"
	TimezoneEnum_Africa_Niamey                  TimezoneEnum = "Africa/Niamey"
	TimezoneEnum_Africa_Nouakchott              TimezoneEnum = "Africa/Nouakchott"
	TimezoneEnum_Africa_Ouagadougou             TimezoneEnum = "Africa/Ouagadougou"
	TimezoneEnum_Africa_Porto_Novo              TimezoneEnum = "Africa/Porto-Novo"
	TimezoneEnum_Africa_Sao_Tome                TimezoneEnum = "Africa/Sao_Tome"
	TimezoneEnum_Africa_Tripoli                 TimezoneEnum = "Africa/Tripoli"
	TimezoneEnum_Africa_Tunis                   TimezoneEnum = "Africa/Tunis"
	TimezoneEnum_Africa_Windhoek                TimezoneEnum = "Africa/Windhoek"
	TimezoneEnum_America_Adak                   TimezoneEnum = "America/Adak"
	TimezoneEnum_America_Anchorage              TimezoneEnum = "America/Anchorage"
	TimezoneEnum_America_Anguilla               TimezoneEnum = "America/Anguilla"
	TimezoneEnum_America_Antigua                TimezoneEnum = "America/Antigua"
	TimezoneEnum_America_Araguaina              TimezoneEnum = "America/Araguaina"
	TimezoneEnum_America_Argentina_Buenos_Aires TimezoneEnum = "America/Argentina/Buenos_Aires"
	TimezoneEnum_America_Argentina_Catamarca    TimezoneEnum = "America/Argentina/Catamarca"
	TimezoneEnum_America_Argentina_Cordoba      TimezoneEnum = "America/Argentina/Cordoba"
	TimezoneEnum_America_Argentina_Jujuy        TimezoneEnum = "America/Argentina/Jujuy"
	TimezoneEnum_America_Argentina_La_Rioja     TimezoneEnum = "America/Argentina/La_Rioja"
	TimezoneEnum_America_Argentina_Mendoza      TimezoneEnum = "America/Argentina/Mendoza"
	TimezoneEnum_America_Argentina_Rio_Gallegos TimezoneEnum = "America/Argentina/Rio_Gallegos"
	TimezoneEnum_America_Argentina_Salta        TimezoneEnum = "America/Argentina/Salta"
	TimezoneEnum_America_Argentina_San_Juan     TimezoneEnum = "America/Argentina/San_Juan"
	TimezoneEnum_America_Argentina_San_Luis     TimezoneEnum = "America/Argentina/San_Luis"
	TimezoneEnum_America_Argentina_Tucuman      TimezoneEnum = "America/Argentina/Tucuman"
	TimezoneEnum_America_Argentina_Ushuaia      TimezoneEnum = "America/Argentina/Ushuaia"
	TimezoneEnum_America_Aruba                  TimezoneEnum = "America/Aruba"
	TimezoneEnum_America_Asuncion               TimezoneEnum = "America/Asuncion"
	TimezoneEnum_America_Atikokan               TimezoneEnum = "America/Atikokan"
	TimezoneEnum_America_Bahia                  TimezoneEnum = "America/Bahia"
	TimezoneEnum_America_Bahia_Banderas         TimezoneEnum = "America/Bahia_Banderas"
	TimezoneEnum_America_Barbados               TimezoneEnum = "America/Barbados"
	TimezoneEnum_America_Belem                  TimezoneEnum = "America/Belem"
	TimezoneEnum_America_Belize                 TimezoneEnum = "America/Belize"
	TimezoneEnum_America_Blanc_Sablon           TimezoneEnum = "America/Blanc-Sablon"
	TimezoneEnum_America_Boa_Vista              TimezoneEnum = "America/Boa_Vista"
	TimezoneEnum_America_Bogota                 TimezoneEnum = "America/Bogota"
	TimezoneEnum_America_Boise                  TimezoneEnum = "America/Boise"
	TimezoneEnum_America_Cambridge_Bay          TimezoneEnum = "America/Cambridge_Bay"
	TimezoneEnum_America_Campo_Grande           TimezoneEnum = "America/Campo_Grande"
	TimezoneEnum_America_Cancun                 TimezoneEnum = "America/Cancun"
	TimezoneEnum_America_Caracas                TimezoneEnum = "America/Caracas"
	TimezoneEnum_America_Cayenne                TimezoneEnum = "America/Cayenne"
	TimezoneEnum_America_Cayman                 TimezoneEnum = "America/Cayman"
	TimezoneEnum_America_Chicago                TimezoneEnum = "America/Chicago"
	TimezoneEnum_America_Chihuahua              TimezoneEnum = "America/Chihuahua"
	TimezoneEnum_America_Ciudad_Juarez          TimezoneEnum = "America/Ciudad_Juarez"
	TimezoneEnum_America_Costa_Rica             TimezoneEnum = "America/Costa_Rica"
	TimezoneEnum_America_Creston                TimezoneEnum = "America/Creston"
	TimezoneEnum_America_Cuiaba                 TimezoneEnum = "America/Cuiaba"
	TimezoneEnum_America_Curacao                TimezoneEnum = "America/Curacao"
	TimezoneEnum_America_Danmarkshavn           TimezoneEnum = "America/Danmarkshavn"
	TimezoneEnum_America_Dawson                 TimezoneEnum = "America/Dawson"
	TimezoneEnum_America_Dawson_Creek           TimezoneEnum = "America/Dawson_Creek"
	TimezoneEnum_America_Denver                 TimezoneEnum = "America/Denver"
	TimezoneEnum_America_Detroit                TimezoneEnum = "America/Detroit"
	TimezoneEnum_America_Dominica               TimezoneEnum = "America/Dominica"
	TimezoneEnum_America_Edmonton               TimezoneEnum = "America/Edmonton"
	TimezoneEnum_America_Eirunepe               TimezoneEnum = "America/Eirunepe"
	TimezoneEnum_America_El_Salvador            TimezoneEnum = "America/El_Salvador"
	TimezoneEnum_America_Fort_Nelson            TimezoneEnum = "America/Fort_Nelson"
	TimezoneEnum_America_Fortaleza              TimezoneEnum = "America/Fortaleza"
	TimezoneEnum_America_Glace_Bay              TimezoneEnum = "America/Glace_Bay"
	TimezoneEnum_America_Goose_Bay              TimezoneEnum = "America/Goose_Bay"
	TimezoneEnum_America_Grand_Turk             TimezoneEnum = "America/Grand_Turk"
	TimezoneEnum_America_Grenada                TimezoneEnum = "America/Grenada"
	TimezoneEnum_America_Guadeloupe             TimezoneEnum = "America/Guadeloupe"
	TimezoneEnum_America_Guatemala              TimezoneEnum = "America/Guatemala"
	TimezoneEnum_America_Guayaquil              TimezoneEnum = "America/Guayaquil"
	TimezoneEnum_America_Guyana                 TimezoneEnum = "America/Guyana"
	TimezoneEnum_America_Halifax                TimezoneEnum = "America/Halifax"
	TimezoneEnum_America_Havana                 TimezoneEnum = "America/Havana"
	TimezoneEnum_America_Hermosillo             TimezoneEnum = "America/Hermosillo"
	TimezoneEnum_America_Indiana_Indianapolis   TimezoneEnum = "America/Indiana/Indianapolis"
	TimezoneEnum_America_Indiana_Knox           TimezoneEnum = "America/Indiana/Knox"
	TimezoneEnum_America_Indiana_Marengo        TimezoneEnum = "America/Indiana/Marengo"
	TimezoneEnum_America_Indiana_Petersburg     TimezoneEnum = "America/Indiana/Petersburg"
	TimezoneEnum_America_Indiana_Tell_City      TimezoneEnum = "America/Indiana/Tell_City"
	TimezoneEnum_America_Indiana_Vevay          TimezoneEnum = "America/Indiana/Vevay"
	TimezoneEnum_America_Indiana_Vincennes      TimezoneEnum = "America/Indiana/Vincennes"
	TimezoneEnum_America_Indiana_Winamac        TimezoneEnum = "America/Indiana/Winamac"
	TimezoneEnum_America_Inuvik                 TimezoneEnum = "America/Inuvik"
	TimezoneEnum_America_Iqaluit                TimezoneEnum = "America/Iqaluit"
	TimezoneEnum_America_Jamaica                TimezoneEnum = "America/Jamaica"
	TimezoneEnum_America_Juneau                 TimezoneEnum = "America/Juneau"
	TimezoneEnum_America_Kentucky_Louisville    TimezoneEnum = "America/Kentucky/Louisville"
	TimezoneEnum_America_Kentucky_Monticello    TimezoneEnum = "America/Kentucky/Monticello"
	TimezoneEnum_America_Kralendijk             TimezoneEnum = "America/Kralendijk"
	TimezoneEnum_America_La_Paz                 TimezoneEnum = "America/La_Paz"
	TimezoneEnum_America_Lima                   TimezoneEnum = "America/Lima"
	TimezoneEnum_America_Los_Angeles            TimezoneEnum = "America/Los_Angeles"
	TimezoneEnum_America_Lower_Princes          TimezoneEnum = "America/Lower_Princes"
	TimezoneEnum_America_Maceio                 TimezoneEnum = "America/Maceio"
	TimezoneEnum_America_Managua                TimezoneEnum = "America/Managua"
	TimezoneEnum_America_Manaus                 TimezoneEnum = "America/Manaus"
	TimezoneEnum_America_Marigot                TimezoneEnum = "America/Marigot"
	TimezoneEnum_America_Martinique             TimezoneEnum = "America/Martinique"
	TimezoneEnum_America_Matamoros              TimezoneEnum = "America/Matamoros"
	TimezoneEnum_America_Mazatlan               TimezoneEnum = "America/Mazatlan"
	TimezoneEnum_America_Menominee              TimezoneEnum = "America/Menominee"
	TimezoneEnum_America_Merida                 TimezoneEnum = "America/Merida"
	TimezoneEnum_America_Metlakatla             TimezoneEnum = "America/Metlakatla"
	TimezoneEnum_America_Mexico_City            TimezoneEnum = "America/Mexico_City"
	TimezoneEnum_America_Miquelon               TimezoneEnum = "America/Miquelon"
	TimezoneEnum_America_Moncton                TimezoneEnum = "America/Moncton"
	TimezoneEnum_America_Monterrey              TimezoneEnum = "America/Monterrey"
	TimezoneEnum_America_Montevideo             TimezoneEnum = "America/Montevideo"
	TimezoneEnum_America_Montserrat             TimezoneEnum = "America/Montserrat"
	TimezoneEnum_America_Nassau                 TimezoneEnum = "America/Nassau"
	TimezoneEnum_America_New_York               TimezoneEnum = "America/New_York"
	TimezoneEnum_America_Nome                   TimezoneEnum = "America/Nome"
	TimezoneEnum_America_Noronha                TimezoneEnum = "America/Noronha"
	TimezoneEnum_America_North_Dakota_Beulah    TimezoneEnum = "America/North_Dakota/Beulah"
	TimezoneEnum_America_North_Dakota_Center    TimezoneEnum = "America/North_Dakota/Center"
	TimezoneEnum_America_North_Dakota_New_Salem TimezoneEnum = "America/North_Dakota/New_Salem"
	TimezoneEnum_America_Nuuk                   TimezoneEnum = "America/Nuuk"
	TimezoneEnum_America_Ojinaga                TimezoneEnum = "America/Ojinaga"
	TimezoneEnum_America_Panama                 TimezoneEnum = "America/Panama"
	TimezoneEnum_America_Paramaribo             TimezoneEnum = "America/Paramaribo"
	TimezoneEnum_America_Phoenix                TimezoneEnum = "America/Phoenix"
	TimezoneEnum_America_Port_au_Prince         TimezoneEnum = "America/Port-au-Prince"
	TimezoneEnum_America_Port_of_Spain          TimezoneEnum = "America/Port_of_Spain"
	TimezoneEnum_America_Porto_Velho            TimezoneEnum = "America/Porto_Velho"
	TimezoneEnum_America_Puerto_Rico            TimezoneEnum = "America/Puerto_Rico"
	TimezoneEnum_America_Punta_Arenas           TimezoneEnum = "America/Punta_Arenas"
	TimezoneEnum_America_Rankin_Inlet           TimezoneEnum = "America/Rankin_Inlet"
	TimezoneEnum_America_Recife                 TimezoneEnum = "America/Recife"
	TimezoneEnum_America_Regina                 TimezoneEnum = "America/Regina"
	TimezoneEnum_America_Resolute               TimezoneEnum = "America/Resolute"
	TimezoneEnum_America_Rio_Branco             TimezoneEnum = "America/Rio_Branco"
	TimezoneEnum_America_Santarem               TimezoneEnum = "America/Santarem"
	TimezoneEnum_America_Santiago               TimezoneEnum = "America/Santiago"
	TimezoneEnum_America_Santo_Domingo          TimezoneEnum = "America/Santo_Domingo"
	TimezoneEnum_America_Sao_Paulo              TimezoneEnum = "America/Sao_Paulo"
	TimezoneEnum_America_Scoresbysund           TimezoneEnum = "America/Scoresbysund"
	TimezoneEnum_America_Sitka                  TimezoneEnum = "America/Sitka"
	TimezoneEnum_America_St_Barthelemy          TimezoneEnum = "America/St_Barthelemy"
	TimezoneEnum_America_St_Johns               TimezoneEnum = "America/St_Johns"
	TimezoneEnum_America_St_Kitts               TimezoneEnum = "America/St_Kitts"
	TimezoneEnum_America_St_Lucia               TimezoneEnum = "America/St_Lucia"
	TimezoneEnum_America_St_Thomas              TimezoneEnum = "America/St_Thomas"
	TimezoneEnum_America_St_Vincent             TimezoneEnum = "America/St_Vincent"
	TimezoneEnum_America_Swift_Current          TimezoneEnum = "America/Swift_Current"
	TimezoneEnum_America_Tegucigalpa            TimezoneEnum = "America/Tegucigalpa"
	TimezoneEnum_America_Thule                  TimezoneEnum = "America/Thule"
	TimezoneEnum_America_Tijuana                TimezoneEnum = "America/Tijuana"
	TimezoneEnum_America_Toronto                TimezoneEnum = "America/Toronto"
	TimezoneEnum_America_Tortola                TimezoneEnum = "America/Tortola"
	TimezoneEnum_America_Vancouver              TimezoneEnum = "America/Vancouver"
	TimezoneEnum_America_Whitehorse             TimezoneEnum = "America/Whitehorse"
	TimezoneEnum_America_Winnipeg               TimezoneEnum = "America/Winnipeg"
	TimezoneEnum_America_Yakutat                TimezoneEnum = "America/Yakutat"
	TimezoneEnum_Antarctica_Casey               TimezoneEnum = "Antarctica/Casey"
	TimezoneEnum_Antarctica_Davis               TimezoneEnum = "Antarctica/Davis"
	TimezoneEnum_Antarctica_DumontDUrville      TimezoneEnum = "Antarctica/DumontDUrville"
	TimezoneEnum_Antarctica_Macquarie           TimezoneEnum = "Antarctica/Macquarie"
	TimezoneEnum_Antarctica_Mawson              TimezoneEnum = "Antarctica/Mawson"
	TimezoneEnum_Antarctica_McMurdo             TimezoneEnum = "Antarctica/McMurdo"
	TimezoneEnum_Antarctica_Palmer              TimezoneEnum = "Antarctica/Palmer"
	TimezoneEnum_Antarctica_Rothera             TimezoneEnum = "Antarctica/Rothera"
	TimezoneEnum_Antarctica_Syowa               TimezoneEnum = "Antarctica/Syowa"
	TimezoneEnum_Antarctica_Troll               TimezoneEnum = "Antarctica/Troll"
	TimezoneEnum_Antarctica_Vostok              TimezoneEnum = "Antarctica/Vostok"
	TimezoneEnum_Arctic_Longyearbyen            TimezoneEnum = "Arctic/Longyearbyen"
	TimezoneEnum_Asia_Aden                      TimezoneEnum = "Asia/Aden"
	TimezoneEnum_Asia_Almaty                    TimezoneEnum = "Asia/Almaty"
	TimezoneEnum_Asia_Amman                     TimezoneEnum = "Asia/Amman"
	TimezoneEnum_Asia_Anadyr                    TimezoneEnum = "Asia/Anadyr"
	TimezoneEnum_Asia_Aqtau                     TimezoneEnum = "Asia/Aqtau"
	TimezoneEnum_Asia_Aqtobe                    TimezoneEnum = "Asia/Aqtobe"
	TimezoneEnum_Asia_Ashgabat                  TimezoneEnum = "Asia/Ashgabat"
	TimezoneEnum_Asia_Atyrau                    TimezoneEnum = "Asia/Atyrau"
	TimezoneEnum_Asia_Baghdad                   TimezoneEnum = "Asia/Baghdad"
	TimezoneEnum_Asia_Bahrain                   TimezoneEnum = "Asia/Bahrain"
	TimezoneEnum_Asia_Baku                      TimezoneEnum = "Asia/Baku"
	TimezoneEnum_Asia_Bangkok                   TimezoneEnum = "Asia/Bangkok"
	TimezoneEnum_Asia_Barnaul                   TimezoneEnum = "Asia/Barnaul"
	TimezoneEnum_Asia_Beirut                    TimezoneEnum = "Asia/Beirut"
	TimezoneEnum_Asia_Bishkek                   TimezoneEnum = "Asia/Bishkek"
	TimezoneEnum_Asia_Brunei                    TimezoneEnum = "Asia/Brunei"
	TimezoneEnum_Asia_Chita                     TimezoneEnum = "Asia/Chita"
	TimezoneEnum_Asia_Choibalsan                TimezoneEnum = "Asia/Choibalsan"
	TimezoneEnum_Asia_Colombo                   TimezoneEnum = "Asia/Colombo"
	TimezoneEnum_Asia_Damascus                  TimezoneEnum = "Asia/Damascus"
	TimezoneEnum_Asia_Dhaka                     TimezoneEnum = "Asia/Dhaka"
	TimezoneEnum_Asia_Dili                      TimezoneEnum = "Asia/Dili"
	TimezoneEnum_Asia_Dubai                     TimezoneEnum = "Asia/Dubai"
	TimezoneEnum_Asia_Dushanbe                  TimezoneEnum = "Asia/Dushanbe"
	TimezoneEnum_Asia_Famagusta                 TimezoneEnum = "Asia/Famagusta"
	TimezoneEnum_Asia_Gaza                      TimezoneEnum = "Asia/Gaza"
	TimezoneEnum_Asia_Hebron                    TimezoneEnum = "Asia/Hebron"
	TimezoneEnum_Asia_Ho_Chi_Minh               TimezoneEnum = "Asia/Ho_Chi_Minh"
	TimezoneEnum_Asia_Hong_Kong                 TimezoneEnum = "Asia/Hong_Kong"
	TimezoneEnum_Asia_Hovd                      TimezoneEnum = "Asia/Hovd"
	TimezoneEnum_Asia_Irkutsk                   TimezoneEnum = "Asia/Irkutsk"
	TimezoneEnum_Asia_Jakarta                   TimezoneEnum = "Asia/Jakarta"
	TimezoneEnum_Asia_Jayapura                  TimezoneEnum = "Asia/Jayapura"
	TimezoneEnum_Asia_Jerusalem                 TimezoneEnum = "Asia/Jerusalem"
	TimezoneEnum_Asia_Kabul                     TimezoneEnum = "Asia/Kabul"
	TimezoneEnum_Asia_Kamchatka                 TimezoneEnum = "Asia/Kamchatka"
	TimezoneEnum_Asia_Karachi                   TimezoneEnum = "Asia/Karachi"
	TimezoneEnum_Asia_Kathmandu                 TimezoneEnum = "Asia/Kathmandu"
	TimezoneEnum_Asia_Khandyga                  TimezoneEnum = "Asia/Khandyga"
	TimezoneEnum_Asia_Kolkata                   TimezoneEnum = "Asia/Kolkata"
	TimezoneEnum_Asia_Krasnoyarsk               TimezoneEnum = "Asia/Krasnoyarsk"
	TimezoneEnum_Asia_Kuala_Lumpur              TimezoneEnum = "Asia/Kuala_Lumpur"
	TimezoneEnum_Asia_Kuching                   TimezoneEnum = "Asia/Kuching"
	TimezoneEnum_Asia_Kuwait                    TimezoneEnum = "Asia/Kuwait"
	TimezoneEnum_Asia_Macau                     TimezoneEnum = "Asia/Macau"
	TimezoneEnum_Asia_Magadan                   TimezoneEnum = "Asia/Magadan"
	TimezoneEnum_Asia_Makassar                  TimezoneEnum = "Asia/Makassar"
	TimezoneEnum_Asia_Manila                    TimezoneEnum = "Asia/Manila"
	TimezoneEnum_Asia_Muscat                    TimezoneEnum = "Asia/Muscat"
	TimezoneEnum_Asia_Nicosia                   TimezoneEnum = "Asia/Nicosia"
	TimezoneEnum_Asia_Novokuznetsk              TimezoneEnum = "Asia/Novokuznetsk"
	TimezoneEnum_Asia_Novosibirsk               TimezoneEnum = "Asia/Novosibirsk"
	TimezoneEnum_Asia_Omsk                      TimezoneEnum = "Asia/Omsk"
	TimezoneEnum_Asia_Oral                      TimezoneEnum = "Asia/Oral"
	TimezoneEnum_Asia_Phnom_Penh                TimezoneEnum = "Asia/Phnom_Penh"
	TimezoneEnum_Asia_Pontianak                 TimezoneEnum = "Asia/Pontianak"
	TimezoneEnum_Asia_Pyongyang                 TimezoneEnum = "Asia/Pyongyang"
	TimezoneEnum_Asia_Qatar                     TimezoneEnum = "Asia/Qatar"
	TimezoneEnum_Asia_Qostanay                  TimezoneEnum = "Asia/Qostanay"
	TimezoneEnum_Asia_Qyzylorda                 TimezoneEnum = "Asia/Qyzylorda"
	TimezoneEnum_Asia_Riyadh                    TimezoneEnum = "Asia/Riyadh"
	TimezoneEnum_Asia_Sakhalin                  TimezoneEnum = "Asia/Sakhalin"
	TimezoneEnum_Asia_Samarkand                 TimezoneEnum = "Asia/Samarkand"
	TimezoneEnum_Asia_Seoul                     TimezoneEnum = "Asia/Seoul"
	TimezoneEnum_Asia_Shanghai                  TimezoneEnum = "Asia/Shanghai"
	TimezoneEnum_Asia_Singapore                 TimezoneEnum = "Asia/Singapore"
	TimezoneEnum_Asia_Srednekolymsk             TimezoneEnum = "Asia/Srednekolymsk"
	TimezoneEnum_Asia_Taipei                    TimezoneEnum = "Asia/Taipei"
	TimezoneEnum_Asia_Tashkent                  TimezoneEnum = "Asia/Tashkent"
	TimezoneEnum_Asia_Tbilisi                   TimezoneEnum = "Asia/Tbilisi"
	TimezoneEnum_Asia_Tehran                    TimezoneEnum = "Asia/Tehran"
	TimezoneEnum_Asia_Thimphu                   TimezoneEnum = "Asia/Thimphu"
	TimezoneEnum_Asia_Tokyo                     TimezoneEnum = "Asia/Tokyo"
	TimezoneEnum_Asia_Tomsk                     TimezoneEnum = "Asia/Tomsk"
	TimezoneEnum_Asia_Ulaanbaatar               TimezoneEnum = "Asia/Ulaanbaatar"
	TimezoneEnum_Asia_Urumqi                    TimezoneEnum = "Asia/Urumqi"
	TimezoneEnum_Asia_Ust_Nera                  TimezoneEnum = "Asia/Ust-Nera"
	TimezoneEnum_Asia_Vientiane                 TimezoneEnum = "Asia/Vientiane"
	TimezoneEnum_Asia_Vladivostok               TimezoneEnum = "Asia/Vladivostok"
	TimezoneEnum_Asia_Yakutsk                   TimezoneEnum = "Asia/Yakutsk"
	TimezoneEnum_Asia_Yangon                    TimezoneEnum = "Asia/Yangon"
	TimezoneEnum_Asia_Yekaterinburg             TimezoneEnum = "Asia/Yekaterinburg"
	TimezoneEnum_Asia_Yerevan                   TimezoneEnum = "Asia/Yerevan"
	TimezoneEnum_Atlantic_Azores                TimezoneEnum = "Atlantic/Azores"
	TimezoneEnum_Atlantic_Bermuda               TimezoneEnum = "Atlantic/Bermuda"
	TimezoneEnum_Atlantic_Canary                TimezoneEnum = "Atlantic/Canary"
	TimezoneEnum_Atlantic_Cape_Verde            TimezoneEnum = "Atlantic/Cape_Verde"
	TimezoneEnum_Atlantic_Faroe                 TimezoneEnum = "Atlantic/Faroe"
	TimezoneEnum_Atlantic_Madeira               TimezoneEnum = "Atlantic/Madeira"
	TimezoneEnum_Atlantic_Reykjavik             TimezoneEnum = "Atlantic/Reykjavik"
	TimezoneEnum_Atlantic_South_Georgia         TimezoneEnum = "Atlantic/South_Georgia"
	TimezoneEnum_Atlantic_St_Helena             TimezoneEnum = "Atlantic/St_Helena"
	TimezoneEnum_Atlantic_Stanley               TimezoneEnum = "Atlantic/Stanley"
	TimezoneEnum_Australia_Adelaide             TimezoneEnum = "Australia/Adelaide"
	TimezoneEnum_Australia_Brisbane             TimezoneEnum = "Australia/Brisbane"
	TimezoneEnum_Australia_Broken_Hill          TimezoneEnum = "Australia/Broken_Hill"
	TimezoneEnum_Australia_Darwin               TimezoneEnum = "Australia/Darwin"
	TimezoneEnum_Australia_Eucla                TimezoneEnum = "Australia/Eucla"
	TimezoneEnum_Australia_Hobart               TimezoneEnum = "Australia/Hobart"
	TimezoneEnum_Australia_Lindeman             TimezoneEnum = "Australia/Lindeman"
	TimezoneEnum_Australia_Lord_Howe            TimezoneEnum = "Australia/Lord_Howe"
	TimezoneEnum_Australia_Melbourne            TimezoneEnum = "Australia/Melbourne"
	TimezoneEnum_Australia_Perth                TimezoneEnum = "Australia/Perth"
	TimezoneEnum_Australia_Sydney               TimezoneEnum = "Australia/Sydney"
	TimezoneEnum_Canada_Atlantic                TimezoneEnum = "Canada/Atlantic"
	TimezoneEnum_Canada_Central                 TimezoneEnum = "Canada/Central"
	TimezoneEnum_Canada_Eastern                 TimezoneEnum = "Canada/Eastern"
	TimezoneEnum_Canada_Mountain                TimezoneEnum = "Canada/Mountain"
	TimezoneEnum_Canada_Newfoundland            TimezoneEnum = "Canada/Newfoundland"
	TimezoneEnum_Canada_Pacific                 TimezoneEnum = "Canada/Pacific"
	TimezoneEnum_Europe_Amsterdam               TimezoneEnum = "Europe/Amsterdam"
	TimezoneEnum_Europe_Andorra                 TimezoneEnum = "Europe/Andorra"
	TimezoneEnum_Europe_Astrakhan               TimezoneEnum = "Europe/Astrakhan"
	TimezoneEnum_Europe_Athens                  TimezoneEnum = "Europe/Athens"
	TimezoneEnum_Europe_Belgrade                TimezoneEnum = "Europe/Belgrade"
	TimezoneEnum_Europe_Berlin                  TimezoneEnum = "Europe/Berlin"
	TimezoneEnum_Europe_Bratislava              TimezoneEnum = "Europe/Bratislava"
	TimezoneEnum_Europe_Brussels                TimezoneEnum = "Europe/Brussels"
	TimezoneEnum_Europe_Bucharest               TimezoneEnum = "Europe/Bucharest"
	TimezoneEnum_Europe_Budapest                TimezoneEnum = "Europe/Budapest"
	TimezoneEnum_Europe_Busingen                TimezoneEnum = "Europe/Busingen"
	TimezoneEnum_Europe_Chisinau                TimezoneEnum = "Europe/Chisinau"
	TimezoneEnum_Europe_Copenhagen              TimezoneEnum = "Europe/Copenhagen"
	TimezoneEnum_Europe_Dublin                  TimezoneEnum = "Europe/Dublin"
	TimezoneEnum_Europe_Gibraltar               TimezoneEnum = "Europe/Gibraltar"
	TimezoneEnum_Europe_Guernsey                TimezoneEnum = "Europe/Guernsey"
	TimezoneEnum_Europe_Helsinki                TimezoneEnum = "Europe/Helsinki"
	TimezoneEnum_Europe_Isle_of_Man             TimezoneEnum = "Europe/Isle_of_Man"
	TimezoneEnum_Europe_Istanbul                TimezoneEnum = "Europe/Istanbul"
	TimezoneEnum_Europe_Jersey                  TimezoneEnum = "Europe/Jersey"
	TimezoneEnum_Europe_Kaliningrad             TimezoneEnum = "Europe/Kaliningrad"
	TimezoneEnum_Europe_Kirov                   TimezoneEnum = "Europe/Kirov"
	TimezoneEnum_Europe_Kyiv                    TimezoneEnum = "Europe/Kyiv"
	TimezoneEnum_Europe_Lisbon                  TimezoneEnum = "Europe/Lisbon"
	TimezoneEnum_Europe_Ljubljana               TimezoneEnum = "Europe/Ljubljana"
	TimezoneEnum_Europe_London                  TimezoneEnum = "Europe/London"
	TimezoneEnum_Europe_Luxembourg              TimezoneEnum = "Europe/Luxembourg"
	TimezoneEnum_Europe_Madrid                  TimezoneEnum = "Europe/Madrid"
	TimezoneEnum_Europe_Malta                   TimezoneEnum = "Europe/Malta"
	TimezoneEnum_Europe_Mariehamn               TimezoneEnum = "Europe/Mariehamn"
	TimezoneEnum_Europe_Minsk                   TimezoneEnum = "Europe/Minsk"
	TimezoneEnum_Europe_Monaco                  TimezoneEnum = "Europe/Monaco"
	TimezoneEnum_Europe_Moscow                  TimezoneEnum = "Europe/Moscow"
	TimezoneEnum_Europe_Oslo                    TimezoneEnum = "Europe/Oslo"
	TimezoneEnum_Europe_Paris                   TimezoneEnum = "Europe/Paris"
	TimezoneEnum_Europe_Podgorica               TimezoneEnum = "Europe/Podgorica"
	TimezoneEnum_Europe_Prague                  TimezoneEnum = "Europe/Prague"
	TimezoneEnum_Europe_Riga                    TimezoneEnum = "Europe/Riga"
	TimezoneEnum_Europe_Rome                    TimezoneEnum = "Europe/Rome"
	TimezoneEnum_Europe_Samara                  TimezoneEnum = "Europe/Samara"
	TimezoneEnum_Europe_San_Marino              TimezoneEnum = "Europe/San_Marino"
	TimezoneEnum_Europe_Sarajevo                TimezoneEnum = "Europe/Sarajevo"
	TimezoneEnum_Europe_Saratov                 TimezoneEnum = "Europe/Saratov"
	TimezoneEnum_Europe_Simferopol              TimezoneEnum = "Europe/Simferopol"
	TimezoneEnum_Europe_Skopje                  TimezoneEnum = "Europe/Skopje"
	TimezoneEnum_Europe_Sofia                   TimezoneEnum = "Europe/Sofia"
	TimezoneEnum_Europe_Stockholm               TimezoneEnum = "Europe/Stockholm"
	TimezoneEnum_Europe_Tallinn                 TimezoneEnum = "Europe/Tallinn"
	TimezoneEnum_Europe_Tirane                  TimezoneEnum = "Europe/Tirane"
	TimezoneEnum_Europe_Ulyanovsk               TimezoneEnum = "Europe/Ulyanovsk"
	TimezoneEnum_Europe_Vaduz                   TimezoneEnum = "Europe/Vaduz"
	TimezoneEnum_Europe_Vatican                 TimezoneEnum = "Europe/Vatican"
	TimezoneEnum_Europe_Vienna                  TimezoneEnum = "Europe/Vienna"
	TimezoneEnum_Europe_Vilnius                 TimezoneEnum = "Europe/Vilnius"
	TimezoneEnum_Europe_Volgograd               TimezoneEnum = "Europe/Volgograd"
	TimezoneEnum_Europe_Warsaw                  TimezoneEnum = "Europe/Warsaw"
	TimezoneEnum_Europe_Zagreb                  TimezoneEnum = "Europe/Zagreb"
	TimezoneEnum_Europe_Zurich                  TimezoneEnum = "Europe/Zurich"
	TimezoneEnum_GMT                            TimezoneEnum = "GMT"
	TimezoneEnum_Indian_Antananarivo            TimezoneEnum = "Indian/Antananarivo"
	TimezoneEnum_Indian_Chagos                  TimezoneEnum = "Indian/Chagos"
	TimezoneEnum_Indian_Christmas               TimezoneEnum = "Indian/Christmas"
	TimezoneEnum_Indian_Cocos                   TimezoneEnum = "Indian/Cocos"
	TimezoneEnum_Indian_Comoro                  TimezoneEnum = "Indian/Comoro"
	TimezoneEnum_Indian_Kerguelen               TimezoneEnum = "Indian/Kerguelen"
	TimezoneEnum_Indian_Mahe                    TimezoneEnum = "Indian/Mahe"
	TimezoneEnum_Indian_Maldives                TimezoneEnum = "Indian/Maldives"
	TimezoneEnum_Indian_Mauritius               TimezoneEnum = "Indian/Mauritius"
	TimezoneEnum_Indian_Mayotte                 TimezoneEnum = "Indian/Mayotte"
	TimezoneEnum_Indian_Reunion                 TimezoneEnum = "Indian/Reunion"
	TimezoneEnum_Pacific_Apia                   TimezoneEnum = "Pacific/Apia"
	TimezoneEnum_Pacific_Auckland               TimezoneEnum = "Pacific/Auckland"
	TimezoneEnum_Pacific_Bougainville           TimezoneEnum = "Pacific/Bougainville"
	TimezoneEnum_Pacific_Chatham                TimezoneEnum = "Pacific/Chatham"
	TimezoneEnum_Pacific_Chuuk                  TimezoneEnum = "Pacific/Chuuk"
	TimezoneEnum_Pacific_Easter                 TimezoneEnum = "Pacific/Easter"
	TimezoneEnum_Pacific_Efate                  TimezoneEnum = "Pacific/Efate"
	TimezoneEnum_Pacific_Fakaofo                TimezoneEnum = "Pacific/Fakaofo"
	TimezoneEnum_Pacific_Fiji                   TimezoneEnum = "Pacific/Fiji"
	TimezoneEnum_Pacific_Funafuti               TimezoneEnum = "Pacific/Funafuti"
	TimezoneEnum_Pacific_Galapagos              TimezoneEnum = "Pacific/Galapagos"
	TimezoneEnum_Pacific_Gambier                TimezoneEnum = "Pacific/Gambier"
	TimezoneEnum_Pacific_Guadalcanal            TimezoneEnum = "Pacific/Guadalcanal"
	TimezoneEnum_Pacific_Guam                   TimezoneEnum = "Pacific/Guam"
	TimezoneEnum_Pacific_Honolulu               TimezoneEnum = "Pacific/Honolulu"
	TimezoneEnum_Pacific_Kanton                 TimezoneEnum = "Pacific/Kanton"
	TimezoneEnum_Pacific_Kiritimati             TimezoneEnum = "Pacific/Kiritimati"
	TimezoneEnum_Pacific_Kosrae                 TimezoneEnum = "Pacific/Kosrae"
	TimezoneEnum_Pacific_Kwajalein              TimezoneEnum = "Pacific/Kwajalein"
	TimezoneEnum_Pacific_Majuro                 TimezoneEnum = "Pacific/Majuro"
	TimezoneEnum_Pacific_Marquesas              TimezoneEnum = "Pacific/Marquesas"
	TimezoneEnum_Pacific_Midway                 TimezoneEnum = "Pacific/Midway"
	TimezoneEnum_Pacific_Nauru                  TimezoneEnum = "Pacific/Nauru"
	TimezoneEnum_Pacific_Niue                   TimezoneEnum = "Pacific/Niue"
	TimezoneEnum_Pacific_Norfolk                TimezoneEnum = "Pacific/Norfolk"
	TimezoneEnum_Pacific_Noumea                 TimezoneEnum = "Pacific/Noumea"
	TimezoneEnum_Pacific_Pago_Pago              TimezoneEnum = "Pacific/Pago_Pago"
	TimezoneEnum_Pacific_Palau                  TimezoneEnum = "Pacific/Palau"
	TimezoneEnum_Pacific_Pitcairn               TimezoneEnum = "Pacific/Pitcairn"
	TimezoneEnum_Pacific_Pohnpei                TimezoneEnum = "Pacific/Pohnpei"
	TimezoneEnum_Pacific_Port_Moresby           TimezoneEnum = "Pacific/Port_Moresby"
	TimezoneEnum_Pacific_Rarotonga              TimezoneEnum = "Pacific/Rarotonga"
	TimezoneEnum_Pacific_Saipan                 TimezoneEnum = "Pacific/Saipan"
	TimezoneEnum_Pacific_Tahiti                 TimezoneEnum = "Pacific/Tahiti"
	TimezoneEnum_Pacific_Tarawa                 TimezoneEnum = "Pacific/Tarawa"
	TimezoneEnum_Pacific_Tongatapu              TimezoneEnum = "Pacific/Tongatapu"
	TimezoneEnum_Pacific_Wake                   TimezoneEnum = "Pacific/Wake"
	TimezoneEnum_Pacific_Wallis                 TimezoneEnum = "Pacific/Wallis"
	TimezoneEnum_US_Alaska                      TimezoneEnum = "US/Alaska"
	TimezoneEnum_US_Arizona                     TimezoneEnum = "US/Arizona"
	TimezoneEnum_US_Central                     TimezoneEnum = "US/Central"
	TimezoneEnum_US_Eastern                     TimezoneEnum = "US/Eastern"
	TimezoneEnum_US_Hawaii                      TimezoneEnum = "US/Hawaii"
	TimezoneEnum_US_Mountain                    TimezoneEnum = "US/Mountain"
	TimezoneEnum_US_Pacific                     TimezoneEnum = "US/Pacific"
	TimezoneEnum_UTC                            TimezoneEnum = "UTC"
)

type UnitEnum int

const (
	// No unit
	UnitEnum_No_unit UnitEnum = 0
	// %
	UnitEnum_percent UnitEnum = 1
	UnitEnum_byte    UnitEnum = 2
	UnitEnum_bit     UnitEnum = 3
	// io/s
	UnitEnum_io_s UnitEnum = 4
	// /s
	UnitEnum_s      UnitEnum = 5
	UnitEnum_second UnitEnum = 6
	UnitEnum_Custom UnitEnum = 7
	UnitEnum_day    UnitEnum = 8
	// °C
	UnitEnum_C UnitEnum = 9
	// byte/s
	UnitEnum_byte_s UnitEnum = 10
	// bit/s
	UnitEnum_bit_s UnitEnum = 11
	UnitEnum_Hz    UnitEnum = 12
	UnitEnum_W     UnitEnum = 13
)
