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

//nolint: gofmt,gofumpt,goimports
package bleemeo

type AgentType string

const (
	AgentType_AWS_Account        AgentType = "aws_account"
	AgentType_AWS_TrustedAdvisor AgentType = "aws_trusted_advisor"
	AgentType_AWS_DynamoDB       AgentType = "aws_dynamodb"
	AgentType_AWS_EC2            AgentType = "aws_ec2"
	AgentType_AWS_ELB            AgentType = "aws_elb"
	AgentType_AWS_RDS            AgentType = "aws_rds"
	AgentType_AWS_S3             AgentType = "aws_s3"
	AgentType_Agent              AgentType = "agent"
	AgentType_Monitor            AgentType = "connection_check"
	AgentType_SNMP               AgentType = "snmp"
	AgentType_K8s                AgentType = "kubernetes"
	AgentType_vSphereCluster     AgentType = "vsphere_cluster"
	AgentType_vSphereHost        AgentType = "vsphere_host"
	AgentType_vSphereVM          AgentType = "vsphere_vm"
	AgentType_Application        AgentType = "application"
)

type DisconnectionReason int

const (
	DisconnectionReason_CleanShutdown    DisconnectionReason = 1
	DisconnectionReason_AgentTimeout     DisconnectionReason = 2
	DisconnectionReason_AgentAutoUpgrade DisconnectionReason = 3
	DisconnectionReason_AgentUpgrade     DisconnectionReason = 4
)

type GloutonDiagnostic int

const (
	GloutonDiagnostic_Crash    GloutonDiagnostic = 0
	GloutonDiagnostic_OnDemand GloutonDiagnostic = 1
)

type Graph int

const (
	Graph_Line                 Graph = 0
	Graph_Stack                Graph = 1
	Graph_Pie                  Graph = 2
	Graph_Gauge                Graph = 3
	Graph_AvailabilityTimeline Graph = 4
	Graph_Number               Graph = 5
	Graph_Status               Graph = 6
	Graph_SnmpStatus           Graph = 7
	Graph_Text                 Graph = 8
	Graph_Image                Graph = 9
	Graph_HeatmapStatus        Graph = 10
	Graph_Bar                  Graph = 11
)

type ReportPeriod int

const (
	ReportPeriod_Weekly  ReportPeriod = 0
	ReportPeriod_Monthly ReportPeriod = 1
)

type ReportIncluded int

const (
	ReportIncluded_None    ReportIncluded = 0
	ReportIncluded_Partial ReportIncluded = 1
	ReportIncluded_Full    ReportIncluded = 2
)

type ConfigItemSource int

const (
	ConfigItemSource_Unknown ConfigItemSource = 0
	ConfigItemSource_Default ConfigItemSource = 1
	ConfigItemSource_File    ConfigItemSource = 2
	ConfigItemSource_Env     ConfigItemSource = 3
	ConfigItemSource_API     ConfigItemSource = 4
)

type ConfigItemType int

const (
	ConfigItemType_Any               ConfigItemType = 0
	ConfigItemType_Int               ConfigItemType = 1
	ConfigItemType_Float             ConfigItemType = 2
	ConfigItemType_Bool              ConfigItemType = 3
	ConfigItemType_String            ConfigItemType = 4
	ConfigItemType_ListString        ConfigItemType = 10
	ConfigItemType_ListInt           ConfigItemType = 11
	ConfigItemType_MapStrStr         ConfigItemType = 20
	ConfigItemType_MapStrInt         ConfigItemType = 21
	ConfigItemType_Thresholds        ConfigItemType = 30
	ConfigItemType_Services          ConfigItemType = 31
	ConfigItemType_NameInstances     ConfigItemType = 32
	ConfigItemType_BlackboxTargets   ConfigItemType = 33
	ConfigItemType_PrometheusTargets ConfigItemType = 34
	ConfigItemType_SNMPTargets       ConfigItemType = 35
	ConfigItemType_LogInputs         ConfigItemType = 36
)

type Status int

const (
	Status_Ok       Status = 0
	Status_Warning  Status = 1
	Status_Critical Status = 2
	Status_Unknown  Status = 3
)

type TagType int

const (
	TagType_AutomaticAPI        TagType = 0
	TagType_CreatedByGlouton    TagType = 1
	TagType_CreatedByFrontend   TagType = 2
	TagType_AutomaticGlouton    TagType = 3
	TagType_AutomaticAPIService TagType = 4
	TagType_NoType              TagType = 10
)
