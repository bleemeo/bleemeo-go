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

type AgentTypeEnum string

const (
	AgentType_AWS_Account        AgentTypeEnum = "aws_account"
	AgentType_AWS_TrustedAdvisor AgentTypeEnum = "aws_trusted_advisor"
	AgentType_AWS_DynamoDB       AgentTypeEnum = "aws_dynamodb"
	AgentType_AWS_EC2            AgentTypeEnum = "aws_ec2"
	AgentType_AWS_ELB            AgentTypeEnum = "aws_elb"
	AgentType_AWS_RDS            AgentTypeEnum = "aws_rds"
	AgentType_AWS_S3             AgentTypeEnum = "aws_s3"
	AgentType_Agent              AgentTypeEnum = "agent"
	AgentType_Monitor            AgentTypeEnum = "connection_check"
	AgentType_Snmp               AgentTypeEnum = "snmp"
	AgentType_K8s                AgentTypeEnum = "kubernetes"
	AgentType_vSphereCluster     AgentTypeEnum = "vsphere_cluster"
	AgentType_vSphereHost        AgentTypeEnum = "vsphere_host"
	AgentType_vSphereVM          AgentTypeEnum = "vsphere_vm"
)

type DisconnectionReasonEnum int

const (
	DisconnectionReason_CleanShutdown    DisconnectionReasonEnum = 1
	DisconnectionReason_AgentTimeout     DisconnectionReasonEnum = 2
	DisconnectionReason_AgentAutoUpgrade DisconnectionReasonEnum = 3
	DisconnectionReason_AgentUpgrade     DisconnectionReasonEnum = 4
)

type GloutonDiagnosticEnum int

const (
	Type_Crash    GloutonDiagnosticEnum = 0
	Type_OnDemand GloutonDiagnosticEnum = 1
)

type GraphEnum int

const (
	Graph_Line                 GraphEnum = 0
	Graph_Stack                GraphEnum = 1
	Graph_Pie                  GraphEnum = 2
	Graph_Gauge                GraphEnum = 3
	Graph_AvailabilityTimeline GraphEnum = 4
	Graph_Number               GraphEnum = 5
	Graph_Status               GraphEnum = 6
	Graph_SnmpStatus           GraphEnum = 7
	Graph_Text                 GraphEnum = 8
	Graph_Image                GraphEnum = 9
	Graph_HeatmapStatus        GraphEnum = 10
	Graph_Bar                  GraphEnum = 11
)

type ReportPeriodEnum int

const (
	ReportPeriod_Weekly  ReportPeriodEnum = 0
	ReportPeriod_Monthly ReportPeriodEnum = 1
)

type ReportIncludedEnum int

const (
	ReportIncluded_None    ReportIncludedEnum = 0
	ReportIncluded_Partial ReportIncludedEnum = 1
	ReportIncluded_Full    ReportIncludedEnum = 2
)

type StatusEnum int

const (
	Status_Ok       StatusEnum = 0
	Status_Warning  StatusEnum = 1
	Status_Critical StatusEnum = 2
	Status_Unknown  StatusEnum = 3
)
