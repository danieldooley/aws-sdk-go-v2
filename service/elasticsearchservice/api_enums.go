// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticsearchservice

type DeploymentStatus string

// Enum values for DeploymentStatus
const (
	DeploymentStatusPendingUpdate DeploymentStatus = "PENDING_UPDATE"
	DeploymentStatusInProgress    DeploymentStatus = "IN_PROGRESS"
	DeploymentStatusCompleted     DeploymentStatus = "COMPLETED"
	DeploymentStatusNotEligible   DeploymentStatus = "NOT_ELIGIBLE"
	DeploymentStatusEligible      DeploymentStatus = "ELIGIBLE"
)

func (enum DeploymentStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeploymentStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DescribePackagesFilterName string

// Enum values for DescribePackagesFilterName
const (
	DescribePackagesFilterNamePackageId     DescribePackagesFilterName = "PackageID"
	DescribePackagesFilterNamePackageName   DescribePackagesFilterName = "PackageName"
	DescribePackagesFilterNamePackageStatus DescribePackagesFilterName = "PackageStatus"
)

func (enum DescribePackagesFilterName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DescribePackagesFilterName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DomainPackageStatus string

// Enum values for DomainPackageStatus
const (
	DomainPackageStatusAssociating        DomainPackageStatus = "ASSOCIATING"
	DomainPackageStatusAssociationFailed  DomainPackageStatus = "ASSOCIATION_FAILED"
	DomainPackageStatusActive             DomainPackageStatus = "ACTIVE"
	DomainPackageStatusDissociating       DomainPackageStatus = "DISSOCIATING"
	DomainPackageStatusDissociationFailed DomainPackageStatus = "DISSOCIATION_FAILED"
)

func (enum DomainPackageStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DomainPackageStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ESPartitionInstanceType string

// Enum values for ESPartitionInstanceType
const (
	ESPartitionInstanceTypeM3MediumElasticsearch         ESPartitionInstanceType = "m3.medium.elasticsearch"
	ESPartitionInstanceTypeM3LargeElasticsearch          ESPartitionInstanceType = "m3.large.elasticsearch"
	ESPartitionInstanceTypeM3XlargeElasticsearch         ESPartitionInstanceType = "m3.xlarge.elasticsearch"
	ESPartitionInstanceTypeM32xlargeElasticsearch        ESPartitionInstanceType = "m3.2xlarge.elasticsearch"
	ESPartitionInstanceTypeM4LargeElasticsearch          ESPartitionInstanceType = "m4.large.elasticsearch"
	ESPartitionInstanceTypeM4XlargeElasticsearch         ESPartitionInstanceType = "m4.xlarge.elasticsearch"
	ESPartitionInstanceTypeM42xlargeElasticsearch        ESPartitionInstanceType = "m4.2xlarge.elasticsearch"
	ESPartitionInstanceTypeM44xlargeElasticsearch        ESPartitionInstanceType = "m4.4xlarge.elasticsearch"
	ESPartitionInstanceTypeM410xlargeElasticsearch       ESPartitionInstanceType = "m4.10xlarge.elasticsearch"
	ESPartitionInstanceTypeM5LargeElasticsearch          ESPartitionInstanceType = "m5.large.elasticsearch"
	ESPartitionInstanceTypeM5XlargeElasticsearch         ESPartitionInstanceType = "m5.xlarge.elasticsearch"
	ESPartitionInstanceTypeM52xlargeElasticsearch        ESPartitionInstanceType = "m5.2xlarge.elasticsearch"
	ESPartitionInstanceTypeM54xlargeElasticsearch        ESPartitionInstanceType = "m5.4xlarge.elasticsearch"
	ESPartitionInstanceTypeM512xlargeElasticsearch       ESPartitionInstanceType = "m5.12xlarge.elasticsearch"
	ESPartitionInstanceTypeR5LargeElasticsearch          ESPartitionInstanceType = "r5.large.elasticsearch"
	ESPartitionInstanceTypeR5XlargeElasticsearch         ESPartitionInstanceType = "r5.xlarge.elasticsearch"
	ESPartitionInstanceTypeR52xlargeElasticsearch        ESPartitionInstanceType = "r5.2xlarge.elasticsearch"
	ESPartitionInstanceTypeR54xlargeElasticsearch        ESPartitionInstanceType = "r5.4xlarge.elasticsearch"
	ESPartitionInstanceTypeR512xlargeElasticsearch       ESPartitionInstanceType = "r5.12xlarge.elasticsearch"
	ESPartitionInstanceTypeC5LargeElasticsearch          ESPartitionInstanceType = "c5.large.elasticsearch"
	ESPartitionInstanceTypeC5XlargeElasticsearch         ESPartitionInstanceType = "c5.xlarge.elasticsearch"
	ESPartitionInstanceTypeC52xlargeElasticsearch        ESPartitionInstanceType = "c5.2xlarge.elasticsearch"
	ESPartitionInstanceTypeC54xlargeElasticsearch        ESPartitionInstanceType = "c5.4xlarge.elasticsearch"
	ESPartitionInstanceTypeC59xlargeElasticsearch        ESPartitionInstanceType = "c5.9xlarge.elasticsearch"
	ESPartitionInstanceTypeC518xlargeElasticsearch       ESPartitionInstanceType = "c5.18xlarge.elasticsearch"
	ESPartitionInstanceTypeUltrawarm1MediumElasticsearch ESPartitionInstanceType = "ultrawarm1.medium.elasticsearch"
	ESPartitionInstanceTypeUltrawarm1LargeElasticsearch  ESPartitionInstanceType = "ultrawarm1.large.elasticsearch"
	ESPartitionInstanceTypeT2MicroElasticsearch          ESPartitionInstanceType = "t2.micro.elasticsearch"
	ESPartitionInstanceTypeT2SmallElasticsearch          ESPartitionInstanceType = "t2.small.elasticsearch"
	ESPartitionInstanceTypeT2MediumElasticsearch         ESPartitionInstanceType = "t2.medium.elasticsearch"
	ESPartitionInstanceTypeR3LargeElasticsearch          ESPartitionInstanceType = "r3.large.elasticsearch"
	ESPartitionInstanceTypeR3XlargeElasticsearch         ESPartitionInstanceType = "r3.xlarge.elasticsearch"
	ESPartitionInstanceTypeR32xlargeElasticsearch        ESPartitionInstanceType = "r3.2xlarge.elasticsearch"
	ESPartitionInstanceTypeR34xlargeElasticsearch        ESPartitionInstanceType = "r3.4xlarge.elasticsearch"
	ESPartitionInstanceTypeR38xlargeElasticsearch        ESPartitionInstanceType = "r3.8xlarge.elasticsearch"
	ESPartitionInstanceTypeI2XlargeElasticsearch         ESPartitionInstanceType = "i2.xlarge.elasticsearch"
	ESPartitionInstanceTypeI22xlargeElasticsearch        ESPartitionInstanceType = "i2.2xlarge.elasticsearch"
	ESPartitionInstanceTypeD2XlargeElasticsearch         ESPartitionInstanceType = "d2.xlarge.elasticsearch"
	ESPartitionInstanceTypeD22xlargeElasticsearch        ESPartitionInstanceType = "d2.2xlarge.elasticsearch"
	ESPartitionInstanceTypeD24xlargeElasticsearch        ESPartitionInstanceType = "d2.4xlarge.elasticsearch"
	ESPartitionInstanceTypeD28xlargeElasticsearch        ESPartitionInstanceType = "d2.8xlarge.elasticsearch"
	ESPartitionInstanceTypeC4LargeElasticsearch          ESPartitionInstanceType = "c4.large.elasticsearch"
	ESPartitionInstanceTypeC4XlargeElasticsearch         ESPartitionInstanceType = "c4.xlarge.elasticsearch"
	ESPartitionInstanceTypeC42xlargeElasticsearch        ESPartitionInstanceType = "c4.2xlarge.elasticsearch"
	ESPartitionInstanceTypeC44xlargeElasticsearch        ESPartitionInstanceType = "c4.4xlarge.elasticsearch"
	ESPartitionInstanceTypeC48xlargeElasticsearch        ESPartitionInstanceType = "c4.8xlarge.elasticsearch"
	ESPartitionInstanceTypeR4LargeElasticsearch          ESPartitionInstanceType = "r4.large.elasticsearch"
	ESPartitionInstanceTypeR4XlargeElasticsearch         ESPartitionInstanceType = "r4.xlarge.elasticsearch"
	ESPartitionInstanceTypeR42xlargeElasticsearch        ESPartitionInstanceType = "r4.2xlarge.elasticsearch"
	ESPartitionInstanceTypeR44xlargeElasticsearch        ESPartitionInstanceType = "r4.4xlarge.elasticsearch"
	ESPartitionInstanceTypeR48xlargeElasticsearch        ESPartitionInstanceType = "r4.8xlarge.elasticsearch"
	ESPartitionInstanceTypeR416xlargeElasticsearch       ESPartitionInstanceType = "r4.16xlarge.elasticsearch"
	ESPartitionInstanceTypeI3LargeElasticsearch          ESPartitionInstanceType = "i3.large.elasticsearch"
	ESPartitionInstanceTypeI3XlargeElasticsearch         ESPartitionInstanceType = "i3.xlarge.elasticsearch"
	ESPartitionInstanceTypeI32xlargeElasticsearch        ESPartitionInstanceType = "i3.2xlarge.elasticsearch"
	ESPartitionInstanceTypeI34xlargeElasticsearch        ESPartitionInstanceType = "i3.4xlarge.elasticsearch"
	ESPartitionInstanceTypeI38xlargeElasticsearch        ESPartitionInstanceType = "i3.8xlarge.elasticsearch"
	ESPartitionInstanceTypeI316xlargeElasticsearch       ESPartitionInstanceType = "i3.16xlarge.elasticsearch"
)

func (enum ESPartitionInstanceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ESPartitionInstanceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ESWarmPartitionInstanceType string

// Enum values for ESWarmPartitionInstanceType
const (
	ESWarmPartitionInstanceTypeUltrawarm1MediumElasticsearch ESWarmPartitionInstanceType = "ultrawarm1.medium.elasticsearch"
	ESWarmPartitionInstanceTypeUltrawarm1LargeElasticsearch  ESWarmPartitionInstanceType = "ultrawarm1.large.elasticsearch"
)

func (enum ESWarmPartitionInstanceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ESWarmPartitionInstanceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// Type of Log File, it can be one of the following:
//    * INDEX_SLOW_LOGS: Index slow logs contain insert requests that took more
//    time than configured index query log threshold to execute.
//
//    * SEARCH_SLOW_LOGS: Search slow logs contain search queries that took
//    more time than configured search query log threshold to execute.
//
//    * ES_APPLICATION_LOGS: Elasticsearch application logs contain information
//    about errors and warnings raised during the operation of the service and
//    can be useful for troubleshooting.
type LogType string

// Enum values for LogType
const (
	LogTypeIndexSlowLogs     LogType = "INDEX_SLOW_LOGS"
	LogTypeSearchSlowLogs    LogType = "SEARCH_SLOW_LOGS"
	LogTypeEsApplicationLogs LogType = "ES_APPLICATION_LOGS"
)

func (enum LogType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum LogType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The state of a requested change. One of the following:
//
//    * Processing: The request change is still in-process.
//
//    * Active: The request change is processed and deployed to the Elasticsearch
//    domain.
type OptionState string

// Enum values for OptionState
const (
	OptionStateRequiresIndexDocuments OptionState = "RequiresIndexDocuments"
	OptionStateProcessing             OptionState = "Processing"
	OptionStateActive                 OptionState = "Active"
)

func (enum OptionState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum OptionState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PackageStatus string

// Enum values for PackageStatus
const (
	PackageStatusCopying          PackageStatus = "COPYING"
	PackageStatusCopyFailed       PackageStatus = "COPY_FAILED"
	PackageStatusValidating       PackageStatus = "VALIDATING"
	PackageStatusValidationFailed PackageStatus = "VALIDATION_FAILED"
	PackageStatusAvailable        PackageStatus = "AVAILABLE"
	PackageStatusDeleting         PackageStatus = "DELETING"
	PackageStatusDeleted          PackageStatus = "DELETED"
	PackageStatusDeleteFailed     PackageStatus = "DELETE_FAILED"
)

func (enum PackageStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PackageStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PackageType string

// Enum values for PackageType
const (
	PackageTypeTxtDictionary PackageType = "TXT-DICTIONARY"
)

func (enum PackageType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PackageType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReservedElasticsearchInstancePaymentOption string

// Enum values for ReservedElasticsearchInstancePaymentOption
const (
	ReservedElasticsearchInstancePaymentOptionAllUpfront     ReservedElasticsearchInstancePaymentOption = "ALL_UPFRONT"
	ReservedElasticsearchInstancePaymentOptionPartialUpfront ReservedElasticsearchInstancePaymentOption = "PARTIAL_UPFRONT"
	ReservedElasticsearchInstancePaymentOptionNoUpfront      ReservedElasticsearchInstancePaymentOption = "NO_UPFRONT"
)

func (enum ReservedElasticsearchInstancePaymentOption) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReservedElasticsearchInstancePaymentOption) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TLSSecurityPolicy string

// Enum values for TLSSecurityPolicy
const (
	TLSSecurityPolicyPolicyMinTls10201907 TLSSecurityPolicy = "Policy-Min-TLS-1-0-2019-07"
	TLSSecurityPolicyPolicyMinTls12201907 TLSSecurityPolicy = "Policy-Min-TLS-1-2-2019-07"
)

func (enum TLSSecurityPolicy) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TLSSecurityPolicy) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UpgradeStatus string

// Enum values for UpgradeStatus
const (
	UpgradeStatusInProgress          UpgradeStatus = "IN_PROGRESS"
	UpgradeStatusSucceeded           UpgradeStatus = "SUCCEEDED"
	UpgradeStatusSucceededWithIssues UpgradeStatus = "SUCCEEDED_WITH_ISSUES"
	UpgradeStatusFailed              UpgradeStatus = "FAILED"
)

func (enum UpgradeStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpgradeStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type UpgradeStep string

// Enum values for UpgradeStep
const (
	UpgradeStepPreUpgradeCheck UpgradeStep = "PRE_UPGRADE_CHECK"
	UpgradeStepSnapshot        UpgradeStep = "SNAPSHOT"
	UpgradeStepUpgrade         UpgradeStep = "UPGRADE"
)

func (enum UpgradeStep) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum UpgradeStep) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The type of EBS volume, standard, gp2, or io1. See Configuring EBS-based
// Storage (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-ebs)for
// more information.
type VolumeType string

// Enum values for VolumeType
const (
	VolumeTypeStandard VolumeType = "standard"
	VolumeTypeGp2      VolumeType = "gp2"
	VolumeTypeIo1      VolumeType = "io1"
)

func (enum VolumeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum VolumeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
