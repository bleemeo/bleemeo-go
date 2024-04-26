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

type AgentTypeType string

const (
	AgentType_AWS_Account        AgentTypeType = "aws_account"
	AgentType_AWS_TrustedAdvisor AgentTypeType = "aws_trusted_advisor"
	AgentType_AWS_DynamoDB       AgentTypeType = "aws_dynamodb"
	AgentType_AWS_EC2            AgentTypeType = "aws_ec2"
	AgentType_AWS_ELB            AgentTypeType = "aws_elb"
	AgentType_AWS_RDS            AgentTypeType = "aws_rds"
	AgentType_AWS_S3             AgentTypeType = "aws_s3"
	AgentType_Agent              AgentTypeType = "agent"
	AgentType_Monitor            AgentTypeType = "connection_check"
	AgentType_Snmp               AgentTypeType = "snmp"
	AgentType_K8s                AgentTypeType = "kubernetes"
	AgentType_vSphereCluster     AgentTypeType = "vsphere_cluster"
	AgentType_vSphereHost        AgentTypeType = "vsphere_host"
	AgentType_vSphereVM          AgentTypeType = "vsphere_vm"
)

type DisconnectionReasonType int

const (
	DisconnectionReason_CleanShutdown    DisconnectionReasonType = 1
	DisconnectionReason_AgentTimeout     DisconnectionReasonType = 2
	DisconnectionReason_AgentAutoUpgrade DisconnectionReasonType = 3
	DisconnectionReason_AgentUpgrade     DisconnectionReasonType = 4
)

type GloutonDiagnosticType int

const (
	GloutonDiagnostic_Crash    GloutonDiagnosticType = 0
	GloutonDiagnostic_OnDemand GloutonDiagnosticType = 1
)

type GraphType int

const (
	Graph_Line                 GraphType = 0
	Graph_Stack                GraphType = 1
	Graph_Pie                  GraphType = 2
	Graph_Gauge                GraphType = 3
	Graph_AvailabilityTimeline GraphType = 4
	Graph_Number               GraphType = 5
	Graph_Status               GraphType = 6
	Graph_SnmpStatus           GraphType = 7
	Graph_Text                 GraphType = 8
	Graph_Image                GraphType = 9
	Graph_HeatmapStatus        GraphType = 10
	Graph_Bar                  GraphType = 11
)

type ReportPeriodType int

const (
	ReportPeriod_Weekly  ReportPeriodType = 0
	ReportPeriod_Monthly ReportPeriodType = 1
)

type ReportIncludedType int

const (
	ReportIncluded_None    ReportIncludedType = 0
	ReportIncluded_Partial ReportIncludedType = 1
	ReportIncluded_Full    ReportIncludedType = 2
)

type StatusType int

const (
	Status_Ok       StatusType = 0
	Status_Warning  StatusType = 1
	Status_Critical StatusType = 2
	Status_Unknown  StatusType = 3
)
